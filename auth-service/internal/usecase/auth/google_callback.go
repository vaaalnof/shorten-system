package auth

import (
	"auth-service/internal/entity"
	avatarutil "auth-service/internal/utils/avatar"
	"context"
	"database/sql"
	"github.com/google/uuid"

	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/model/converter"
	"auth-service/internal/repository"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"

	"github.com/go-playground/validator/v10"
)

type GoogleCallbackUseCase struct {
	repo *repository.Repository

	validate *validator.Validate

	userRepo             port.UserRepository
	userAuthProviderRepo port.UserAuthProviderRepository

	createSession *CreateSession

	cache port.Cache

	googleOAuth port.GoogleOAuthService
}

func NewGoogleCallbackUseCase(
	repo *repository.Repository,
	validate *validator.Validate,
	userRepo port.UserRepository,
	userAuthProviderRepo port.UserAuthProviderRepository,
	createSession *CreateSession,
	cache port.Cache,
	googleOAuth port.GoogleOAuthService,
) *GoogleCallbackUseCase {

	return &GoogleCallbackUseCase{
		repo: repo,

		validate: validate,

		userRepo: userRepo,

		userAuthProviderRepo: userAuthProviderRepo,

		createSession: createSession,

		cache: cache,

		googleOAuth: googleOAuth,
	}
}

func (u *GoogleCallbackUseCase) Execute(
	ctx context.Context,
	req *model.GoogleCallbackRequest,
) (*model.WebResponse[*model.LoginResponse], error) {

	// =====================================================
	// VALIDATION
	// =====================================================

	if err := u.validate.Struct(
		req,
	); err != nil {

		return nil, exception.Validation(
			err,
		)
	}

	// =====================================================
	// CHECK OAUTH STATE
	// =====================================================

	if !u.cache.Exists(
		ctx,
		cachekey.OAuthState(
			"google",
			req.State,
		),
	) {

		return nil, exception.BadRequest(
			"invalid oauth state",
		)
	}

	// =====================================================
	// DELETE OAUTH STATE
	// =====================================================

	_ = u.cache.Delete(
		ctx,
		cachekey.OAuthState(
			"google",
			req.State,
		),
	)

	// =====================================================
	// EXCHANGE GOOGLE USER
	// =====================================================

	googleUser, err := u.googleOAuth.Exchange(
		ctx,
		req.Code,
	)

	if err != nil {

		return nil, exception.Unauthorized(
			"google authentication failed",
		)
	}

	// =====================================================
	// FIND GOOGLE PROVIDER
	// =====================================================

	authProvider, err := u.userAuthProviderRepo.FindByProviderUserID(
		ctx,
		string(entity.AuthProviderGoogle),
		googleUser.ID,
	)

	if err != nil {

		return nil, exception.Internal(
			"google login failed",
		)
	}

	// =====================================================
	// GOOGLE PROVIDER EXISTS
	// =====================================================

	if authProvider != nil {

		user, err := u.userRepo.FindByUserID(
			ctx,
			authProvider.UserID,
		)

		if err != nil {

			return nil, exception.Internal(
				"google login failed",
			)
		}

		if user == nil {

			return nil, exception.Internal(
				"user not found",
			)
		}

		if !user.IsActive {

			return nil, exception.Unauthorized(
				"account is inactive",
			)
		}

		loginResponse, err := u.createSession.Execute(
			ctx,
			user,
		)

		if err != nil {
			return nil, err
		}

		return &model.WebResponse[*model.LoginResponse]{
			Message: "login success",
			Data: converter.ToLoginResponse(
				loginResponse.AccessToken,
				loginResponse.RefreshToken,
			),
		}, nil
	}

	// =====================================================
	// FIND USER BY EMAIL
	// =====================================================

	user, err := u.userRepo.FindByEmail(
		ctx,
		googleUser.Email,
	)

	if err != nil {

		return nil, exception.Internal(
			"google login failed",
		)
	}

	// =====================================================
	// USER EXISTS
	// =====================================================

	if user != nil {

		if !user.IsActive {

			return nil, exception.Unauthorized(
				"account is inactive",
			)
		}

		// =====================================================
		// CHECK GOOGLE PROVIDER
		// =====================================================

		existingProvider, err := u.userAuthProviderRepo.FindByUserIDAndProvider(
			ctx,
			user.ID,
			"google",
		)

		if err != nil {

			return nil, exception.Internal(
				"google login failed",
			)
		}

		// =====================================================
		// LINK GOOGLE PROVIDER
		// =====================================================

		if existingProvider == nil {

			providerUserID := googleUser.ID

			authProvider = &entity.UserAuthProvider{
				ID:             uuid.NewString(),
				UserID:         user.ID,
				Provider:       "google",
				ProviderUserID: &providerUserID,
			}

			err = u.repo.WithTransaction(
				ctx,
				func(
					ctx context.Context,
					tx *sql.Tx,
				) error {

					return u.userAuthProviderRepo.AddAuthProvider(
						ctx,
						tx,
						authProvider,
					)
				},
			)

			if err != nil {

				return nil, exception.Internal(
					"google login failed",
				)
			}
		}

		// =====================================================
		// CREATE SESSION
		// =====================================================

		loginResponse, err := u.createSession.Execute(
			ctx,
			user,
		)

		if err != nil {
			return nil, err
		}

		return &model.WebResponse[*model.LoginResponse]{
			Message: "login success",
			Data: converter.ToLoginResponse(
				loginResponse.AccessToken,
				loginResponse.RefreshToken,
			),
		}, nil
	}

	// =====================================================
	// PREPARE USER
	// =====================================================

	newUser := &entity.User{
		ID:        uuid.NewString(),
		Email:     googleUser.Email,
		FirstName: googleUser.FirstName,
		LastName:  googleUser.LastName,
		AvatarURL: avatarutil.DefaultAvatar(
			googleUser.FirstName,
			googleUser.LastName,
		),
		IsActive:      true,
		EmailVerified: googleUser.EmailVerified,
	}

	// =====================================================
	// PREPARE GOOGLE PROVIDER
	// =====================================================

	providerUserID := googleUser.ID

	authProvider = &entity.UserAuthProvider{
		ID:             uuid.NewString(),
		UserID:         newUser.ID,
		Provider:       "google",
		ProviderUserID: &providerUserID,
	}

	// =====================================================
	// TRANSACTION
	// =====================================================

	err = u.repo.WithTransaction(
		ctx,
		func(
			ctx context.Context,
			tx *sql.Tx,
		) error {

			if err := u.userRepo.AddUser(
				ctx,
				tx,
				newUser,
			); err != nil {

				return err
			}

			if err := u.userAuthProviderRepo.AddAuthProvider(
				ctx,
				tx,
				authProvider,
			); err != nil {

				return err
			}

			return nil
		},
	)

	if err != nil {

		return nil, exception.Internal(
			"google login failed",
		)
	}

	// =====================================================
	// CREATE SESSION
	// =====================================================

	loginResponse, err := u.createSession.Execute(
		ctx,
		newUser,
	)

	if err != nil {
		return nil, err
	}

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.WebResponse[*model.LoginResponse]{
		Message: "login success",
		Data: converter.ToLoginResponse(
			loginResponse.AccessToken,
			loginResponse.RefreshToken,
		),
	}, nil
}
