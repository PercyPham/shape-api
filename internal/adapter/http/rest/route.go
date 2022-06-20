package rest

import v1 "shape-api/internal/adapter/http/rest/v1"

func (s *server) setupAPIs() {
	apiV1 := s.r.Group("/api/v1")

	apiV1.POST("/users/register", v1.Register(s.userRepo))
}
