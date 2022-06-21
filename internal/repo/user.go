package repo

import (
	"shape-api/internal/common/apperr"
	"shape-api/internal/entity"
)

var (
	ErrDuplicateUsername = apperr.New("duplicate username")
)

type User interface {
	GetUserByUsername(username string) *entity.User
	CreateUser(*entity.User) error
}
