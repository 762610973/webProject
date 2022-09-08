package controller

import (
	"errors"
	"fmt"
	"net/http"
	"webProject/dao/mysql"
	"webProject/logic"
	"webProject/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	// 判断字段类型和格式是否正确
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			/*c.JSON(http.StatusOK, gin.H{
				"msg": "请求参数有误",
			})*/
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 这里触发了错误校验类型
		/*c.JSON(http.StatusOK, gin.H{
			"msg": errs.Translate(trans),
		})*/
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
	}
	// 手动对参数进行详细的业务规则校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//}
	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
	}

	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
	ResponseSuccess(c, nil)
}

// LoginHandler 登录接口
func LoginHandler(c *gin.Context) {
	// 参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数校验错误，返回响应
			ResponseError(c, CodeInvalidParam)
			/*c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})*/
			return
		}
		/*c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})*/
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2.处理登录逻辑
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
