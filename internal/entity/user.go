package entity

import (
	"shape-api/internal/common/apperr"
	"shape-api/internal/common/util"
	"time"
)

var (
	ErrEmptyUsername = apperr.New("empty username")
	ErrEmptyPassword = apperr.New("empty password")
)

type User struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordSalt string `json:"password_salt"`
}

func NewUser(t time.Time, username, password string) (*User, error) {
	if username == "" {
		return nil, ErrEmptyUsername
	}
	if password == "" {
		return nil, ErrEmptyPassword
	}
	hashed, salt := hashPassword(t, password)
	return &User{
		Username:     username,
		Password:     hashed,
		PasswordSalt: salt,
	}, nil
}

func hashPassword(t time.Time, password string) (hashed, salt string) {
	salt = util.GenRanStr(t, 50)
	hashed = util.HashHmacSHA256(password, salt)
	return hashed, salt
}

func (u *User) IsCorrectPassword(password string) bool {
	hashed := util.HashHmacSHA256(password, u.PasswordSalt)
	return hashed == u.Password
}
