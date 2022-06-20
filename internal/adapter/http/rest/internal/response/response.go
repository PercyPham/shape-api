package response

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, successResponse{data})
}

type successResponse struct {
	Data interface{} `json:"data"`
}

func Error(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, errorResponse{err.Error()})
}

type errorResponse struct {
	Error string `json:"error"`
}
