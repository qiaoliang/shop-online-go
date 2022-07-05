package ad

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Banners []Banner

func FetchBanners(c *gin.Context) {
	bt, err := c.GetQuery("type")
	if !err {
		bt = "indexBanner"
	}
	t, err := c.GetQuery("token")
	if !err {
		t = "NoToken"
	}

	result := GetBannerService().FetchBanners(bt, t)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
