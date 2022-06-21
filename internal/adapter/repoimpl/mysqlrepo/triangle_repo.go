package mysqlrepo

import (
	"fmt"
	"shape-api/internal/entity"
	"shape-api/internal/repo"

	"gorm.io/gorm"
)

func NewTriangleRepo(db *gorm.DB) repo.Triangle {
	return &triangleRepo{db}
}

type triangleRepo struct {
	db *gorm.DB
}

func (ur *triangleRepo) GetTriangleByID(id int64) *entity.Triangle {
	triangle := new(entity.Triangle)
	result := ur.db.Where("id = ?", id).First(triangle)
	if result.Error != nil {
		return nil
	}
	return triangle
}

func (r *triangleRepo) CreateTriangle(triangle *entity.Triangle) (*entity.Triangle, error) {
	result := r.db.Create(triangle)
	if result.Error != nil {
		return nil, fmt.Errorf("gorm create triangle: %w", result.Error)
	}
	return triangle, nil
}

func (r *triangleRepo) UpdateTriangle(triangle *entity.Triangle) (*entity.Triangle, error) {
	result := r.db.Save(triangle)
	if result.Error != nil {
		return nil, fmt.Errorf("gorm create triangle: %w", result.Error)
	}
	return triangle, nil
}

func (ur *triangleRepo) DeleteTriangleByID(id int64) error {
	triangle := new(entity.Triangle)
	result := ur.db.Delete(triangle, id)
	if result.Error != nil {
		return fmt.Errorf("gorm delete triangle: %w", result.Error)
	}
	return nil
}
