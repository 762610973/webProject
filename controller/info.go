package controller

import (
	"webProject/logic"
	"webProject/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetInfoHandler(c *gin.Context) {
	// 获取参数，简单校验，然后直接存到数据库里
	p := new(models.Info)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("controller.SetInfoHandler", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.SetInfo(c, p); err != nil {
		zap.L().Error("controller.SetInfoHandler.logic.SetInfo", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
