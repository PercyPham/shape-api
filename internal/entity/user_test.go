package entity_test

import (
	"errors"
	"shape-api/internal/entity"
	"testing"
	"time"
)

func TestCreateNewUserWithEmptyUsername(t *testing.T) {
	_, err := entity.NewUser(time.Now(), "", "password")

	if !errors.Is(err, entity.ErrEmptyUsername) {
		t.Errorf("expected error '%v', got '%v'", entity.ErrEmptyUsername, err)
	}
}

func TestCreateNewUserWithEmptyPassword(t *testing.T) {
	_, err := entity.NewUser(time.Now(), "username", "")

	if !errors.Is(err, entity.ErrEmptyPassword) {
		t.Errorf("expected error '%v', got '%v'", entity.ErrEmptyPassword, err)
	}
}

func TestCreateNewUser(t *testing.T) {
	testedTime := time.Date(2022, time.June, 20, 0, 0, 0, 0, time.UTC)

	user, err := entity.NewUser(testedTime, "username", "password")

	expectedSalt := "ikv0PNxOS8VazgZ3fH3c1M7nb3rbpsO2j9gVJBDa40r84DOnuI"
	expectedHash := "196ecaebcca78086904c38edf18cc8d375bf5a57eda19774273bda86d144a60e"

	if err != nil {
		t.Errorf("expected nil error, got '%v'", err)
	}
	if user.PasswordSalt != expectedSalt {
		t.Errorf("expected salt to be '%v', got '%v'", expectedSalt, user.PasswordSalt)
	}
	if user.Password != expectedHash {
		t.Errorf("expected hash to be '%v', got '%v'", expectedHash, user.Password)
	}
}

func TestCheckUserPassword(t *testing.T) {
	user, err := entity.NewUser(time.Now(), "username", "password")

	if err != nil {
		t.Errorf("expected nil error, got '%v'", err)
	}

	var tests = []struct {
		password string
		expected bool
	}{
		{"", false},
		{"wrong", false},
		{"password", true},
	}

	for _, test := range tests {
		if result := user.IsCorrectPassword(test.password); result != test.expected {
			t.Errorf("expected %v, got %v when checking password '%v'", test.expected, result, test.password)
		}
	}
}
