package authcheck

import (
	"fmt"
	"shape-api/internal/common/apperr"
	"shape-api/internal/usecase"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ErrEmptyAccessToken   = apperr.New("empty access token")
	ErrInvalidBearerToken = apperr.New("invalid bearer token")
)

func IsAuthenticated(t time.Time, c *gin.Context) error {
	accessToken, err := extractAccessToken(c)
	if err != nil {
		return fmt.Errorf("extract access token: %w", err)
	}
	authUC := usecase.NewAuthUsecase()
	_, err = authUC.ValidateAccessToken(t, accessToken)
	if err != nil {
		return fmt.Errorf("validate access token: %w", err)
	}
	return nil
}

func extractAccessToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken == "" {
		return "", ErrEmptyAccessToken
	}
	bearerTokenParts := strings.Split(bearerToken, " ")
	if len(bearerTokenParts) != 2 || bearerTokenParts[0] != "Bearer" {
		return "", ErrInvalidBearerToken
	}
	return bearerTokenParts[1], nil
}
