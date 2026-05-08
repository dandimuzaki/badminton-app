package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUintParam(c *gin.Context, name string) (uint, error) {
	idStr := c.Param(name)

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, errors.New("invalid id")
	}

	return uint(id64), nil
}