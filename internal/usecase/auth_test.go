package usecase_test

import (
	"errors"
	"shape-api/internal/entity"
	"shape-api/internal/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenAccessToken(t *testing.T) {
	testedTime := time.Now()
	authUC := usecase.NewAuthUsecase()
	user, _ := entity.NewUser(testedTime, "percy", "password")

	_, err := authUC.GenAccessToken(testedTime, user)
	assert.Nil(t, err)
}

func TestValidateExpiredAccessToken(t *testing.T) {
	testedTime := time.Now()
	authUC := usecase.NewAuthUsecase()
	user, _ := entity.NewUser(testedTime, "percy", "password")

	expiredTime := testedTime.
		Add(2 * time.Hour).
		Add(1 * time.Second)

	accessToken, err := authUC.GenAccessToken(testedTime, user)
	if assert.Nil(t, err) {
		_, err := authUC.ValidateAccessToken(expiredTime, accessToken)
		if !errors.Is(err, usecase.ErrExpiredAccessToken) {
			t.Errorf("expected error '%v', got '%v'", usecase.ErrExpiredAccessToken, err)
		}
	}
}

func TestValidateValidAccessToken(t *testing.T) {
	testedTime := time.Now()
	authUC := usecase.NewAuthUsecase()
	user, _ := entity.NewUser(testedTime, "percy", "password")

	validTime := testedTime.
		Add(1 * time.Second)

	accessToken, err := authUC.GenAccessToken(testedTime, user)
	if assert.Nil(t, err) {
		claims, err := authUC.ValidateAccessToken(validTime, accessToken)
		assert.Nil(t, err)
		assert.Equal(t, claims.Username, "percy")
	}
}
