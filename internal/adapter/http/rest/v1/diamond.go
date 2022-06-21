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

func CreateDiamond(diamondRepo repo.Diamond) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := authcheck.IsAuthenticated(time.Now(), c); err != nil {
			response.Error(c, http.StatusUnauthorized, fmt.Errorf("authenticate: %w", err))
			return
		}

		input := new(usecase.CreateDiamondInput)
		if err := c.BindJSON(input); err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Errorf("bind json: %v", err))
			return
		}

		diamondUC := usecase.NewDiamondUsecase(diamondRepo)
		diamond, err := diamondUC.CreateDiamond(input)
		if err != nil {
			if errors.Is(err, usecase.ErrInvalidDiamondInput) || errors.Is(err, entity.ErrInvalidDiamond) {
				response.Error(c, http.StatusBadRequest, err)
				return
			}
			response.Error(c, http.StatusInternalServerError, err)
			return
		}

		response.Success(c, http.StatusCreated, diamond)
	}
}

func GetDiamondByID(diamondRepo repo.Diamond) func(*gin.Context) {
	return func(c *gin.Context) {
		diamond, errHttpStatusCode, err := getDiamondByID(c, diamondRepo)
		if err != nil {
			response.Error(c, errHttpStatusCode, err)
			return
		}

		response.Success(c, http.StatusOK, diamond)
	}
}

func GetDiamondAreaByID(diamondRepo repo.Diamond) func(*gin.Context) {
	return func(c *gin.Context) {
		diamond, errHttpStatusCode, err := getDiamondByID(c, diamondRepo)
		if err != nil {
			response.Error(c, errHttpStatusCode, err)
			return
		}

		response.Success(c, http.StatusOK, diamond.Area())
	}
}

func GetDiamondPerimeterByID(diamondRepo repo.Diamond) func(*gin.Context) {
	return func(c *gin.Context) {
		diamond, errHttpStatusCode, err := getDiamondByID(c, diamondRepo)
		if err != nil {
			response.Error(c, errHttpStatusCode, err)
			return
		}

		response.Success(c, http.StatusOK, diamond.Perimeter())
	}
}

func getDiamondByID(c *gin.Context, diamondRepo repo.Diamond) (diamond *entity.Diamond, errHttpStatusCode int, err error) {
	if err := authcheck.IsAuthenticated(time.Now(), c); err != nil {
		return nil, http.StatusUnauthorized, fmt.Errorf("authenticate: %w", err)
	}

	id, err := internal.GetIDFromParam(c, "id")
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("get id from param: %w", err)
	}

	diamondUC := usecase.NewDiamondUsecase(diamondRepo)
	diamond = diamondUC.GetDiamondByID(id)
	if diamond == nil {
		return nil, http.StatusNotFound, fmt.Errorf("diamond with id '%d' not found", id)
	}

	return diamond, 0, nil
}

func UpdateDiamond(diamondRepo repo.Diamond) func(*gin.Context) {
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

		input := new(usecase.UpdateDiamondInput)
		if err := c.BindJSON(input); err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Errorf("bind json: %v", err))
			return
		}
		input.ID = id

		diamondUC := usecase.NewDiamondUsecase(diamondRepo)
		diamond, err := diamondUC.UpdateDiamond(input)
		if err != nil {
			if errors.Is(err, usecase.ErrInvalidDiamondInput) ||
				errors.Is(err, entity.ErrInvalidDiamond) {
				response.Error(c, http.StatusBadRequest, err)
				return
			}
			if errors.Is(err, repo.ErrDiamondNotFound) {
				response.Error(c, http.StatusNotFound, err)
				return
			}
			response.Error(c, http.StatusInternalServerError, err)
			return
		}

		response.Success(c, http.StatusCreated, diamond)
	}
}

func DeleteDiamondByID(diamondRepo repo.Diamond) func(*gin.Context) {
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

		diamondUC := usecase.NewDiamondUsecase(diamondRepo)
		diamondUC.DeleteDiamondByID(id)

		response.Success(c, http.StatusOK, true)
	}
}
