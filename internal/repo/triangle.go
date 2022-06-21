package repo

import (
	"shape-api/internal/common/apperr"
	"shape-api/internal/entity"
)

var (
	ErrTriangleNotFound = apperr.New("triangle not found")
)

type Triangle interface {
	GetTriangleByID(id int64) *entity.Triangle
	CreateTriangle(*entity.Triangle) (*entity.Triangle, error)
	UpdateTriangle(*entity.Triangle) (*entity.Triangle, error)
	DeleteTriangleByID(id int64) error
}
