package rest

import v1 "shape-api/internal/adapter/http/rest/v1"

func (s *server) setupAPIs() {
	apiV1 := s.r.Group("/api/v1")

	apiV1.POST("/users/register", v1.Register(s.userRepo))
	apiV1.POST("/users/login", v1.Login(s.userRepo))

	apiV1.POST("/triangles", v1.CreateTriangle(s.triangleRepo))
	apiV1.GET("/triangles/:id", v1.GetTriangleByID(s.triangleRepo))
	apiV1.PUT("/triangles/:id", v1.UpdateTriangle(s.triangleRepo))
	apiV1.DELETE("/triangles/:id", v1.DeleteTriangleByID(s.triangleRepo))
}
