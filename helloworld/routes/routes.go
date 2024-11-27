package routes

import (
	"net/http"

	"test.com/helloworld/controller"

	"github.com/gin-gonic/gin"
	"test.com/helloworld/logger"
)

func SetUp() (r *gin.Engine) {
	r = gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	{
		r.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "hello world")
		})
		r.POST("/signup", controller.UserSignUpHandler)
	}
	return
}
