package usecase_test

import (
	"errors"
	"shape-api/internal/entity"
	"shape-api/internal/usecase"
	"shape-api/test/repomock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoginWithEmptyUsername(t *testing.T) {
	userRepoMock := repomock.NewUserRepoMockBuilder().Build()
	loginUC := usecase.NewLoginUsecase(userRepoMock)

	accessToken, err := loginUC.Login(time.Now(), &usecase.LoginInput{Password: "password"})
	if assert.Equal(t, accessToken, "") {
		if !errors.Is(err, entity.ErrEmptyUsername) {
			t.Errorf("expected error '%v', got '%v'", entity.ErrEmptyUsername, err)
		}
	}
}

func TestLoginWithEmptyPassword(t *testing.T) {
	userRepoMock := repomock.NewUserRepoMockBuilder().Build()
	loginUC := usecase.NewLoginUsecase(userRepoMock)

	accessToken, err := loginUC.Login(time.Now(), &usecase.LoginInput{Username: "username"})
	if assert.Equal(t, accessToken, "") {
		if !errors.Is(err, entity.ErrEmptyPassword) {
			t.Errorf("expected error '%v', got '%v'", entity.ErrEmptyPassword, err)
		}
	}
}

func TestLoginWithInvalidUsername(t *testing.T) {
	userRepoMock := repomock.
		NewUserRepoMockBuilder().
		WithGetUserByUsernameMock(func(_ string) *entity.User {
			return nil
		}).
		Build()
	loginUC := usecase.NewLoginUsecase(userRepoMock)

	accessToken, err := loginUC.Login(time.Now(), &usecase.LoginInput{
		Username: "not-exist",
		Password: "password",
	})
	if assert.Equal(t, accessToken, "") {
		if !errors.Is(err, usecase.ErrInvalidUsernameOrPassword) {
			t.Errorf("expected error '%v', got '%v'", usecase.ErrInvalidUsernameOrPassword, err)
		}
	}
}

func TestLoginWithInvalidPassword(t *testing.T) {
	userRepoMock := repomock.
		NewUserRepoMockBuilder().
		WithGetUserByUsernameMock(func(username string) *entity.User {
			user, _ := entity.NewUser(time.Now(), username, "password")
			return user
		}).
		Build()
	loginUC := usecase.NewLoginUsecase(userRepoMock)

	accessToken, err := loginUC.Login(time.Now(), &usecase.LoginInput{
		Username: "username",
		Password: "invalid",
	})
	if assert.Equal(t, accessToken, "") {
		if !errors.Is(err, usecase.ErrInvalidUsernameOrPassword) {
			t.Errorf("expected error '%v', got '%v'", usecase.ErrInvalidUsernameOrPassword, err)
		}
	}
}

func TestLoginWithValidInfo(t *testing.T) {
	userRepoMock := repomock.
		NewUserRepoMockBuilder().
		WithGetUserByUsernameMock(func(username string) *entity.User {
			user, _ := entity.NewUser(time.Now(), username, "password")
			return user
		}).
		Build()
	loginUC := usecase.NewLoginUsecase(userRepoMock)

	accessToken, err := loginUC.Login(time.Now(), &usecase.LoginInput{
		Username: "username",
		Password: "password",
	})
	if assert.Nil(t, err) {
		claims, err := usecase.NewAuthUsecase().ValidateAccessToken(time.Now(), accessToken)
		assert.Nil(t, err)
		assert.Equal(t, claims.Username, "username")
	}
}
