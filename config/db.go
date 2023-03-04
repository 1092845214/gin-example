package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yangkaiyue/gin-exp/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitMysql() (err error) {

	// dsn 相关配置参数
	// https://github.com/go-sql-driver/mysql#parameters
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=30s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.db_name"),
	)

	logMode := logger.Error
	if viper.GetBool("server.debug") {
		logMode = logger.Info
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "epoch_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return
	}

	db2, _ := db.DB()
	db2.SetMaxOpenConns(viper.GetInt("mysql.max_idle_conn"))
	db2.SetMaxIdleConns(viper.GetInt("mysql.max_open _conn"))
	db2.SetConnMaxLifetime(time.Hour)

	// 如果表不存在则创建
	//db.AutoMigrate(&model.User{})

	// 赋值,如果
	global.DB = db
	return
}
