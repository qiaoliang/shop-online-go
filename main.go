package main

import (
	"bookstore/app/configs"
	"bookstore/app/routers"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()
	configs.DbMigrate()
	configs.InitMysqlDB()
	routers.InitRouter()
}
