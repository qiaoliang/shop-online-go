package controllers

import (
	"bookstore/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchBanners(c *gin.Context) {
	result := initBannerData()
	c.JSON(http.StatusOK, gin.H{"code":0,"data": &result,"msg":"OK"})
}
func initBannerData() []models.Banner {
	bans := []models.Banner{}
	ban1 := &models.Banner{
		0,
		"2022-05-05 11:26:09",
		222083,
		"https://gitee.com/joeshu/v-shop",
		0,
		"http://localhost:9090/pic/banners/b7225946.jpeg",
		"跳转gitee v-shop",
		0,
		"q",
		"p",
		"s",
		1605,
	}


	ban2 := &models.Banner{
		1,
		"2022-05-05 11:26:09",
		222084,
		"https://gitee.com/joeshu/v-shop",
		0,
		"http://localhost:9090/pic/banners/b7225947.jpeg",
		"跳转gitee v-shop",
		0,
		"q",
		"p",
		"s",
		1606,
	}
	bans = append(bans, *ban1)
	bans = append(bans, *ban2)
	return bans
}