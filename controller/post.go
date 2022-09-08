package controller

import (
	"strconv"
	"webProject/logic"
	"webProject/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 创建帖子/文章
func CreatePostHandler(c *gin.Context) {
	// 1.获取参数及参数的校验
	p := new(models.Post)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("controller.Post.BindJSON", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从c中获取当前用的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		// 解析不出来，重新登陆一下
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情
func GetPostDetailHandler(c *gin.Context) {
	// 1.获取参数（从URL中获取帖子id）
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
	}
	// 2.根据id查询帖子数据（查数据库）
	data, err := logic.GetPostById(id)
	if err != nil {
		zap.L().Error("logic.GetPostById(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表
func GetPostListHandler(c *gin.Context) {
	page, size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler2 获取帖子列表根据前端传来的参数（按照分数、创建时间）动态获取帖子列表
// 解析参数、从Redis中查询id列表、根据id去数据库查询帖子详细信息
func GetPostListHandler2(c *gin.Context) {
	// get请求参数：/api/v1/posts2?page=1&size=10&order=time
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostListHandler2() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

func GetCommunityPostListHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostListNew(p)
	//data, err := logic.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("logic.GetCommunityPostListHandler() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}
