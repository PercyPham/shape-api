package rest

import v1 "shape-api/internal/adapter/http/rest/v1"

func (s *server) setupAPIs() {
	apiV1 := s.r.Group("/api/v1")

	apiV1.POST("/users/register", v1.Register(s.userRepo))
	apiV1.POST("/users/login", v1.Login(s.userRepo))

	apiV1.POST("/triangles", v1.CreateTriangle(s.triangleRepo))
	apiV1.GET("/triangles/:id", v1.GetTriangleByID(s.triangleRepo))
	apiV1.GET("/triangles/:id/area", v1.GetTriangleAreaByID(s.triangleRepo))
	apiV1.GET("/triangles/:id/perimeter", v1.GetTrianglePerimeterByID(s.triangleRepo))
	apiV1.PUT("/triangles/:id", v1.UpdateTriangle(s.triangleRepo))
	apiV1.DELETE("/triangles/:id", v1.DeleteTriangleByID(s.triangleRepo))

	apiV1.POST("/rectangles", v1.CreateRectangle(s.rectangleRepo))
	apiV1.GET("/rectangles/:id", v1.GetRectangleByID(s.rectangleRepo))
	apiV1.GET("/rectangles/:id/area", v1.GetRectangleAreaByID(s.rectangleRepo))
	apiV1.GET("/rectangles/:id/perimeter", v1.GetRectanglePerimeterByID(s.rectangleRepo))
	apiV1.PUT("/rectangles/:id", v1.UpdateRectangle(s.rectangleRepo))
	apiV1.DELETE("/rectangles/:id", v1.DeleteRectangleByID(s.rectangleRepo))

	apiV1.POST("/squares", v1.CreateSquare(s.squareRepo))
	apiV1.GET("/squares/:id", v1.GetSquareByID(s.squareRepo))
	apiV1.GET("/squares/:id/area", v1.GetSquareAreaByID(s.squareRepo))
	apiV1.GET("/squares/:id/perimeter", v1.GetSquarePerimeterByID(s.squareRepo))
	apiV1.PUT("/squares/:id", v1.UpdateSquare(s.squareRepo))
	apiV1.DELETE("/squares/:id", v1.DeleteSquareByID(s.squareRepo))

	apiV1.POST("/diamonds", v1.CreateDiamond(s.diamondRepo))
	apiV1.GET("/diamonds/:id", v1.GetDiamondByID(s.diamondRepo))
	apiV1.GET("/diamonds/:id/area", v1.GetDiamondAreaByID(s.diamondRepo))
	apiV1.GET("/diamonds/:id/perimeter", v1.GetDiamondPerimeterByID(s.diamondRepo))
	apiV1.PUT("/diamonds/:id", v1.UpdateDiamond(s.diamondRepo))
	apiV1.DELETE("/diamonds/:id", v1.DeleteDiamondByID(s.diamondRepo))
}
