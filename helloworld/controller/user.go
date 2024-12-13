package controller

import (
	"errors"
	"test.com/helloworld/dao/mysql"
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
		if errs, ok := err.(validator.ValidationErrors); ok {
			ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		} else {
			ResponseErrorWithMsg(c, CodeUnknownError, err.Error())
		}
		return
	}
	if err := service.UserSignUp(&params); err != nil {
		zap.L().Error("failed in user signup: ", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExisted)
		} else {
			ResponseErrorWithMsg(c, CodeUnknownError, err.Error())
		}
		return
	}
	zap.L().Info("success in user signup: %v", zap.Any("params", params))
	ResponseSuccess(c, nil)
	return
}

func UserSignInHandler(c *gin.Context) {
	var params models.ParamSignIn
	if err := c.ShouldBindJSON(&params); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		} else {
			ResponseErrorWithMsg(c, CodeUnknownError, err.Error())
		}
		return
	}
	if err := service.UserSignIn(&params); err != nil {
		zap.L().Error("failed in user signin: ", zap.Error(err))
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeUserNotExisted)
		} else {
			ResponseErrorWithMsg(c, CodeUnknownError, err.Error())
		}
		return
	}
	zap.L().Info("success in user signin: %v", zap.Any("params", params))
	ResponseSuccess(c, nil)
}
