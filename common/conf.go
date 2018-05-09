package common

import (
	"os"

	"github.com/astaxie/beego/config"
)

var (
	Sconfig config.Configer
)

func init() {
	fileDir, _ := os.Getwd()
	path := fileDir + "/config/redis.ini"
	_, err := os.Stat(path)

	if err != nil {
		panic("config file not found")
	}
	conf, err := config.NewConfig("ini", path)
	if err != nil {
		panic("config init faild")
	}
	Sconfig = conf
}

func GetConfig(key string) string {
	return Sconfig.String(key)
}
