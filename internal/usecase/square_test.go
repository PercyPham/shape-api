package usecase_test

import (
	"errors"
	"shape-api/internal/entity"
	"shape-api/internal/repo"
	"shape-api/internal/usecase"
	"shape-api/test/repomock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSquareWithInvalidInput(t *testing.T) {
	squareRepoMock := repomock.NewSquareRepoMockBuilder().Build()
	squareUC := usecase.NewSquareUsecase(squareRepoMock)
	input := &usecase.CreateSquareInput{A: -1}

	_, err := squareUC.CreateSquare(input)

	if !errors.Is(err, entity.ErrInvalidSquare) {
		t.Errorf("expected error '%v', got '%v'", entity.ErrInvalidSquare, err)
	}
}

func TestCreateSquare(t *testing.T) {
	squareRepoMock := repomock.
		NewSquareRepoMockBuilder().
		WithCreateSquareMock(func(t *entity.Square) (*entity.Square, error) {
			return &entity.Square{ID: 1, A: 3}, nil
		}).
		Build()
	squareUC := usecase.NewSquareUsecase(squareRepoMock)
	input := &usecase.CreateSquareInput{A: 3}

	square, err := squareUC.CreateSquare(input)

	assert.Nil(t, err)
	assert.NotNil(t, square)
}

func TestGetNonExistSquare(t *testing.T) {
	squareRepoMock := repomock.
		NewSquareRepoMockBuilder().
		WithGetSquareByIDMock(func(_ int64) *entity.Square { return nil }).
		Build()
	squareUC := usecase.NewSquareUsecase(squareRepoMock)

	square := squareUC.GetSquareByID(100)

	assert.Nil(t, square)
}

func TestGetSquare(t *testing.T) {
	squareRepoMock := repomock.
		NewSquareRepoMockBuilder().
		WithGetSquareByIDMock(func(id int64) *entity.Square {
			if id == 10 {
				return &entity.Square{ID: 10, A: 3}
			}
			return nil
		}).
		Build()
	squareUC := usecase.NewSquareUsecase(squareRepoMock)

	square := squareUC.GetSquareByID(10)

	if assert.NotNil(t, square) {
		assert.Equal(t, square.ID, int64(10))
		assert.Equal(t, square.A, 3.0)
	}
}

func TestUpdateNonExistSquare(t *testing.T) {
	squareRepoMock := repomock.
		NewSquareRepoMockBuilder().
		WithGetSquareByIDMock(func(_ int64) *entity.Square { return nil }).
		Build()
	squareUC := usecase.NewSquareUsecase(squareRepoMock)
	input := genValidUpdateSquareInput()
	input.ID = 1000

	_, err := squareUC.UpdateSquare(input)

	if !errors.Is(err, repo.ErrSquareNotFound) {
		t.Errorf("expected error '%v', got '%v'", repo.ErrSquareNotFound, err)
	}
}

func TestUpdateWithInvalidSquare(t *testing.T) {
	squareRepoMock := repomock.
		NewSquareRepoMockBuilder().
		WithGetSquareByIDMock(func(_ int64) *entity.Square {
			return &entity.Square{ID: 1, A: 3}
		}).
		Build()
	squareUC := usecase.NewSquareUsecase(squareRepoMock)
	input := genValidUpdateSquareInput()
	input.A = -1000

	_, err := squareUC.UpdateSquare(input)

	if !errors.Is(err, entity.ErrInvalidSquare) {
		t.Errorf("expected error '%v', got '%v'", entity.ErrInvalidSquare, err)
	}
}

func TestUpdateSquare(t *testing.T) {
	squareRepoMock := repomock.
		NewSquareRepoMockBuilder().
		WithGetSquareByIDMock(func(id int64) *entity.Square {
			if id == 1 {
				return &entity.Square{ID: 1, A: 5}
			}
			return nil
		}).
		WithUpdateSquareMock(func(t *entity.Square) (*entity.Square, error) {
			if t.ID == 1 {
				return t, nil
			}
			return nil, nil
		}).
		Build()
	squareUC := usecase.NewSquareUsecase(squareRepoMock)
	input := genValidUpdateSquareInput()

	square, err := squareUC.UpdateSquare(input)

	if assert.Nil(t, err) {
		assert.Equal(t, square.ID, int64(1))
		assert.Equal(t, square.A, 3.0)
	}
}

func genValidUpdateSquareInput() *usecase.UpdateSquareInput {
	return &usecase.UpdateSquareInput{
		ID:                1,
		CreateSquareInput: usecase.CreateSquareInput{A: 3},
	}
}

func TestDeleteNonExistSquare(t *testing.T) {
	squareRepoMock := repomock.
		NewSquareRepoMockBuilder().
		WithGetSquareByIDMock(func(_ int64) *entity.Square {
			return nil
		}).
		Build()
	squareUC := usecase.NewSquareUsecase(squareRepoMock)

	err := squareUC.DeleteSquareByID(1)
	assert.Nil(t, err)
}

func TestDeleteSquareWithErrorRepo(t *testing.T) {
	expectedErr := errors.New("error")
	squareRepoMock := repomock.
		NewSquareRepoMockBuilder().
		WithGetSquareByIDMock(func(id int64) *entity.Square {
			if id == 1 {
				return &entity.Square{ID: 1, A: 3}
			}
			return nil
		}).
		WithDeleteSquareByIDMock(func(i int64) error {
			return expectedErr
		}).
		Build()
	squareUC := usecase.NewSquareUsecase(squareRepoMock)

	err := squareUC.DeleteSquareByID(1)
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error '%v', got '%v'", expectedErr, err)
	}
}

func TestDeleteSquare(t *testing.T) {
	squareRepoMock := repomock.
		NewSquareRepoMockBuilder().
		WithGetSquareByIDMock(func(id int64) *entity.Square {
			if id == 1 {
				return &entity.Square{ID: 1, A: 3}
			}
			return nil
		}).
		WithDeleteSquareByIDMock(func(i int64) error { return nil }).
		Build()
	squareUC := usecase.NewSquareUsecase(squareRepoMock)

	err := squareUC.DeleteSquareByID(1)
	assert.Nil(t, err)
}
