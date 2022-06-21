package mysqlrepo

import (
	"fmt"
	"shape-api/internal/entity"
	"shape-api/internal/repo"

	"gorm.io/gorm"
)

func NewRectangleRepo(db *gorm.DB) repo.Rectangle {
	return &rectangleRepo{db}
}

type rectangleRepo struct {
	db *gorm.DB
}

func (ur *rectangleRepo) GetRectangleByID(id int64) *entity.Rectangle {
	rectangle := new(entity.Rectangle)
	result := ur.db.Where("id = ?", id).First(rectangle)
	if result.Error != nil {
		return nil
	}
	return rectangle
}

func (r *rectangleRepo) CreateRectangle(rectangle *entity.Rectangle) (*entity.Rectangle, error) {
	result := r.db.Create(rectangle)
	if result.Error != nil {
		return nil, fmt.Errorf("gorm create rectangle: %w", result.Error)
	}
	return rectangle, nil
}

func (r *rectangleRepo) UpdateRectangle(rectangle *entity.Rectangle) (*entity.Rectangle, error) {
	result := r.db.Save(rectangle)
	if result.Error != nil {
		return nil, fmt.Errorf("gorm create rectangle: %w", result.Error)
	}
	return rectangle, nil
}

func (ur *rectangleRepo) DeleteRectangleByID(id int64) error {
	rectangle := new(entity.Rectangle)
	result := ur.db.Delete(rectangle, id)
	if result.Error != nil {
		return fmt.Errorf("gorm delete rectangle: %w", result.Error)
	}
	return nil
}
