package entity

import (
	"math"
	"shape-api/internal/common/apperr"
)

var (
	ErrInvalidTriangle = apperr.New("invalid triangle")
)

func NewTriangle(a, b, c float64) (*Triangle, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return nil, ErrInvalidTriangle
	}
	if a+b <= c || b+c <= a || a+c <= b {
		return nil, ErrInvalidTriangle
	}
	return &Triangle{0, a, b, c}, nil
}

type Triangle struct {
	ID int64 `json:"id"`

	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
}

func (t *Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func (t *Triangle) Area() float64 {
	// Heron's formula
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}
