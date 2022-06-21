package repo

import (
	"shape-api/internal/common/apperr"
	"shape-api/internal/entity"
)

var (
	ErrDiamondNotFound = apperr.New("diamond not found")
)

type Diamond interface {
	CreateDiamond(*entity.Diamond) (*entity.Diamond, error)
	GetDiamondByID(id int64) *entity.Diamond
	UpdateDiamond(*entity.Diamond) (*entity.Diamond, error)
	DeleteDiamondByID(id int64) error
}
