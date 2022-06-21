package entity_test

import (
	"errors"
	"shape-api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInvalidTriangles(t *testing.T) {
	invalidTriangleInputs := [][]float64{
		{2, 2, -3},
		{2, -2, 3},
		{-2, 2, 3},
		{0, 1, 2},
		{1, 0, 2},
		{1, 2, 0},
		{1, 2, 3},
		{1, 3, 2},
		{3, 1, 2},
		{1, 2, 5},
		{5, 1, 2},
		{1, 5, 2},
	}
	for _, input := range invalidTriangleInputs {
		_, err := entity.NewTriangle(input[0], input[1], input[2])
		if !errors.Is(err, entity.ErrInvalidTriangle) {
			t.Errorf("expected error '%v', got '%v'", entity.ErrInvalidTriangle, err)
		}
	}
}

func TestNewValidTriangle(t *testing.T) {
	triangle, err := entity.NewTriangle(2, 3, 4)
	if assert.Nil(t, err) {
		assert.Equal(t, triangle.ID, int64(0))
		assert.Equal(t, triangle.A, 2.0)
		assert.Equal(t, triangle.B, 3.0)
		assert.Equal(t, triangle.C, 4.0)
	}
}

func TestGetTrianglePerimeter(t *testing.T) {
	triangle, err := entity.NewTriangle(2, 3, 4)
	if assert.Nil(t, err) {
		assert.Equal(t, triangle.Perimeter(), 9.0)
	}
}

func TestGetTriangleArea(t *testing.T) {
	triangle, err := entity.NewTriangle(3, 4, 5)
	if assert.Nil(t, err) {
		assert.Equal(t, triangle.Area(), 6.0)
	}
}
