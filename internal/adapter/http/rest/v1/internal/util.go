package internal

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIDFromParam(c *gin.Context, paramName string) (int64, error) {
	idRaw := c.Param(paramName)
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid id: '%v'", idRaw)
	}
	return id, nil
}
