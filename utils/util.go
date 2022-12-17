package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetLimitAndOffset(c *gin.Context) (int, int) {
	var limit, offset int
	if c.Query("limit") == "" {
		limit = 10
	} else {
		limit, _ = strconv.Atoi(c.Query("limit"))
	}
	if c.Query("offset") == "" {
		offset = 0
	} else {
		offset, _ = strconv.Atoi(c.Query("offset"))
	}

	return limit, offset
}
