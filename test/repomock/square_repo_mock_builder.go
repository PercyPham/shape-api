package repomock

import (
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

func NewSquareRepoMockBuilder() *squareRepoMockBuilder {
	return &squareRepoMockBuilder{
		&squareRepoMock{},
	}
}

type squareRepoMockBuilder struct {
	squareRepoMock *squareRepoMock
}

func (b *squareRepoMockBuilder) WithCreateSquareMock(mockFunc func(*entity.Square) (*entity.Square, error)) *squareRepoMockBuilder {
	b.squareRepoMock.createSquareMock = mockFunc
	return b
}

func (b *squareRepoMockBuilder) WithGetSquareByIDMock(mockFunc func(int64) *entity.Square) *squareRepoMockBuilder {
	b.squareRepoMock.getSquareByIDMock = mockFunc
	return b
}

func (b *squareRepoMockBuilder) WithUpdateSquareMock(mockFunc func(*entity.Square) (*entity.Square, error)) *squareRepoMockBuilder {
	b.squareRepoMock.updateSquareMock = mockFunc
	return b
}

func (b *squareRepoMockBuilder) WithDeleteSquareByIDMock(mockFunc func(int64) error) *squareRepoMockBuilder {
	b.squareRepoMock.deleteSquareByIDMock = mockFunc
	return b
}

func (urmb *squareRepoMockBuilder) Build() repo.Square {
	return urmb.squareRepoMock
}
