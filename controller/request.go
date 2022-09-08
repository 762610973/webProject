package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// 获取分页数据
func getPageInfo(c *gin.Context) (int64, int64) {
	offsetStr := c.Query("page")
	sizeStr := c.Query("size")
	// 获取数据
	var (
		page int64
		size int64
		err  error
	)
	size, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		size = 1
	}
	page, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		page = 10
	}
	return page, size
}
