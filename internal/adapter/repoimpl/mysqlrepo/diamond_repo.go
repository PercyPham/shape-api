package mysqlrepo

import (
	"fmt"
	"shape-api/internal/entity"
	"shape-api/internal/repo"

	"gorm.io/gorm"
)

func NewDiamondRepo(db *gorm.DB) repo.Diamond {
	return &diamondRepo{db}
}

type diamondRepo struct {
	db *gorm.DB
}

func (ur *diamondRepo) GetDiamondByID(id int64) *entity.Diamond {
	diamond := new(entity.Diamond)
	result := ur.db.Where("id = ?", id).First(diamond)
	if result.Error != nil {
		return nil
	}
	return diamond
}

func (r *diamondRepo) CreateDiamond(diamond *entity.Diamond) (*entity.Diamond, error) {
	result := r.db.Create(diamond)
	if result.Error != nil {
		return nil, fmt.Errorf("gorm create diamond: %w", result.Error)
	}
	return diamond, nil
}

func (r *diamondRepo) UpdateDiamond(diamond *entity.Diamond) (*entity.Diamond, error) {
	result := r.db.Save(diamond)
	if result.Error != nil {
		return nil, fmt.Errorf("gorm create diamond: %w", result.Error)
	}
	return diamond, nil
}

func (ur *diamondRepo) DeleteDiamondByID(id int64) error {
	diamond := new(entity.Diamond)
	result := ur.db.Delete(diamond, id)
	if result.Error != nil {
		return fmt.Errorf("gorm delete diamond: %w", result.Error)
	}
	return nil
}
