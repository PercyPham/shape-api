package v1

import (
	"errors"
	"fmt"
	"net/http"
	"shape-api/internal/adapter/http/rest/internal/response"
	"shape-api/internal/common/apperr"
	"shape-api/internal/repo"
	"shape-api/internal/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(userRepo repo.User) func(*gin.Context) {
	return func(c *gin.Context) {
		input := new(usecase.LoginInput)

		if err := c.BindJSON(input); err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Errorf("bind json: %v", err))
			return
		}

		loginUC := usecase.NewLoginUsecase(userRepo)
		accessToken, err := loginUC.Login(time.Now(), input)
		if err != nil {
			err := fmt.Errorf("register usecase: %w", err)
			if errors.Is(err, apperr.ErrApp) {
				response.Error(c, http.StatusBadRequest, err)
				return
			}
			response.Error(c, http.StatusInternalServerError, err)
			return
		}

		response.Success(c, http.StatusOK, gin.H{
			"accessToken": accessToken,
		})
	}
}
