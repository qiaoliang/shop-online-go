package main

import (
	"github.com/example/project/app/configs"
	"github.com/example/project/app/routers"
)

func init() {

}
func main() {
	configs.GetConfigInstance("config.yaml")
	if configs.Cfg.Persistence {
		configs.Cfg.Upgrade()
		configs.Cfg.DBConnection()
	}
	routers.InitRouter()
}
