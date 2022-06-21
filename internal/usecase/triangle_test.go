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

func TestCreateTriangleWithInvalidInput(t *testing.T) {
	triangleRepoMock := repomock.NewTriangleRepoMockBuilder().Build()
	triangleUC := usecase.NewTriangleUsecase(triangleRepoMock)
	input := &usecase.CreateTriangleInput{A: -1, B: 1, C: 1}

	_, err := triangleUC.CreateTriangle(input)

	if !errors.Is(err, entity.ErrInvalidTriangle) {
		t.Errorf("expected error '%v', got '%v'", entity.ErrInvalidTriangle, err)
	}
}

func TestCreateTriangle(t *testing.T) {
	triangleRepoMock := repomock.
		NewTriangleRepoMockBuilder().
		WithCreateTriangleMock(func(t *entity.Triangle) (*entity.Triangle, error) {
			return &entity.Triangle{ID: 1, A: 3, B: 4, C: 5}, nil
		}).
		Build()
	triangleUC := usecase.NewTriangleUsecase(triangleRepoMock)
	input := &usecase.CreateTriangleInput{A: 3, B: 4, C: 5}

	triangle, err := triangleUC.CreateTriangle(input)

	assert.Nil(t, err)
	assert.NotNil(t, triangle)
}

func TestGetNonExistTriangle(t *testing.T) {
	triangleRepoMock := repomock.
		NewTriangleRepoMockBuilder().
		WithGetTriangleByIDMock(func(_ int64) *entity.Triangle { return nil }).
		Build()
	triangleUC := usecase.NewTriangleUsecase(triangleRepoMock)

	triangle := triangleUC.GetTriangleByID(100)

	assert.Nil(t, triangle)
}

func TestGetTriangle(t *testing.T) {
	triangleRepoMock := repomock.
		NewTriangleRepoMockBuilder().
		WithGetTriangleByIDMock(func(id int64) *entity.Triangle {
			if id == 10 {
				return &entity.Triangle{ID: 10, A: 3, B: 4, C: 5}
			}
			return nil
		}).
		Build()
	triangleUC := usecase.NewTriangleUsecase(triangleRepoMock)

	triangle := triangleUC.GetTriangleByID(10)

	if assert.NotNil(t, triangle) {
		assert.Equal(t, triangle.ID, int64(10))
		assert.Equal(t, triangle.A, 3.0)
		assert.Equal(t, triangle.B, 4.0)
		assert.Equal(t, triangle.C, 5.0)
	}
}

func TestUpdateNonExistTriangle(t *testing.T) {
	triangleRepoMock := repomock.
		NewTriangleRepoMockBuilder().
		WithGetTriangleByIDMock(func(_ int64) *entity.Triangle { return nil }).
		Build()
	triangleUC := usecase.NewTriangleUsecase(triangleRepoMock)
	input := genValidUpdateTriangleInput()
	input.ID = 1000

	_, err := triangleUC.UpdateTriangle(input)

	if !errors.Is(err, repo.ErrTriangleNotFound) {
		t.Errorf("expected error '%v', got '%v'", repo.ErrTriangleNotFound, err)
	}
}

func TestUpdateWithInvalidTriangle(t *testing.T) {
	triangleRepoMock := repomock.
		NewTriangleRepoMockBuilder().
		WithGetTriangleByIDMock(func(_ int64) *entity.Triangle {
			return &entity.Triangle{ID: 1, A: 3, B: 4, C: 5}
		}).
		Build()
	triangleUC := usecase.NewTriangleUsecase(triangleRepoMock)
	input := genValidUpdateTriangleInput()
	input.C = 1000

	_, err := triangleUC.UpdateTriangle(input)

	if !errors.Is(err, entity.ErrInvalidTriangle) {
		t.Errorf("expected error '%v', got '%v'", entity.ErrInvalidTriangle, err)
	}
}

func TestUpdateTriangle(t *testing.T) {
	triangleRepoMock := repomock.
		NewTriangleRepoMockBuilder().
		WithGetTriangleByIDMock(func(id int64) *entity.Triangle {
			if id == 1 {
				return &entity.Triangle{ID: 1, A: 2, B: 2, C: 3}
			}
			return nil
		}).
		WithUpdateTriangleMock(func(t *entity.Triangle) (*entity.Triangle, error) {
			if t.ID == 1 {
				return t, nil
			}
			return nil, nil
		}).
		Build()
	triangleUC := usecase.NewTriangleUsecase(triangleRepoMock)
	input := genValidUpdateTriangleInput()

	triangle, err := triangleUC.UpdateTriangle(input)

	if assert.Nil(t, err) {
		assert.Equal(t, triangle.ID, int64(1))
		assert.Equal(t, triangle.A, 3.0)
		assert.Equal(t, triangle.B, 4.0)
		assert.Equal(t, triangle.C, 5.0)
	}
}

func genValidUpdateTriangleInput() *usecase.UpdateTriangleInput {
	return &usecase.UpdateTriangleInput{
		ID:                  1,
		CreateTriangleInput: usecase.CreateTriangleInput{A: 3, B: 4, C: 5},
	}
}

func TestDeleteNonExistTriangle(t *testing.T) {
	triangleRepoMock := repomock.
		NewTriangleRepoMockBuilder().
		WithGetTriangleByIDMock(func(_ int64) *entity.Triangle {
			return nil
		}).
		Build()
	triangleUC := usecase.NewTriangleUsecase(triangleRepoMock)

	err := triangleUC.DeleteTriangleByID(1)
	assert.Nil(t, err)
}

func TestDeleteTriangleWithErrorRepo(t *testing.T) {
	expectedErr := errors.New("error")
	triangleRepoMock := repomock.
		NewTriangleRepoMockBuilder().
		WithGetTriangleByIDMock(func(id int64) *entity.Triangle {
			if id == 1 {
				return &entity.Triangle{ID: 1, A: 2, B: 2, C: 3}
			}
			return nil
		}).
		WithDeleteTriangleByIDMock(func(i int64) error {
			return expectedErr
		}).
		Build()
	triangleUC := usecase.NewTriangleUsecase(triangleRepoMock)

	err := triangleUC.DeleteTriangleByID(1)
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error '%v', got '%v'", expectedErr, err)
	}
}

func TestDeleteTriangle(t *testing.T) {
	triangleRepoMock := repomock.
		NewTriangleRepoMockBuilder().
		WithGetTriangleByIDMock(func(id int64) *entity.Triangle {
			if id == 1 {
				return &entity.Triangle{ID: 1, A: 2, B: 2, C: 3}
			}
			return nil
		}).
		WithDeleteTriangleByIDMock(func(i int64) error { return nil }).
		Build()
	triangleUC := usecase.NewTriangleUsecase(triangleRepoMock)

	err := triangleUC.DeleteTriangleByID(1)
	assert.Nil(t, err)
}
