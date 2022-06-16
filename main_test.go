package main

import (
	"bookstore/app/configs"
	"os"

	"github.com/stretchr/testify/suite"
)

type ProdConfigTestBeforeRuningSuite struct {
	suite.Suite
}

func (suite *ProdConfigTestBeforeRuningSuite) SetupSuite() {
}
func (s *ProdConfigTestBeforeRuningSuite) Test_init() {
	os.Setenv("SHOPONLINE_CFG_FILE", "config.yaml")
	expConfig := configs.Config{
		IsTestEnv:       false,
		Persistence:     false,
		Host:            "",
		Port:            0,
		DBUser:          "",
		DBPasswd:        "",
		DBAddr:          "",
		DBPort:          0,
		DBName:          "",
		DBMigrationPath: "",
		DBMigrateProto:  "",
		DBMigrateDir:    "",
		StaticPic:       "",
		BannerPath:      "",
		GoodsPath:       "",
		AvatarPath:      "",
	}
	s.EqualValues(expConfig, configs.Cfg)
}
