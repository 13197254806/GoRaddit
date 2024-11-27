package controller

import (
	"fmt"
	"net/http"
	"test.com/helloworld/service"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"test.com/helloworld/models"
)

func UserSignUpHandler(c *gin.Context) {
	var params models.ParamSignUp
	if err := c.ShouldBindJSON(&params); err != nil {
		//zap.L().Error("signup with invalid params: %v", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	if err := service.UserSignUp(&params); err != nil {
		zap.L().Error("failed in user signup: ", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("注册失败：%s", err.Error()),
		})
		return
	}
	zap.L().Info("success in user signup: %v", zap.Any("params", params))
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
	return
}
