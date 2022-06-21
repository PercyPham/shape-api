package repo

import (
	"shape-api/internal/common/apperr"
	"shape-api/internal/entity"
)

var (
	ErrRectangleNotFound = apperr.New("rectangle not found")
)

type Rectangle interface {
	CreateRectangle(*entity.Rectangle) (*entity.Rectangle, error)
	GetRectangleByID(id int64) *entity.Rectangle
	UpdateRectangle(*entity.Rectangle) (*entity.Rectangle, error)
	DeleteRectangleByID(id int64) error
}
