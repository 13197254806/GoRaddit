package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthHandler(c *gin.Context) {
	userId, exists := c.Get("UserId")
	if exists {
		fmt.Print(userId)
		ResponseSuccess(c, userId)
		return
	}
	ResponseError(c, CodeInvalidToken)
}
