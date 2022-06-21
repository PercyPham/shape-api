package repo

import (
	"shape-api/internal/common/apperr"
	"shape-api/internal/entity"
)

var (
	ErrSquareNotFound = apperr.New("square not found")
)

type Square interface {
	CreateSquare(*entity.Square) (*entity.Square, error)
	GetSquareByID(id int64) *entity.Square
	UpdateSquare(*entity.Square) (*entity.Square, error)
	DeleteSquareByID(id int64) error
}
