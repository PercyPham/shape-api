package repomock

import (
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

var _ repo.Square = new(squareRepoMock)

type squareRepoMock struct {
	getSquareByIDMock    func(int64) *entity.Square
	createSquareMock     func(*entity.Square) (*entity.Square, error)
	updateSquareMock     func(*entity.Square) (*entity.Square, error)
	deleteSquareByIDMock func(int64) error
}

func (m *squareRepoMock) GetSquareByID(id int64) *entity.Square {
	return m.getSquareByIDMock(id)
}

func (m *squareRepoMock) CreateSquare(square *entity.Square) (*entity.Square, error) {
	return m.createSquareMock(square)
}

func (m *squareRepoMock) UpdateSquare(square *entity.Square) (*entity.Square, error) {
	return m.updateSquareMock(square)
}

func (m *squareRepoMock) DeleteSquareByID(id int64) error {
	return m.deleteSquareByIDMock(id)
}
