package auth

import (
	"context"
	"time"

	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"
)

type LogoutUseCase struct {
	*AccessSessionUseCase

	userSessionRepo port.UserSessionRepository
	cache           port.Cache
}

func NewLogoutUseCase(
	base *AccessSessionUseCase,
	userSessionRepo port.UserSessionRepository,
	cache port.Cache,
) *LogoutUseCase {

	return &LogoutUseCase{
		AccessSessionUseCase: base,
		userSessionRepo:      userSessionRepo,
		cache:                cache,
	}
}

func (u *LogoutUseCase) Execute(
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
	// FIND SESSION
	// =====================================================

	userSession, err := u.userSessionRepo.FindSessionByID(
		ctx,
		session.SessionID,
	)

	if err != nil {

		return nil, exception.Internal(
			"logout failed",
		)
	}

	if userSession == nil {

		return nil, exception.Unauthorized(
			"session not found",
		)
	}

	// =====================================================
	// REVOKE SESSION
	// =====================================================

	err = u.userSessionRepo.RevokeByID(
		ctx,
		session.SessionID,
		time.Now().Unix(),
	)

	if err != nil {

		return nil, exception.Internal(
			"logout failed",
		)
	}

	// =====================================================
	// DELETE CACHE
	// =====================================================

	_ = u.cache.Delete(
		ctx,
		cachekey.Session(
			session.SessionID,
		),
	)

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.WebResponse[any]{
		Message: "logout success",
	}, nil
}
