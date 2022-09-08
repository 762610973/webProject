package mysql

import (
	"database/sql"
	"webProject/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `SELECT community_id, community_name FROM community`
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			// 空记录
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (*models.CommunityDetail, error) {
	community := new(models.CommunityDetail)
	sqlStr := `SELECT 
				community_id, community_name, introduction,create_time 
				FROM community 
				WHERE community_id = ?`
	var err error
	if err = db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}
