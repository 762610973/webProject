package logic

import (
	"webProject/dao/mysql"
	"webProject/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库，查找到所有的community
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
