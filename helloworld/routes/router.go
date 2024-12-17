package routes

import (
	"net/http"
	"test.com/helloworld/controller"
	"test.com/helloworld/middlewares"

	"github.com/gin-gonic/gin"
	"test.com/helloworld/logger"
)

func SetUp(mode string) (r *gin.Engine) {
	if mode == gin.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else if mode == gin.TestMode {
		gin.SetMode(gin.TestMode)
	}
	r = gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	{
		r.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "hello world")
		})
		r.POST("/signup", controller.UserSignUpHandler)
		r.POST("/signin", controller.UserSignInHandler)
		r.POST("/auth", middlewares.JWTAuthMiddleware(), controller.AuthHandler)
	}
	return
}
