package entity_test

import (
	"errors"
	"shape-api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInvalidDiamonds(t *testing.T) {
	invalidDiamondInputs := [][]float64{
		{-1, 1},
		{0, 1},
	}
	for _, input := range invalidDiamondInputs {
		_, err := entity.NewDiamond(input[0], input[1])
		if !errors.Is(err, entity.ErrInvalidDiamond) {
			t.Errorf("expected error '%v', got '%v'", entity.ErrInvalidDiamond, err)
		}
	}
}

func TestNewValidDiamond(t *testing.T) {
	diamond, err := entity.NewDiamond(2, 3)
	if assert.Nil(t, err) {
		assert.Equal(t, diamond.ID, int64(0))
		assert.Equal(t, diamond.DiagonalA, 2.0)
		assert.Equal(t, diamond.DiagonalB, 3.0)
	}
}

func TestGetDiamondPerimeter(t *testing.T) {
	diamond, err := entity.NewDiamond(3, 4)
	if assert.Nil(t, err) {
		assert.Equal(t, diamond.Perimeter(), 10.0)
	}
}

func TestGetDiamondArea(t *testing.T) {
	diamond, err := entity.NewDiamond(3, 4)
	if assert.Nil(t, err) {
		assert.Equal(t, diamond.Area(), 6.0)
	}
}
