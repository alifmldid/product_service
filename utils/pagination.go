package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetPaginationParams(c *gin.Context) (limit, offset int) {
	limit = 10
	offset = 0

	if l := c.Query("limit"); l != "" {
		fmt.Sscan(l, &limit)
	}

	if o := c.Query("offset"); o != "" {
		fmt.Sscan(o, &offset)
	}

	return
}
