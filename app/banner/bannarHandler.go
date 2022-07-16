package ad

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Banners []BannerVM

func FetchBanners(c *gin.Context) {
	bt, err := c.GetQuery("type")
	if !err || bt == "" {
		bt = "indexBanner"
	}
	t, err := c.GetQuery("token")
	if !err || t == "" {
		t = "NoToken"
	}

	result := GetBannerService().FetchBanners(bt, t)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
