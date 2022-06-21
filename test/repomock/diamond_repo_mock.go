package repomock

import (
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

var _ repo.Diamond = new(diamondRepoMock)

type diamondRepoMock struct {
	getDiamondByIDMock    func(int64) *entity.Diamond
	createDiamondMock     func(*entity.Diamond) (*entity.Diamond, error)
	updateDiamondMock     func(*entity.Diamond) (*entity.Diamond, error)
	deleteDiamondByIDMock func(int64) error
}

func (m *diamondRepoMock) GetDiamondByID(id int64) *entity.Diamond {
	return m.getDiamondByIDMock(id)
}

func (m *diamondRepoMock) CreateDiamond(diamond *entity.Diamond) (*entity.Diamond, error) {
	return m.createDiamondMock(diamond)
}

func (m *diamondRepoMock) UpdateDiamond(diamond *entity.Diamond) (*entity.Diamond, error) {
	return m.updateDiamondMock(diamond)
}

func (m *diamondRepoMock) DeleteDiamondByID(id int64) error {
	return m.deleteDiamondByIDMock(id)
}
