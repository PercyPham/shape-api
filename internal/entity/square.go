package entity

import "shape-api/internal/common/apperr"

var (
	ErrInvalidSquare = apperr.New("invalid square")
)

func NewSquare(a float64) (*Square, error) {
	if a <= 0 {
		return nil, ErrInvalidSquare
	}
	return &Square{0, a}, nil
}

type Square struct {
	ID int64 `json:"id"`

	A float64 `json:"a"`
}

func (s *Square) Perimeter() float64 {
	return 4 * s.A
}

func (s *Square) Area() float64 {
	return s.A * s.A
}
