package usecase

import (
	"fmt"
	"shape-api/internal/common/apperr"
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

var (
	ErrInvalidDiamondInput = apperr.New("invalid diamond input")
)

func NewDiamondUsecase(tr repo.Diamond) *DiamondUsecase {
	return &DiamondUsecase{tr}
}

type DiamondUsecase struct {
	diamondRepo repo.Diamond
}

func (u *DiamondUsecase) CreateDiamond(input *CreateDiamondInput) (*entity.Diamond, error) {
	if err := input.validate(); err != nil {
		return nil, fmt.Errorf("validate input: %w", err)
	}

	diamond, err := entity.NewDiamond(input.DiagonalA, input.DiagonalB)
	if err != nil {
		return nil, fmt.Errorf("new diamond: %w", err)
	}

	diamond, err = u.diamondRepo.CreateDiamond(diamond)
	if err != nil {
		return nil, fmt.Errorf("repo create diamond: %w", err)
	}

	return diamond, nil
}

type CreateDiamondInput struct {
	DiagonalA float64 `json:"diagonalA"`
	DiagonalB float64 `json:"diagonalB"`
}

func (i *CreateDiamondInput) validate() error {
	if i.DiagonalA == 0 {
		return fmt.Errorf("%w: diagonalA must be greater than 0", ErrInvalidDiamondInput)
	}
	if i.DiagonalB == 0 {
		return fmt.Errorf("%w: diagonalB must be greater than 0", ErrInvalidDiamondInput)
	}
	return nil
}

func (u *DiamondUsecase) GetDiamondByID(id int64) *entity.Diamond {
	return u.diamondRepo.GetDiamondByID(id)
}

func (u *DiamondUsecase) UpdateDiamond(input *UpdateDiamondInput) (*entity.Diamond, error) {
	if err := input.validate(); err != nil {
		return nil, fmt.Errorf("validate input: %w", err)
	}

	found := u.diamondRepo.GetDiamondByID(input.ID)
	if found == nil {
		return nil, repo.ErrDiamondNotFound
	}

	diamond, err := entity.NewDiamond(input.DiagonalA, input.DiagonalB)
	if err != nil {
		return nil, fmt.Errorf("new diamond: %w", err)
	}
	diamond.ID = input.ID

	diamond, err = u.diamondRepo.UpdateDiamond(diamond)
	if err != nil {
		return nil, fmt.Errorf("repo update diamond: %w", err)
	}

	return diamond, nil
}

type UpdateDiamondInput struct {
	ID int64
	CreateDiamondInput
}

func (u *DiamondUsecase) DeleteDiamondByID(id int64) error {
	found := u.diamondRepo.GetDiamondByID(id)
	if found == nil {
		return nil
	}

	err := u.diamondRepo.DeleteDiamondByID(id)
	if err != nil {
		return fmt.Errorf("repo delete diamond: %w", err)
	}

	return nil
}
