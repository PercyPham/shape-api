package repomock

import (
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

var _ repo.User = new(userRepoMock)

type userRepoMock struct {
	getUserByUsernameMock func(string) *entity.User
	createUserMock        func(*entity.User) error
}

func (urm *userRepoMock) GetUserByUsername(username string) *entity.User {
	return urm.getUserByUsernameMock(username)
}

func (urm *userRepoMock) CreateUser(user *entity.User) error {
	return urm.createUserMock(user)
}
