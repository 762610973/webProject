package logic

import (
	"webProject/dao/mysql"
	"webProject/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetInfo(c *gin.Context, p *models.Info) error {
	uid, ok := c.Get("userID")
	if !ok {
		zap.L().Error("logic.SetInfo(*gin.Context,*models.Info) failed,")
	}
	userID := uid.(int64)
	if err := mysql.SetInfo(userID, p); err != nil {
		zap.L().Error("logic.SetInfo.mysql.SetInfo(userID,*gin.Context) failed", zap.Error(err))
		return err
	}
	return nil
}
