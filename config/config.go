package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/yangkaiyue/gin-exp/global"
	"path"
)

func InitConf() {

	viper.SetConfigName("conf.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path.Join(global.ProjectPath, "conf"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	setDefault()

	go func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			global.Logger.Info("Config File Changed. ", e.String())
		})
	}()
}

func setDefault() {
	viper.SetDefault("server.debug", "true")
	viper.SetDefault("server.port", 9000)
}
