package logic

import (
	"webProject/dao/mysql"
	"webProject/models"
	"webProject/pkg/jwt"
	"webProject/pkg/snowflake"

	"go.uber.org/zap"
)

// 存放业务逻辑的代码

// SignUp 注册
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		return err
	}
	// 2.生成UID
	userID := snowflake.GenID()
	u := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进数据库
	err = mysql.InsertUser(&u)
	if err != nil {
		zap.L().Error("数据库存储出错", zap.Error(err), zap.String("username", p.Username))
		return err
	}
	return nil
	/*return mysql.InsertUser(&u)*/
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
