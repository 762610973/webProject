package mysql

import (
	"fmt"
	"webProject/models"

	"go.uber.org/zap"
)

func SetInfo(userID int64, info *models.Info) error {
	sqlStr := `UPDATE user 
				SET username = ?,password = ?,email = ?,gender = ? 
				WHERE user_id = ?`
	exec, err := db.Exec(sqlStr, info.Username, info.Password, info.Email, info.Gender, userID)
	fmt.Println(exec)
	if err != nil {
		zap.L().Error("mysql.info.SetInfo", zap.Error(err))
		return err
	}
	return nil
}
