package entity_test

import (
	"errors"
	"shape-api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInvalidRectangles(t *testing.T) {
	invalidRectangleInputs := [][]float64{
		{-1, 1},
		{0, 1},
	}
	for _, input := range invalidRectangleInputs {
		_, err := entity.NewRectangle(input[0], input[1])
		if !errors.Is(err, entity.ErrInvalidRectangle) {
			t.Errorf("expected error '%v', got '%v'", entity.ErrInvalidRectangle, err)
		}
	}
}

func TestNewValidRectangle(t *testing.T) {
	rectangle, err := entity.NewRectangle(2, 3)
	if assert.Nil(t, err) {
		assert.Equal(t, rectangle.ID, int64(0))
		assert.Equal(t, rectangle.L, 2.0)
		assert.Equal(t, rectangle.W, 3.0)
	}
}

func TestGetRectanglePerimeter(t *testing.T) {
	rectangle, err := entity.NewRectangle(2, 3)
	if assert.Nil(t, err) {
		assert.Equal(t, rectangle.Perimeter(), 10.0)
	}
}

func TestGetRectangleArea(t *testing.T) {
	rectangle, err := entity.NewRectangle(3, 4)
	if assert.Nil(t, err) {
		assert.Equal(t, rectangle.Area(), 12.0)
	}
}
