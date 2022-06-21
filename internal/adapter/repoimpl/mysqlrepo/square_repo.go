package mysqlrepo

import (
	"fmt"
	"shape-api/internal/entity"
	"shape-api/internal/repo"

	"gorm.io/gorm"
)

func NewSquareRepo(db *gorm.DB) repo.Square {
	return &squareRepo{db}
}

type squareRepo struct {
	db *gorm.DB
}

func (ur *squareRepo) GetSquareByID(id int64) *entity.Square {
	square := new(entity.Square)
	result := ur.db.Where("id = ?", id).First(square)
	if result.Error != nil {
		return nil
	}
	return square
}

func (r *squareRepo) CreateSquare(square *entity.Square) (*entity.Square, error) {
	result := r.db.Create(square)
	if result.Error != nil {
		return nil, fmt.Errorf("gorm create square: %w", result.Error)
	}
	return square, nil
}

func (r *squareRepo) UpdateSquare(square *entity.Square) (*entity.Square, error) {
	result := r.db.Save(square)
	if result.Error != nil {
		return nil, fmt.Errorf("gorm create square: %w", result.Error)
	}
	return square, nil
}

func (ur *squareRepo) DeleteSquareByID(id int64) error {
	square := new(entity.Square)
	result := ur.db.Delete(square, id)
	if result.Error != nil {
		return fmt.Errorf("gorm delete square: %w", result.Error)
	}
	return nil
}
