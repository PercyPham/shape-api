package entity

import (
	"math"
	"shape-api/internal/common/apperr"
)

var (
	ErrInvalidDiamond = apperr.New("invalid diamond")
)

func NewDiamond(diagonalA, diagonalB float64) (*Diamond, error) {
	if diagonalA <= 0 || diagonalB <= 0 {
		return nil, ErrInvalidDiamond
	}
	return &Diamond{0, diagonalA, diagonalB}, nil
}

type Diamond struct {
	ID int64 `json:"id"`

	DiagonalA float64 `json:"diagonalA"`
	DiagonalB float64 `json:"diagonalB"`
}

func (d *Diamond) Perimeter() float64 {
	return 2 * math.Sqrt(d.DiagonalA*d.DiagonalA+d.DiagonalB*d.DiagonalB)
}

func (d *Diamond) Area() float64 {
	return (d.DiagonalA * d.DiagonalB) / 2
}
