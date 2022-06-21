package usecase

import (
	"fmt"
	"shape-api/internal/entity"
	"shape-api/internal/repo"
	"time"
)

func NewRegisterUsecase(ur repo.User) *RegisterUsecase {
	return &RegisterUsecase{ur}
}

type RegisterUsecase struct {
	userRepo repo.User
}

func (u *RegisterUsecase) Register(t time.Time, input *RegisterInput) error {
	user, err := entity.NewUser(t, input.Username, input.Password)
	if err != nil {
		return fmt.Errorf("new user: %w", err)
	}

	if err := u.userRepo.CreateUser(user); err != nil {
		return fmt.Errorf("repo create user: %w", err)
	}

	return nil
}

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
