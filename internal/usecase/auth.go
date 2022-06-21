package usecase

import (
	"fmt"
	"shape-api/internal/common/apperr"
	"shape-api/internal/common/config"
	"shape-api/internal/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrExpiredAccessToken            = apperr.New("expired access token")
	ErrWrongAccessTokenSigningMethod = apperr.New("wrong access token signing method")
)

func NewAuthUsecase() *AuthUsecase {
	return &AuthUsecase{}
}

type AuthUsecase struct {
}

func (u *AuthUsecase) GenAccessToken(t time.Time, user *entity.User) (string, error) {
	claims := &AccessTokenClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: t.UnixNano() + int64(2*time.Hour),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.App().Secret))
	if err != nil {
		return "", fmt.Errorf("signing token with jwt: %w", err)
	}

	return signedToken, nil
}

type AccessTokenClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (u *AuthUsecase) ValidateAccessToken(t time.Time, accessToken string) (claims *AccessTokenClaims, err error) {
	token, err := jwt.ParseWithClaims(accessToken, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("%w: expected '%v', got '%v'",
				ErrWrongAccessTokenSigningMethod,
				jwt.SigningMethodHS256.Name,
				token.Header["alg"])
		}
		return []byte(config.App().Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse token with claims: %w", err)
	}

	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		if !claims.VerifyExpiresAt(t.UnixNano(), false) {
			return nil, ErrExpiredAccessToken
		}
		return claims, nil
	} else {
		return nil, fmt.Errorf("parse token: %w", err)
	}
}
