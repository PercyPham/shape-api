package entity

import "shape-api/internal/common/apperr"

var (
	ErrInvalidRectangle = apperr.New("invalid rectangle")
)

func NewRectangle(l, w float64) (*Rectangle, error) {
	if l <= 0 || w <= 0 {
		return nil, ErrInvalidRectangle
	}
	return &Rectangle{0, l, w}, nil
}

type Rectangle struct {
	ID int64 `json:"id"`

	L float64 `json:"l"`
	W float64 `json:"w"`
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.L + 2*r.W
}

func (r *Rectangle) Area() float64 {
	return r.L * r.W
}
