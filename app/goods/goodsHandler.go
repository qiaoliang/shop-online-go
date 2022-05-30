package goods

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchGoodsList(c *gin.Context) {
	// params
	page := c.PostForm("page")
	pageSize := c.DefaultPostForm("pageSize", "10")
	categoryId := c.PostForm("categoryId")
	catalogueId, _ := strconv.Atoi(categoryId)
	id := uint(catalogueId)

	// logic
	result := getGoods(page, pageSize, id)

	// response
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"totalRow": len(result),
			"result":   &result,
		},
		"msg": "OK",
	})
}

func GetGoodsDetail(c *gin.Context) {

	// params
	token, _ := c.GetQuery("token")
	idStr, _ := c.GetQuery("id")
	idUint, _ := strconv.Atoi(idStr)
	id := uint(idUint)
	fmt.Printf(" goods detail token =%v, id=%v\n\n\n", token, id)
	result := getItemDetail(id, token)

	// response
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": &result,
		"msg":  "OK",
	})
}

func getItemDetail(id uint, token string) GoodsDetail {
	gr := GetGoodsRepo()
	gr.loadGoods()
	return gr.getItemDetail(id, token)
}

func getGoods(page string, pageSize string, catalogueId uint) []GoodsItem {
	gr := GetGoodsRepo()
	goods := gr.loadGoods()
	result := make([]GoodsItem, 0)
	for _, item := range goods {
		if item.blongsTo(catalogueId) {
			result = append(result, item)
		}
	}
	return result
}
