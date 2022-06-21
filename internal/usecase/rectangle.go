package usecase

import (
	"fmt"
	"shape-api/internal/common/apperr"
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

var (
	ErrInvalidRectangleInput = apperr.New("invalid rectangle input")
)

func NewRectangleUsecase(tr repo.Rectangle) *RectangleUsecase {
	return &RectangleUsecase{tr}
}

type RectangleUsecase struct {
	rectangleRepo repo.Rectangle
}

func (u *RectangleUsecase) CreateRectangle(input *CreateRectangleInput) (*entity.Rectangle, error) {
	if err := input.validate(); err != nil {
		return nil, fmt.Errorf("validate input: %w", err)
	}

	rectangle, err := entity.NewRectangle(input.L, input.W)
	if err != nil {
		return nil, fmt.Errorf("new rectangle: %w", err)
	}

	rectangle, err = u.rectangleRepo.CreateRectangle(rectangle)
	if err != nil {
		return nil, fmt.Errorf("repo create rectangle: %w", err)
	}

	return rectangle, nil
}

type CreateRectangleInput struct {
	L float64 `json:"l"`
	W float64 `json:"w"`
}

func (i *CreateRectangleInput) validate() error {
	if i.L == 0 {
		return fmt.Errorf("%w: l must be greater than 0", ErrInvalidRectangleInput)
	}
	if i.W == 0 {
		return fmt.Errorf("%w: w must be greater than 0", ErrInvalidRectangleInput)
	}
	return nil
}

func (u *RectangleUsecase) GetRectangleByID(id int64) *entity.Rectangle {
	return u.rectangleRepo.GetRectangleByID(id)
}

func (u *RectangleUsecase) UpdateRectangle(input *UpdateRectangleInput) (*entity.Rectangle, error) {
	if err := input.validate(); err != nil {
		return nil, fmt.Errorf("validate input: %w", err)
	}

	found := u.rectangleRepo.GetRectangleByID(input.ID)
	if found == nil {
		return nil, repo.ErrRectangleNotFound
	}

	rectangle, err := entity.NewRectangle(input.L, input.W)
	if err != nil {
		return nil, fmt.Errorf("new rectangle: %w", err)
	}
	rectangle.ID = input.ID

	rectangle, err = u.rectangleRepo.UpdateRectangle(rectangle)
	if err != nil {
		return nil, fmt.Errorf("repo update rectangle: %w", err)
	}

	return rectangle, nil
}

type UpdateRectangleInput struct {
	ID int64
	CreateRectangleInput
}

func (u *RectangleUsecase) DeleteRectangleByID(id int64) error {
	found := u.rectangleRepo.GetRectangleByID(id)
	if found == nil {
		return nil
	}

	err := u.rectangleRepo.DeleteRectangleByID(id)
	if err != nil {
		return fmt.Errorf("repo delete rectangle: %w", err)
	}

	return nil
}
