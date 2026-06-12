package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenType  = "access"
	RefreshTokenType = "refresh"
)

type JWTService interface {
	GenerateAccessToken(
		userID string,
		sessionID string,
	) (string, error)

	GenerateRefreshToken(
		userID string,
		sessionID string,
	) (string, error)

	AccessTokenTTL() time.Duration

	RefreshTokenTTL() time.Duration

	ParseToken(
		tokenString string,
	) (jwt.MapClaims, error)
}

type jwtService struct {
	secretKey       string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewJWTService(
	secretKey string,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) JWTService {

	return &jwtService{
		secretKey:       secretKey,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

// =====================================================
// ACCESS TOKEN
// =====================================================

func (j *jwtService) GenerateAccessToken(
	userID string,
	sessionID string,
) (string, error) {

	return j.generateToken(
		userID,
		sessionID,
		AccessTokenType,
		j.accessTokenTTL,
	)
}

func (j *jwtService) AccessTokenTTL() time.Duration {
	return j.accessTokenTTL
}

// =====================================================
// REFRESH TOKEN
// =====================================================

func (j *jwtService) GenerateRefreshToken(
	userID string,
	sessionID string,
) (string, error) {

	return j.generateToken(
		userID,
		sessionID,
		RefreshTokenType,
		j.refreshTokenTTL,
	)
}

func (j *jwtService) RefreshTokenTTL() time.Duration {
	return j.refreshTokenTTL
}

// =====================================================
// INTERNAL GENERATOR
// =====================================================

func (j *jwtService) generateToken(
	userID string,
	sessionID string,
	tokenType string,
	duration time.Duration,
) (string, error) {

	now := time.Now()

	claims := jwt.MapClaims{
		"sub": userID,
		"sid": sessionID,
		"typ": tokenType,
		"iat": now.Unix(),
		"exp": now.Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(j.secretKey),
	)
}

// =====================================================
// PARSE TOKEN
// =====================================================

func (j *jwtService) ParseToken(
	tokenString string,
) (jwt.MapClaims, error) {

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New(
					"invalid signing method",
				)
			}

			return []byte(
				j.secretKey,
			), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New(
			"invalid token",
		)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New(
			"invalid claims",
		)
	}

	return claims, nil
}
