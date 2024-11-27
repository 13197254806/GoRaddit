package controller

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"

	"test.com/helloworld/service"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"test.com/helloworld/models"
)

func UserSignUpHandler(c *gin.Context) {
	var params models.ParamSignIn
	if err := c.ShouldBindJSON(&params); err != nil {
		zap.L().Error("signup with invalid params: %v", zap.Error(err))
		if _, ok := err.(validator.ValidationErrors); !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	if err := service.UserSignUp(&params); err != nil {
		zap.L().Error("failed in user signup: %v", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "用户已经存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	zap.L().Info("success in user signup: %v", zap.Any("params", params))
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
	return
}
