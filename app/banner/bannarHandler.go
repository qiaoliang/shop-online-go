package banner

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Banners 是 BannerVM 的切片类型
type Banners = []BannerVM

type BannerHandler struct {
	service *BannerService
}

func NewBannerHandler(service *BannerService) *BannerHandler {
	return &BannerHandler{service: service}
}

func (h *BannerHandler) FetchBanners(c *gin.Context) {
	bt, err := c.GetQuery("type")
	if !err || bt == "" {
		bt = "indexBanner"
	}
	t, err := c.GetQuery("token")
	if !err || t == "" {
		t = "NoToken"
	}
	result := h.service.FetchBanners(bt, t)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": result, "msg": "OK"})
}
