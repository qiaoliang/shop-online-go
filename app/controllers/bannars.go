package controllers

import (
	"bookstore/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchBanners(c *gin.Context) {
	result := initBannerData()
	c.JSON(http.StatusOK, gin.H{"data": &result})
}
func initBannerData() models.Responsable {
	bans := []models.Banner{}
	ban1 := &models.Banner{
		0,
		"2022-05-05 11:26:09",
		222083,
		"https://gitee.com/joeshu/v-shop",
		0,
		"https://dcdn.it120.cc/2022/05/05/ac956ae3-151f-418e-b0e9-fadd76a9ea6d.jpeg",
		"跳转gitee v-shop",
		0,
		"q",
		"p",
		"s",
		1605,
	}

	bans = append(bans, *ban1)

	res := models.Responsable{0,bans,"OK"}
	return res
}