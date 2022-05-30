package ad

import (
	"bookstore/app/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchBanners(c *gin.Context) {
	result := initBannerData()
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
func initBannerData() []Banner {
	bans := []Banner{}
	ban1 := &Banner{
		0,
		"2022-05-05 11:26:09",
		222083,
		"https://gitee.com/joeshu/v-shop",
		0,
		configs.Cfg.BannerPicPrefix() + "b7225946.jpeg",
		"跳转gitee v-shop",
		0,
		"q",
		"p",
		"s",
		1605,
	}

	ban2 := &Banner{
		1,
		"2022-05-05 11:26:09",
		222084,
		"https://gitee.com/joeshu/v-shop",
		0,
		configs.Cfg.BannerPicPrefix() + "b7225947.jpeg",
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
