package main

import (
	"bookstore/app/configs"
	"bookstore/app/routers"
)

func init() {

}
func main() {
	configs.GetConfigInstance("config.yaml")
	configs.Cfg.Upgrade()
	configs.Cfg.GetMysqlDBConn()
	routers.InitRouter()
}
