package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"webProject/models"

	"go.uber.org/zap"
)

// 将每一步的数据库操作封装成函数
// 等待logic层根据业务需求调用

const secret = "umbrella"

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		// 结果集为空，返回错误
		return err
	}
	if count > 0 {
		//zap.L().Error("用户已存在")
		return ErrorUserExist
	}
	return
}

// InsertUser 保存进数据库
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL入库
	sqlStr := `INSERT INTO user(user_id, username, password) VALUES (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return

}

// 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// Login 登录，插入用户id，姓名，密码
func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `SELECT user_id,username,password FROM user WHERE username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

// GetUserById 根据id获取用户信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `SELECT user_id,username FROM user WHERE user_id = ?`
	err = db.Get(user, sqlStr, uid)
	if err == sql.ErrNoRows {
		zap.L().Error("mysql.GetUserById failed...", zap.Error(err))
	}
	return
}
