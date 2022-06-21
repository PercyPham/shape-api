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

func TestCreateDiamondWithInvalidInput(t *testing.T) {
	diamondRepoMock := repomock.NewDiamondRepoMockBuilder().Build()
	diamondUC := usecase.NewDiamondUsecase(diamondRepoMock)
	input := &usecase.CreateDiamondInput{DiagonalA: -1, DiagonalB: 1}

	_, err := diamondUC.CreateDiamond(input)

	if !errors.Is(err, entity.ErrInvalidDiamond) {
		t.Errorf("expected error '%v', got '%v'", entity.ErrInvalidDiamond, err)
	}
}

func TestCreateDiamond(t *testing.T) {
	diamondRepoMock := repomock.
		NewDiamondRepoMockBuilder().
		WithCreateDiamondMock(func(t *entity.Diamond) (*entity.Diamond, error) {
			created := t
			created.ID = 1
			return created, nil
		}).
		Build()
	diamondUC := usecase.NewDiamondUsecase(diamondRepoMock)
	input := &usecase.CreateDiamondInput{DiagonalA: 3, DiagonalB: 4}

	diamond, err := diamondUC.CreateDiamond(input)

	assert.Nil(t, err)
	assert.NotNil(t, diamond)
}

func TestGetNonExistDiamond(t *testing.T) {
	diamondRepoMock := repomock.
		NewDiamondRepoMockBuilder().
		WithGetDiamondByIDMock(func(_ int64) *entity.Diamond { return nil }).
		Build()
	diamondUC := usecase.NewDiamondUsecase(diamondRepoMock)

	diamond := diamondUC.GetDiamondByID(100)

	assert.Nil(t, diamond)
}

func TestGetDiamond(t *testing.T) {
	diamondRepoMock := repomock.
		NewDiamondRepoMockBuilder().
		WithGetDiamondByIDMock(func(id int64) *entity.Diamond {
			if id == 10 {
				return &entity.Diamond{ID: 10, DiagonalA: 3, DiagonalB: 4}
			}
			return nil
		}).
		Build()
	diamondUC := usecase.NewDiamondUsecase(diamondRepoMock)

	diamond := diamondUC.GetDiamondByID(10)

	if assert.NotNil(t, diamond) {
		assert.Equal(t, diamond.ID, int64(10))
		assert.Equal(t, diamond.DiagonalA, 3.0)
		assert.Equal(t, diamond.DiagonalB, 4.0)
	}
}

func TestUpdateNonExistDiamond(t *testing.T) {
	diamondRepoMock := repomock.
		NewDiamondRepoMockBuilder().
		WithGetDiamondByIDMock(func(_ int64) *entity.Diamond { return nil }).
		Build()
	diamondUC := usecase.NewDiamondUsecase(diamondRepoMock)
	input := genValidUpdateDiamondInput()
	input.ID = 1000

	_, err := diamondUC.UpdateDiamond(input)

	if !errors.Is(err, repo.ErrDiamondNotFound) {
		t.Errorf("expected error '%v', got '%v'", repo.ErrDiamondNotFound, err)
	}
}

func TestUpdateWithInvalidDiamond(t *testing.T) {
	diamondRepoMock := repomock.
		NewDiamondRepoMockBuilder().
		WithGetDiamondByIDMock(func(_ int64) *entity.Diamond {
			return &entity.Diamond{ID: 1, DiagonalA: 3, DiagonalB: 4}
		}).
		Build()
	diamondUC := usecase.NewDiamondUsecase(diamondRepoMock)
	input := genValidUpdateDiamondInput()
	input.DiagonalA = -1000

	_, err := diamondUC.UpdateDiamond(input)

	if !errors.Is(err, entity.ErrInvalidDiamond) {
		t.Errorf("expected error '%v', got '%v'", entity.ErrInvalidDiamond, err)
	}
}

func TestUpdateDiamond(t *testing.T) {
	diamondRepoMock := repomock.
		NewDiamondRepoMockBuilder().
		WithGetDiamondByIDMock(func(id int64) *entity.Diamond {
			if id == 1 {
				return &entity.Diamond{ID: 1, DiagonalA: 4, DiagonalB: 5}
			}
			return nil
		}).
		WithUpdateDiamondMock(func(t *entity.Diamond) (*entity.Diamond, error) {
			if t.ID == 1 {
				return t, nil
			}
			return nil, nil
		}).
		Build()
	diamondUC := usecase.NewDiamondUsecase(diamondRepoMock)
	input := genValidUpdateDiamondInput()

	diamond, err := diamondUC.UpdateDiamond(input)

	if assert.Nil(t, err) {
		assert.Equal(t, diamond.ID, int64(1))
		assert.Equal(t, diamond.DiagonalA, 3.0)
		assert.Equal(t, diamond.DiagonalB, 4.0)
	}
}

func genValidUpdateDiamondInput() *usecase.UpdateDiamondInput {
	return &usecase.UpdateDiamondInput{
		ID:                 1,
		CreateDiamondInput: usecase.CreateDiamondInput{DiagonalA: 3, DiagonalB: 4},
	}
}

func TestDeleteNonExistDiamond(t *testing.T) {
	diamondRepoMock := repomock.
		NewDiamondRepoMockBuilder().
		WithGetDiamondByIDMock(func(_ int64) *entity.Diamond {
			return nil
		}).
		Build()
	diamondUC := usecase.NewDiamondUsecase(diamondRepoMock)

	err := diamondUC.DeleteDiamondByID(1)
	assert.Nil(t, err)
}

func TestDeleteDiamondWithErrorRepo(t *testing.T) {
	expectedErr := errors.New("error")
	diamondRepoMock := repomock.
		NewDiamondRepoMockBuilder().
		WithGetDiamondByIDMock(func(id int64) *entity.Diamond {
			if id == 1 {
				return &entity.Diamond{ID: 1, DiagonalA: 3, DiagonalB: 4}
			}
			return nil
		}).
		WithDeleteDiamondByIDMock(func(i int64) error {
			return expectedErr
		}).
		Build()
	diamondUC := usecase.NewDiamondUsecase(diamondRepoMock)

	err := diamondUC.DeleteDiamondByID(1)
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error '%v', got '%v'", expectedErr, err)
	}
}

func TestDeleteDiamond(t *testing.T) {
	diamondRepoMock := repomock.
		NewDiamondRepoMockBuilder().
		WithGetDiamondByIDMock(func(id int64) *entity.Diamond {
			if id == 1 {
				return &entity.Diamond{ID: 1, DiagonalA: 3, DiagonalB: 4}
			}
			return nil
		}).
		WithDeleteDiamondByIDMock(func(i int64) error { return nil }).
		Build()
	diamondUC := usecase.NewDiamondUsecase(diamondRepoMock)

	err := diamondUC.DeleteDiamondByID(1)
	assert.Nil(t, err)
}
