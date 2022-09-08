package main

// @title webProject文档
// @version 1.0
// @description webProject项目

// @contact.name umbrella

// @host 127.0.0.1:1112
// @BasePath /api/v1

import (
	"fmt"
	"os"
	"webProject/controller"
	"webProject/dao/mysql"
	"webProject/dao/redis"
	"webProject/logger"
	"webProject/pkg/snowflake"
	"webProject/router"
	"webProject/setting"

	"go.uber.org/zap"
)

func main() {
	/*
		1.加载配置settings
		2.初始化日志
		3.初始化MySQL连接
		4.初始化Redis连接
		5.注册路由
		6.启动服务（优雅关机）
	*/
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: webProject config.yaml")
		return
	}
	// 加载配置
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	// 初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			fmt.Printf("zap.L().Sync() failed, err:%v\n", err)
		}
	}(zap.L())
	// 连接MySQL
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	// 连接Redis
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := router.SetUpRouter(setting.Conf.Mode)
	fmt.Println()
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
