package repomock

import (
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

func NewRectangleRepoMockBuilder() *rectangleRepoMockBuilder {
	return &rectangleRepoMockBuilder{
		&rectangleRepoMock{},
	}
}

type rectangleRepoMockBuilder struct {
	rectangleRepoMock *rectangleRepoMock
}

func (b *rectangleRepoMockBuilder) WithCreateRectangleMock(mockFunc func(*entity.Rectangle) (*entity.Rectangle, error)) *rectangleRepoMockBuilder {
	b.rectangleRepoMock.createRectangleMock = mockFunc
	return b
}

func (b *rectangleRepoMockBuilder) WithGetRectangleByIDMock(mockFunc func(int64) *entity.Rectangle) *rectangleRepoMockBuilder {
	b.rectangleRepoMock.getRectangleByIDMock = mockFunc
	return b
}

func (b *rectangleRepoMockBuilder) WithUpdateRectangleMock(mockFunc func(*entity.Rectangle) (*entity.Rectangle, error)) *rectangleRepoMockBuilder {
	b.rectangleRepoMock.updateRectangleMock = mockFunc
	return b
}

func (b *rectangleRepoMockBuilder) WithDeleteRectangleByIDMock(mockFunc func(int64) error) *rectangleRepoMockBuilder {
	b.rectangleRepoMock.deleteRectangleByIDMock = mockFunc
	return b
}

func (urmb *rectangleRepoMockBuilder) Build() repo.Rectangle {
	return urmb.rectangleRepoMock
}
