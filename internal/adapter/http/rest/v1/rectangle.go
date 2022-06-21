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

func CreateRectangle(rectangleRepo repo.Rectangle) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := authcheck.IsAuthenticated(time.Now(), c); err != nil {
			response.Error(c, http.StatusUnauthorized, fmt.Errorf("authenticate: %w", err))
			return
		}

		input := new(usecase.CreateRectangleInput)
		if err := c.BindJSON(input); err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Errorf("bind json: %v", err))
			return
		}

		rectangleUC := usecase.NewRectangleUsecase(rectangleRepo)
		rectangle, err := rectangleUC.CreateRectangle(input)
		if err != nil {
			if errors.Is(err, usecase.ErrInvalidRectangleInput) || errors.Is(err, entity.ErrInvalidRectangle) {
				response.Error(c, http.StatusBadRequest, err)
				return
			}
			response.Error(c, http.StatusInternalServerError, err)
			return
		}

		response.Success(c, http.StatusCreated, rectangle)
	}
}

func GetRectangleByID(rectangleRepo repo.Rectangle) func(*gin.Context) {
	return func(c *gin.Context) {
		rectangle, errHttpStatusCode, err := getRectangleByID(c, rectangleRepo)
		if err != nil {
			response.Error(c, errHttpStatusCode, err)
			return
		}

		response.Success(c, http.StatusOK, rectangle)
	}
}

func GetRectangleAreaByID(rectangleRepo repo.Rectangle) func(*gin.Context) {
	return func(c *gin.Context) {
		rectangle, errHttpStatusCode, err := getRectangleByID(c, rectangleRepo)
		if err != nil {
			response.Error(c, errHttpStatusCode, err)
			return
		}

		response.Success(c, http.StatusOK, rectangle.Area())
	}
}

func GetRectanglePerimeterByID(rectangleRepo repo.Rectangle) func(*gin.Context) {
	return func(c *gin.Context) {
		rectangle, errHttpStatusCode, err := getRectangleByID(c, rectangleRepo)
		if err != nil {
			response.Error(c, errHttpStatusCode, err)
			return
		}

		response.Success(c, http.StatusOK, rectangle.Perimeter())
	}
}

func getRectangleByID(c *gin.Context, rectangleRepo repo.Rectangle) (rectangle *entity.Rectangle, errHttpStatusCode int, err error) {
	if err := authcheck.IsAuthenticated(time.Now(), c); err != nil {
		return nil, http.StatusUnauthorized, fmt.Errorf("authenticate: %w", err)
	}

	id, err := internal.GetIDFromParam(c, "id")
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("get id from param: %w", err)
	}

	rectangleUC := usecase.NewRectangleUsecase(rectangleRepo)
	rectangle = rectangleUC.GetRectangleByID(id)
	if rectangle == nil {
		return nil, http.StatusNotFound, fmt.Errorf("rectangle with id '%d' not found", id)
	}

	return rectangle, 0, nil
}

func UpdateRectangle(rectangleRepo repo.Rectangle) func(*gin.Context) {
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

		input := new(usecase.UpdateRectangleInput)
		if err := c.BindJSON(input); err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Errorf("bind json: %v", err))
			return
		}
		input.ID = id

		rectangleUC := usecase.NewRectangleUsecase(rectangleRepo)
		rectangle, err := rectangleUC.UpdateRectangle(input)
		if err != nil {
			if errors.Is(err, usecase.ErrInvalidRectangleInput) ||
				errors.Is(err, entity.ErrInvalidRectangle) {
				response.Error(c, http.StatusBadRequest, err)
				return
			}
			if errors.Is(err, repo.ErrRectangleNotFound) {
				response.Error(c, http.StatusNotFound, err)
				return
			}
			response.Error(c, http.StatusInternalServerError, err)
			return
		}

		response.Success(c, http.StatusCreated, rectangle)
	}
}

func DeleteRectangleByID(rectangleRepo repo.Rectangle) func(*gin.Context) {
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

		rectangleUC := usecase.NewRectangleUsecase(rectangleRepo)
		rectangleUC.DeleteRectangleByID(id)

		response.Success(c, http.StatusOK, true)
	}
}
