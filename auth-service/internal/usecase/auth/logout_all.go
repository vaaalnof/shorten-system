package auth

import (
	"context"
	"time"

	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"
)

type LogoutAllUseCase struct {
	*AccessSessionUseCase

	userSessionRepo port.UserSessionRepository
	cache           port.Cache
}

func NewLogoutAllUseCase(
	base *AccessSessionUseCase,
	userSessionRepo port.UserSessionRepository,
	cache port.Cache,
) *LogoutAllUseCase {

	return &LogoutAllUseCase{
		AccessSessionUseCase: base,
		userSessionRepo:      userSessionRepo,
		cache:                cache,
	}
}

func (u *LogoutAllUseCase) Execute(
	ctx context.Context,
) (*model.WebResponse[any], error) {

	// =====================================================
	// AUTH SESSION
	// =====================================================

	session, err := u.GetSession(
		ctx,
	)

	if err != nil {
		return nil, err
	}

	// =====================================================
	// FIND USER SESSIONS
	// =====================================================

	sessions, err := u.userSessionRepo.FindSessionByUserID(
		ctx,
		session.UserID,
	)

	if err != nil {

		return nil, exception.Internal(
			"logout all failed",
		)
	}

	// =====================================================
	// REVOKE ALL SESSIONS
	// =====================================================

	err = u.userSessionRepo.RevokeByUserID(
		ctx,
		session.UserID,
		time.Now().Unix(),
	)

	if err != nil {

		return nil, exception.Internal(
			"logout all failed",
		)
	}

	// =====================================================
	// DELETE SESSION CACHE
	// =====================================================

	for _, userSession := range sessions {

		_ = u.cache.Delete(
			ctx,
			cachekey.Session(
				userSession.ID,
			),
		)
	}

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.WebResponse[any]{
		Message: "logout all success",
	}, nil
}
