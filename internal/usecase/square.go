package usecase

import (
	"fmt"
	"shape-api/internal/common/apperr"
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

var (
	ErrInvalidSquareInput = apperr.New("invalid square input")
)

func NewSquareUsecase(tr repo.Square) *SquareUsecase {
	return &SquareUsecase{tr}
}

type SquareUsecase struct {
	squareRepo repo.Square
}

func (u *SquareUsecase) CreateSquare(input *CreateSquareInput) (*entity.Square, error) {
	if err := input.validate(); err != nil {
		return nil, fmt.Errorf("validate input: %w", err)
	}

	square, err := entity.NewSquare(input.A)
	if err != nil {
		return nil, fmt.Errorf("new square: %w", err)
	}

	square, err = u.squareRepo.CreateSquare(square)
	if err != nil {
		return nil, fmt.Errorf("repo create square: %w", err)
	}

	return square, nil
}

type CreateSquareInput struct {
	A float64 `json:"a"`
}

func (i *CreateSquareInput) validate() error {
	if i.A == 0 {
		return fmt.Errorf("%w: a must be greater than 0", ErrInvalidSquareInput)
	}
	return nil
}

func (u *SquareUsecase) GetSquareByID(id int64) *entity.Square {
	return u.squareRepo.GetSquareByID(id)
}

func (u *SquareUsecase) UpdateSquare(input *UpdateSquareInput) (*entity.Square, error) {
	if err := input.validate(); err != nil {
		return nil, fmt.Errorf("validate input: %w", err)
	}

	found := u.squareRepo.GetSquareByID(input.ID)
	if found == nil {
		return nil, repo.ErrSquareNotFound
	}

	square, err := entity.NewSquare(input.A)
	if err != nil {
		return nil, fmt.Errorf("new square: %w", err)
	}
	square.ID = input.ID

	square, err = u.squareRepo.UpdateSquare(square)
	if err != nil {
		return nil, fmt.Errorf("repo update square: %w", err)
	}

	return square, nil
}

type UpdateSquareInput struct {
	ID int64
	CreateSquareInput
}

func (u *SquareUsecase) DeleteSquareByID(id int64) error {
	found := u.squareRepo.GetSquareByID(id)
	if found == nil {
		return nil
	}

	err := u.squareRepo.DeleteSquareByID(id)
	if err != nil {
		return fmt.Errorf("repo delete square: %w", err)
	}

	return nil
}
