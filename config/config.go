package config

import (
	"github.com/spf13/viper"
	"os"
	"path"
)

var (
	exePath, _  = os.Executable()
	projectPath = path.Dir(path.Dir(exePath))
)

func InitConf() {

	viper.SetConfigName("conf.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path.Join(projectPath, "conf"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
}
