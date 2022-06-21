package repomock

import (
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

func NewDiamondRepoMockBuilder() *diamondRepoMockBuilder {
	return &diamondRepoMockBuilder{
		&diamondRepoMock{},
	}
}

type diamondRepoMockBuilder struct {
	diamondRepoMock *diamondRepoMock
}

func (b *diamondRepoMockBuilder) WithCreateDiamondMock(mockFunc func(*entity.Diamond) (*entity.Diamond, error)) *diamondRepoMockBuilder {
	b.diamondRepoMock.createDiamondMock = mockFunc
	return b
}

func (b *diamondRepoMockBuilder) WithGetDiamondByIDMock(mockFunc func(int64) *entity.Diamond) *diamondRepoMockBuilder {
	b.diamondRepoMock.getDiamondByIDMock = mockFunc
	return b
}

func (b *diamondRepoMockBuilder) WithUpdateDiamondMock(mockFunc func(*entity.Diamond) (*entity.Diamond, error)) *diamondRepoMockBuilder {
	b.diamondRepoMock.updateDiamondMock = mockFunc
	return b
}

func (b *diamondRepoMockBuilder) WithDeleteDiamondByIDMock(mockFunc func(int64) error) *diamondRepoMockBuilder {
	b.diamondRepoMock.deleteDiamondByIDMock = mockFunc
	return b
}

func (urmb *diamondRepoMockBuilder) Build() repo.Diamond {
	return urmb.diamondRepoMock
}
