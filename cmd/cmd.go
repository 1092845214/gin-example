package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/yangkaiyue/gin-exp/config"
	"github.com/yangkaiyue/gin-exp/cron"
	"github.com/yangkaiyue/gin-exp/global"
	"github.com/yangkaiyue/gin-exp/router"
)

func Start() {

	var initErr error

	// 初始化配置文件
	config.InitConf()

	// 初始化日志组件
	global.Logger = config.InitLogger()

	// 启动后台进程
	cron.CrontabJob()

	// 只连一个库的话可以在这里初始化, 或者在 service 中针对不同请求进行连接
	//// 初始化数据库连接
	//if err := utils.InitMysql(); err != nil {
	//	initErr = utils.AppendErr(initErr, err)
	//}
	//
	//// 初始化redis连接
	//if err := utils.InitRedis(); err != nil {
	//	initErr = utils.AppendErr(initErr, err)
	//}

	// 初始化错误最终处理
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

	global.Logger.Infoln("===== Stop Server =====")
}
