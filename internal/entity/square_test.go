package entity_test

import (
	"errors"
	"shape-api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInvalidSquares(t *testing.T) {
	invalidSquareInputs := []float64{-1, 0}
	for _, input := range invalidSquareInputs {
		_, err := entity.NewSquare(input)
		if !errors.Is(err, entity.ErrInvalidSquare) {
			t.Errorf("expected error '%v', got '%v'", entity.ErrInvalidSquare, err)
		}
	}
}

func TestNewValidSquare(t *testing.T) {
	square, err := entity.NewSquare(2)
	if assert.Nil(t, err) {
		assert.Equal(t, square.ID, int64(0))
		assert.Equal(t, square.A, 2.0)
	}
}

func TestGetSquarePerimeter(t *testing.T) {
	square, err := entity.NewSquare(2)
	if assert.Nil(t, err) {
		assert.Equal(t, square.Perimeter(), 8.0)
	}
}

func TestGetSquareArea(t *testing.T) {
	rectangle, err := entity.NewSquare(2)
	if assert.Nil(t, err) {
		assert.Equal(t, rectangle.Area(), 4.0)
	}
}
