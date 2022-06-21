package v1

import (
	"errors"
	"fmt"
	"net/http"
	"shape-api/internal/adapter/http/rest/internal/authcheck"
	"shape-api/internal/adapter/http/rest/internal/response"
	"shape-api/internal/adapter/http/rest/v1/internal"
	"shape-api/internal/entity"
	"shape-api/internal/repo"
	"shape-api/internal/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTriangle(triangleRepo repo.Triangle) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := authcheck.IsAuthenticated(time.Now(), c); err != nil {
			response.Error(c, http.StatusUnauthorized, fmt.Errorf("authenticate: %w", err))
			return
		}

		input := new(usecase.CreateTriangleInput)
		if err := c.BindJSON(input); err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Errorf("bind json: %v", err))
			return
		}

		triangleUC := usecase.NewTriangleUsecase(triangleRepo)
		triangle, err := triangleUC.CreateTriangle(input)
		if err != nil {
			if errors.Is(err, usecase.ErrInvalidTriangleInput) || errors.Is(err, entity.ErrInvalidTriangle) {
				response.Error(c, http.StatusBadRequest, err)
				return
			}
			response.Error(c, http.StatusInternalServerError, err)
			return
		}

		response.Success(c, http.StatusCreated, triangle)
	}
}

func GetTriangleByID(triangleRepo repo.Triangle) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := authcheck.IsAuthenticated(time.Now(), c); err != nil {
			response.Error(c, http.StatusUnauthorized, fmt.Errorf("authenticate: %w", err))
			return
		}

		id, err := internal.GetIDFromParam(c, "id")
		if err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Errorf("get id from param: %w", err))
			return
		}

		triangleUC := usecase.NewTriangleUsecase(triangleRepo)
		triangle := triangleUC.GetTriangleByID(id)
		if triangle == nil {
			response.Error(c, http.StatusNotFound, fmt.Errorf("triangle with id '%d' not found", id))
			return
		}

		response.Success(c, http.StatusOK, triangle)
	}
}

func UpdateTriangle(triangleRepo repo.Triangle) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := authcheck.IsAuthenticated(time.Now(), c); err != nil {
			response.Error(c, http.StatusUnauthorized, fmt.Errorf("authenticate: %w", err))
			return
		}

		id, err := internal.GetIDFromParam(c, "id")
		if err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Errorf("get id from param: %w", err))
			return
		}

		input := new(usecase.UpdateTriangleInput)
		if err := c.BindJSON(input); err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Errorf("bind json: %v", err))
			return
		}
		input.ID = id

		triangleUC := usecase.NewTriangleUsecase(triangleRepo)
		triangle, err := triangleUC.UpdateTriangle(input)
		if err != nil {
			if errors.Is(err, usecase.ErrInvalidTriangleInput) ||
				errors.Is(err, entity.ErrInvalidTriangle) {
				response.Error(c, http.StatusBadRequest, err)
				return
			}
			if errors.Is(err, repo.ErrTriangleNotFound) {
				response.Error(c, http.StatusNotFound, err)
				return
			}
			response.Error(c, http.StatusInternalServerError, err)
			return
		}

		response.Success(c, http.StatusCreated, triangle)
	}
}

func DeleteTriangleByID(triangleRepo repo.Triangle) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := authcheck.IsAuthenticated(time.Now(), c); err != nil {
			response.Error(c, http.StatusUnauthorized, fmt.Errorf("authenticate: %w", err))
			return
		}

		id, err := internal.GetIDFromParam(c, "id")
		if err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Errorf("get id from param: %w", err))
			return
		}

		triangleUC := usecase.NewTriangleUsecase(triangleRepo)
		triangleUC.DeleteTriangleByID(id)

		response.Success(c, http.StatusOK, true)
	}
}
