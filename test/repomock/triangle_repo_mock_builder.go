package repomock

import (
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

func NewTriangleRepoMockBuilder() *triangleRepoMockBuilder {
	return &triangleRepoMockBuilder{
		&triangleRepoMock{},
	}
}

type triangleRepoMockBuilder struct {
	triangleRepoMock *triangleRepoMock
}

func (b *triangleRepoMockBuilder) WithCreateTriangleMock(mockFunc func(*entity.Triangle) (*entity.Triangle, error)) *triangleRepoMockBuilder {
	b.triangleRepoMock.createTriangleMock = mockFunc
	return b
}

func (b *triangleRepoMockBuilder) WithGetTriangleByIDMock(mockFunc func(int64) *entity.Triangle) *triangleRepoMockBuilder {
	b.triangleRepoMock.getTriangleByIDMock = mockFunc
	return b
}

func (b *triangleRepoMockBuilder) WithUpdateTriangleMock(mockFunc func(*entity.Triangle) (*entity.Triangle, error)) *triangleRepoMockBuilder {
	b.triangleRepoMock.updateTriangleMock = mockFunc
	return b
}

func (b *triangleRepoMockBuilder) WithDeleteTriangleByIDMock(mockFunc func(int64) error) *triangleRepoMockBuilder {
	b.triangleRepoMock.deleteTriangleByIDMock = mockFunc
	return b
}

func (urmb *triangleRepoMockBuilder) Build() repo.Triangle {
	return urmb.triangleRepoMock
}
