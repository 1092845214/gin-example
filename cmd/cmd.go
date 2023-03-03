package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/yangkaiyue/gin-exp/config"
	"github.com/yangkaiyue/gin-exp/global"
	"github.com/yangkaiyue/gin-exp/router"
	"github.com/yangkaiyue/gin-exp/utils"
)

func Start() {

	var initErr error

	// 初始化配置文件
	config.InitConf()

	// 初始化日志组件
	global.Logger = config.InitLogger()

	// 初始化数据库连接
	db, err := config.InitDB()
	global.DB = db

	// 初始化错误处理
	initErr = utils.AppendErr(initErr, err)
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Errorln(initErr.Error())
		}
		panic(initErr.Error())
	}

	// gin 启动模式
	gin.SetMode("release")
	if viper.GetBool("server.debug") {
		gin.SetMode("debug")
	}

	// 初始化路由
	router.InitRouter()
}

func Clean() {

	global.Logger.Infoln("=======clean==========")
}
