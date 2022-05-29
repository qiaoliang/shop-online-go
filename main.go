package main

import (
	"bookstore/app/configs"
	"bookstore/app/routers"
)

func init() {

}
func main() {
	configs.NewConfig("config.yaml")
	configs.Cfg.DbMigrate()
	configs.Cfg.InitMysqlDB()
	routers.InitRouter()
}
