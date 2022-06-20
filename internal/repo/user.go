package repo

import (
	"shape-api/internal/common/apperr"
	"shape-api/internal/entity"
)

var (
	ErrDuplicateUsername = apperr.New("duplicate username")
)

type User interface {
	CreateUser(*entity.User) error
}
