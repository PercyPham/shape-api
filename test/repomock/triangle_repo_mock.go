package repomock

import (
	"shape-api/internal/entity"
	"shape-api/internal/repo"
)

var _ repo.Triangle = new(triangleRepoMock)

type triangleRepoMock struct {
	getTriangleByIDMock    func(int64) *entity.Triangle
	createTriangleMock     func(*entity.Triangle) (*entity.Triangle, error)
	updateTriangleMock     func(*entity.Triangle) (*entity.Triangle, error)
	deleteTriangleByIDMock func(int64) error
}

func (m *triangleRepoMock) GetTriangleByID(id int64) *entity.Triangle {
	return m.getTriangleByIDMock(id)
}

func (m *triangleRepoMock) CreateTriangle(triangle *entity.Triangle) (*entity.Triangle, error) {
	return m.createTriangleMock(triangle)
}

func (m *triangleRepoMock) UpdateTriangle(triangle *entity.Triangle) (*entity.Triangle, error) {
	return m.updateTriangleMock(triangle)
}

func (m *triangleRepoMock) DeleteTriangleByID(id int64) error {
	return m.deleteTriangleByIDMock(id)
}
