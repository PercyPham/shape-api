package usecase

import (
	"fmt"
	"shape-api/internal/common/apperr"
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

var (
	ErrInvalidTriangleInput = apperr.New("invalid triangle input")
)

func NewTriangleUsecase(tr repo.Triangle) *TriangleUsecase {
	return &TriangleUsecase{tr}
}

type TriangleUsecase struct {
	triangleRepo repo.Triangle
}

func (u *TriangleUsecase) CreateTriangle(input *CreateTriangleInput) (*entity.Triangle, error) {
	if err := input.validate(); err != nil {
		return nil, fmt.Errorf("validate input: %w", err)
	}

	triangle, err := entity.NewTriangle(input.A, input.B, input.C)
	if err != nil {
		return nil, fmt.Errorf("new triangle: %w", err)
	}

	triangle, err = u.triangleRepo.CreateTriangle(triangle)
	if err != nil {
		return nil, fmt.Errorf("repo create triangle: %w", err)
	}

	return triangle, nil
}

type CreateTriangleInput struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
}

func (i *CreateTriangleInput) validate() error {
	if i.A == 0 {
		return fmt.Errorf("%w: a must be greater than 0", ErrInvalidTriangleInput)
	}
	if i.B == 0 {
		return fmt.Errorf("%w: b must be greater than 0", ErrInvalidTriangleInput)
	}
	if i.C == 0 {
		return fmt.Errorf("%w: c must be greater than 0", ErrInvalidTriangleInput)
	}
	return nil
}

func (u *TriangleUsecase) GetTriangleByID(id int64) *entity.Triangle {
	return u.triangleRepo.GetTriangleByID(id)
}

func (u *TriangleUsecase) UpdateTriangle(input *UpdateTriangleInput) (*entity.Triangle, error) {
	if err := input.validate(); err != nil {
		return nil, fmt.Errorf("validate input: %w", err)
	}

	found := u.triangleRepo.GetTriangleByID(input.ID)
	if found == nil {
		return nil, repo.ErrTriangleNotFound
	}

	triangle, err := entity.NewTriangle(input.A, input.B, input.C)
	if err != nil {
		return nil, fmt.Errorf("new triangle: %w", err)
	}
	triangle.ID = input.ID

	triangle, err = u.triangleRepo.UpdateTriangle(triangle)
	if err != nil {
		return nil, fmt.Errorf("repo update triangle: %w", err)
	}

	return triangle, nil
}

type UpdateTriangleInput struct {
	ID int64
	CreateTriangleInput
}

func (u *TriangleUsecase) DeleteTriangleByID(id int64) error {
	found := u.triangleRepo.GetTriangleByID(id)
	if found == nil {
		return nil
	}

	err := u.triangleRepo.DeleteTriangleByID(id)
	if err != nil {
		return fmt.Errorf("repo delete triangle: %w", err)
	}

	return nil
}
