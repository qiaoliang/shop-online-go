package goods

import (
	"net/http"
	"strconv"

	"bookstore/app/testutils"

	"github.com/gin-gonic/gin"
)

type GoodsListJson struct {
	testutils.JsonResult
	Data *ResultData `json:"data"`
}
type ResultData struct {
	TotalRow int         `json:"totalRow"`
	Result   []GoodsItem `json:"result"`
}

type GoodsHandler struct {
	service *GoodsService
}

func NewGoodsHandler(service *GoodsService) *GoodsHandler {
	return &GoodsHandler{service: service}
}

func (h *GoodsHandler) GetGoodsDetail(c *gin.Context) {
	gid, _ := c.GetQuery("id")
	result := h.service.GetItemDetail(gid)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": &result,
		"msg":  "OK",
	})
}

func (h *GoodsHandler) FetchGoodsList(c *gin.Context) {
	page := c.PostForm("page")
	pageSize := c.DefaultPostForm("pageSize", "10")
	categoryId := c.PostForm("categoryId")
	catalogueId, _ := strconv.Atoi(categoryId)
	cateId := uint(catalogueId)
	result := h.getGoods(page, pageSize, cateId)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"totalRow": len(result),
			"result":   &result,
		},
		"msg": "OK",
	})
}

func (h *GoodsHandler) FetchCatalogues(c *gin.Context) {
	// 使用 service 中的 cateRepo 获取分类数据
	result := h.service.cateRepo.GetList()
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func (h *GoodsHandler) FetchItemReputation(c *gin.Context) {
	// TODO: 实现商品评价查询逻辑
	c.JSON(200, gin.H{"code": 0, "data": []string{}, "msg": "OK"})
}

func (h *GoodsHandler) getGoods(page string, pageSize string, catalogueId uint) []GoodsItem {
	skus := h.service.repo.FindAll()
	items := make([]GoodsItem, 0)
	for _, sku := range skus {
		if sku.CategoryId == catalogueId {
			i := h.service.skuToGoodsItem(sku)
			items = append(items, *i)
		}
	}
	return items
}
