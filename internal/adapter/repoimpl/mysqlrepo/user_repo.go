package mysqlrepo

import (
	"fmt"
	"shape-api/internal/entity"
	"shape-api/internal/repo"
	"strings"

	"gorm.io/gorm"
)

func NewUserRepo(db *gorm.DB) repo.User {
	return &userRepo{db}
}

type userRepo struct {
	db *gorm.DB
}

func (ur *userRepo) GetUserByUsername(username string) *entity.User {
	user := new(entity.User)
	result := ur.db.Where("username = ?", username).First(user)
	if result.Error != nil {
		return nil
	}
	return user
}

func (r *userRepo) CreateUser(user *entity.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			return fmt.Errorf("gorm create user: %w", repo.ErrDuplicateUsername)
		}
		return fmt.Errorf("gorm create user: %w", result.Error)
	}
	return nil
}
