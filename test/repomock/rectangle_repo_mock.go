package repomock

import (
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

var _ repo.Rectangle = new(rectangleRepoMock)

type rectangleRepoMock struct {
	getRectangleByIDMock    func(int64) *entity.Rectangle
	createRectangleMock     func(*entity.Rectangle) (*entity.Rectangle, error)
	updateRectangleMock     func(*entity.Rectangle) (*entity.Rectangle, error)
	deleteRectangleByIDMock func(int64) error
}

func (m *rectangleRepoMock) GetRectangleByID(id int64) *entity.Rectangle {
	return m.getRectangleByIDMock(id)
}

func (m *rectangleRepoMock) CreateRectangle(rectangle *entity.Rectangle) (*entity.Rectangle, error) {
	return m.createRectangleMock(rectangle)
}

func (m *rectangleRepoMock) UpdateRectangle(rectangle *entity.Rectangle) (*entity.Rectangle, error) {
	return m.updateRectangleMock(rectangle)
}

func (m *rectangleRepoMock) DeleteRectangleByID(id int64) error {
	return m.deleteRectangleByIDMock(id)
}
