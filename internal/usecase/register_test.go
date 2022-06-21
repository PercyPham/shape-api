package usecase_test

import (
	"errors"
	"shape-api/internal/entity"
	"shape-api/internal/repo"
	"shape-api/internal/usecase"
	"shape-api/test/repomock"
	"testing"
	"time"
)

func TestRegisterWithInvalidInput(t *testing.T) {
	userRepoMock := repomock.NewUserRepoMockBuilder().Build()

	registerUC := usecase.NewRegisterUsecase(userRepoMock)
	err := registerUC.Register(time.Now(), &usecase.RegisterInput{})

	if !errors.Is(err, entity.ErrEmptyUsername) {
		t.Errorf("expected error '%v', got '%v'", entity.ErrEmptyUsername, err)
	}
}

func TestRegisterWithDuplicateUsername(t *testing.T) {
	userRepoMock := repomock.
		NewUserRepoMockBuilder().
		WithCreateUserMock(func(u *entity.User) error {
			return repo.ErrDuplicateUsername
		}).
		Build()

	registerUC := usecase.NewRegisterUsecase(userRepoMock)
	err := registerUC.Register(time.Now(), createValidRegisterInput())

	if !errors.Is(err, repo.ErrDuplicateUsername) {
		t.Errorf("expected error '%v', got '%v'", repo.ErrDuplicateUsername, err)
	}
}

func TestRegisterWithValidInput(t *testing.T) {
	userRepoMock := repomock.
		NewUserRepoMockBuilder().
		WithCreateUserMock(func(u *entity.User) error { return nil }).
		Build()

	registerUC := usecase.NewRegisterUsecase(userRepoMock)
	err := registerUC.Register(time.Now(), createValidRegisterInput())

	if err != nil {
		t.Errorf("expected nil, got '%v'", err)
	}
}

func createValidRegisterInput() *usecase.RegisterInput {
	return &usecase.RegisterInput{
		Username: "username",
		Password: "password",
	}
}
