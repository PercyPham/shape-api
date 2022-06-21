package usecase

import (
	"fmt"
	"shape-api/internal/common/apperr"
	"shape-api/internal/entity"
	"shape-api/internal/repo"
	"time"
)

var (
	ErrInvalidUsernameOrPassword = apperr.New("invalid username or password")
)

func NewLoginUsecase(ur repo.User) *LoginUsecase {
	return &LoginUsecase{ur}
}

type LoginUsecase struct {
	userRepo repo.User
}

func (u *LoginUsecase) Login(t time.Time, input *LoginInput) (accessToken string, err error) {
	if input.Username == "" {
		return "", entity.ErrEmptyUsername
	}
	if input.Password == "" {
		return "", entity.ErrEmptyPassword
	}

	user := u.userRepo.GetUserByUsername(input.Username)
	if user == nil || !user.IsCorrectPassword(input.Password) {
		return "", ErrInvalidUsernameOrPassword
	}

	accessToken, err = NewAuthUsecase().GenAccessToken(t, user)
	if err != nil {
		return "", fmt.Errorf("gen access token: %w", err)
	}

	return accessToken, nil
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
