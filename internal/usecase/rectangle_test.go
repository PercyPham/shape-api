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

func TestCreateRectangleWithInvalidInput(t *testing.T) {
	rectangleRepoMock := repomock.NewRectangleRepoMockBuilder().Build()
	rectangleUC := usecase.NewRectangleUsecase(rectangleRepoMock)
	input := &usecase.CreateRectangleInput{L: -1, W: 1}

	_, err := rectangleUC.CreateRectangle(input)

	if !errors.Is(err, entity.ErrInvalidRectangle) {
		t.Errorf("expected error '%v', got '%v'", entity.ErrInvalidRectangle, err)
	}
}

func TestCreateRectangle(t *testing.T) {
	rectangleRepoMock := repomock.
		NewRectangleRepoMockBuilder().
		WithCreateRectangleMock(func(t *entity.Rectangle) (*entity.Rectangle, error) {
			return &entity.Rectangle{ID: 1, L: 3, W: 4}, nil
		}).
		Build()
	rectangleUC := usecase.NewRectangleUsecase(rectangleRepoMock)
	input := &usecase.CreateRectangleInput{L: 3, W: 4}

	rectangle, err := rectangleUC.CreateRectangle(input)

	assert.Nil(t, err)
	assert.NotNil(t, rectangle)
}

func TestGetNonExistRectangle(t *testing.T) {
	rectangleRepoMock := repomock.
		NewRectangleRepoMockBuilder().
		WithGetRectangleByIDMock(func(_ int64) *entity.Rectangle { return nil }).
		Build()
	rectangleUC := usecase.NewRectangleUsecase(rectangleRepoMock)

	rectangle := rectangleUC.GetRectangleByID(100)

	assert.Nil(t, rectangle)
}

func TestGetRectangle(t *testing.T) {
	rectangleRepoMock := repomock.
		NewRectangleRepoMockBuilder().
		WithGetRectangleByIDMock(func(id int64) *entity.Rectangle {
			if id == 10 {
				return &entity.Rectangle{ID: 10, L: 3, W: 4}
			}
			return nil
		}).
		Build()
	rectangleUC := usecase.NewRectangleUsecase(rectangleRepoMock)

	rectangle := rectangleUC.GetRectangleByID(10)

	if assert.NotNil(t, rectangle) {
		assert.Equal(t, rectangle.ID, int64(10))
		assert.Equal(t, rectangle.L, 3.0)
		assert.Equal(t, rectangle.W, 4.0)
	}
}

func TestUpdateNonExistRectangle(t *testing.T) {
	rectangleRepoMock := repomock.
		NewRectangleRepoMockBuilder().
		WithGetRectangleByIDMock(func(_ int64) *entity.Rectangle { return nil }).
		Build()
	rectangleUC := usecase.NewRectangleUsecase(rectangleRepoMock)
	input := genValidUpdateRectangleInput()
	input.ID = 1000

	_, err := rectangleUC.UpdateRectangle(input)

	if !errors.Is(err, repo.ErrRectangleNotFound) {
		t.Errorf("expected error '%v', got '%v'", repo.ErrRectangleNotFound, err)
	}
}

func TestUpdateWithInvalidRectangle(t *testing.T) {
	rectangleRepoMock := repomock.
		NewRectangleRepoMockBuilder().
		WithGetRectangleByIDMock(func(_ int64) *entity.Rectangle {
			return &entity.Rectangle{ID: 1, L: 3, W: 4}
		}).
		Build()
	rectangleUC := usecase.NewRectangleUsecase(rectangleRepoMock)
	input := genValidUpdateRectangleInput()
	input.W = -1000

	_, err := rectangleUC.UpdateRectangle(input)

	if !errors.Is(err, entity.ErrInvalidRectangle) {
		t.Errorf("expected error '%v', got '%v'", entity.ErrInvalidRectangle, err)
	}
}

func TestUpdateRectangle(t *testing.T) {
	rectangleRepoMock := repomock.
		NewRectangleRepoMockBuilder().
		WithGetRectangleByIDMock(func(id int64) *entity.Rectangle {
			if id == 1 {
				return &entity.Rectangle{ID: 1, L: 4, W: 5}
			}
			return nil
		}).
		WithUpdateRectangleMock(func(t *entity.Rectangle) (*entity.Rectangle, error) {
			if t.ID == 1 {
				return t, nil
			}
			return nil, nil
		}).
		Build()
	rectangleUC := usecase.NewRectangleUsecase(rectangleRepoMock)
	input := genValidUpdateRectangleInput()

	rectangle, err := rectangleUC.UpdateRectangle(input)

	if assert.Nil(t, err) {
		assert.Equal(t, rectangle.ID, int64(1))
		assert.Equal(t, rectangle.L, 3.0)
		assert.Equal(t, rectangle.W, 4.0)
	}
}

func genValidUpdateRectangleInput() *usecase.UpdateRectangleInput {
	return &usecase.UpdateRectangleInput{
		ID:                   1,
		CreateRectangleInput: usecase.CreateRectangleInput{L: 3, W: 4},
	}
}

func TestDeleteNonExistRectangle(t *testing.T) {
	rectangleRepoMock := repomock.
		NewRectangleRepoMockBuilder().
		WithGetRectangleByIDMock(func(_ int64) *entity.Rectangle {
			return nil
		}).
		Build()
	rectangleUC := usecase.NewRectangleUsecase(rectangleRepoMock)

	err := rectangleUC.DeleteRectangleByID(1)
	assert.Nil(t, err)
}

func TestDeleteRectangleWithErrorRepo(t *testing.T) {
	expectedErr := errors.New("error")
	rectangleRepoMock := repomock.
		NewRectangleRepoMockBuilder().
		WithGetRectangleByIDMock(func(id int64) *entity.Rectangle {
			if id == 1 {
				return &entity.Rectangle{ID: 1, L: 3, W: 4}
			}
			return nil
		}).
		WithDeleteRectangleByIDMock(func(i int64) error {
			return expectedErr
		}).
		Build()
	rectangleUC := usecase.NewRectangleUsecase(rectangleRepoMock)

	err := rectangleUC.DeleteRectangleByID(1)
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error '%v', got '%v'", expectedErr, err)
	}
}

func TestDeleteRectangle(t *testing.T) {
	rectangleRepoMock := repomock.
		NewRectangleRepoMockBuilder().
		WithGetRectangleByIDMock(func(id int64) *entity.Rectangle {
			if id == 1 {
				return &entity.Rectangle{ID: 1, L: 3, W: 4}
			}
			return nil
		}).
		WithDeleteRectangleByIDMock(func(i int64) error { return nil }).
		Build()
	rectangleUC := usecase.NewRectangleUsecase(rectangleRepoMock)

	err := rectangleUC.DeleteRectangleByID(1)
	assert.Nil(t, err)
}
