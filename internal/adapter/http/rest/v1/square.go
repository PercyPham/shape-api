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

func CreateSquare(squareRepo repo.Square) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := authcheck.IsAuthenticated(time.Now(), c); err != nil {
			response.Error(c, http.StatusUnauthorized, fmt.Errorf("authenticate: %w", err))
			return
		}

		input := new(usecase.CreateSquareInput)
		if err := c.BindJSON(input); err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Errorf("bind json: %v", err))
			return
		}

		squareUC := usecase.NewSquareUsecase(squareRepo)
		square, err := squareUC.CreateSquare(input)
		if err != nil {
			if errors.Is(err, usecase.ErrInvalidSquareInput) || errors.Is(err, entity.ErrInvalidSquare) {
				response.Error(c, http.StatusBadRequest, err)
				return
			}
			response.Error(c, http.StatusInternalServerError, err)
			return
		}

		response.Success(c, http.StatusCreated, square)
	}
}

func GetSquareByID(squareRepo repo.Square) func(*gin.Context) {
	return func(c *gin.Context) {
		square, errHttpStatusCode, err := getSquareByID(c, squareRepo)
		if err != nil {
			response.Error(c, errHttpStatusCode, err)
			return
		}

		response.Success(c, http.StatusOK, square)
	}
}

func GetSquareAreaByID(squareRepo repo.Square) func(*gin.Context) {
	return func(c *gin.Context) {
		square, errHttpStatusCode, err := getSquareByID(c, squareRepo)
		if err != nil {
			response.Error(c, errHttpStatusCode, err)
			return
		}

		response.Success(c, http.StatusOK, square.Area())
	}
}

func GetSquarePerimeterByID(squareRepo repo.Square) func(*gin.Context) {
	return func(c *gin.Context) {
		square, errHttpStatusCode, err := getSquareByID(c, squareRepo)
		if err != nil {
			response.Error(c, errHttpStatusCode, err)
			return
		}

		response.Success(c, http.StatusOK, square.Perimeter())
	}
}

func getSquareByID(c *gin.Context, squareRepo repo.Square) (square *entity.Square, errHttpStatusCode int, err error) {
	if err := authcheck.IsAuthenticated(time.Now(), c); err != nil {
		return nil, http.StatusUnauthorized, fmt.Errorf("authenticate: %w", err)
	}

	id, err := internal.GetIDFromParam(c, "id")
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("get id from param: %w", err)
	}

	squareUC := usecase.NewSquareUsecase(squareRepo)
	square = squareUC.GetSquareByID(id)
	if square == nil {
		return nil, http.StatusNotFound, fmt.Errorf("square with id '%d' not found", id)
	}

	return square, 0, nil
}

func UpdateSquare(squareRepo repo.Square) func(*gin.Context) {
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

		input := new(usecase.UpdateSquareInput)
		if err := c.BindJSON(input); err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Errorf("bind json: %v", err))
			return
		}
		input.ID = id

		squareUC := usecase.NewSquareUsecase(squareRepo)
		square, err := squareUC.UpdateSquare(input)
		if err != nil {
			if errors.Is(err, usecase.ErrInvalidSquareInput) ||
				errors.Is(err, entity.ErrInvalidSquare) {
				response.Error(c, http.StatusBadRequest, err)
				return
			}
			if errors.Is(err, repo.ErrSquareNotFound) {
				response.Error(c, http.StatusNotFound, err)
				return
			}
			response.Error(c, http.StatusInternalServerError, err)
			return
		}

		response.Success(c, http.StatusCreated, square)
	}
}

func DeleteSquareByID(squareRepo repo.Square) func(*gin.Context) {
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

		squareUC := usecase.NewSquareUsecase(squareRepo)
		squareUC.DeleteSquareByID(id)

		response.Success(c, http.StatusOK, true)
	}
}
