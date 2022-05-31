package goods

import (
	"bookstore/app/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoodsListJson struct {
	utils.JsonResult
	Data *ResultData `json:"data"`
}
type ResultData struct {
	TotalRow int         `json:"totalRow"`
	Result   []GoodsItem `json:"result"`
}

func FetchItemReputation(c *gin.Context) {
	//TODO have not implemented,please fix it.
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": "",
		"msg":  "OK",
	})
}

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
	gid, _ := c.GetQuery("id")
	fmt.Printf(" goods detail token =%v, id=%v\n\n\n", token, gid)
	result := getItemDetail(gid, token)

	// response
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": &result,
		"msg":  "OK",
	})
}

func getItemDetail(gid string, token string) GoodsDetail {
	gr := GetGoodsRepo()
	gr.loadGoods()
	return gr.getItemDetail(gid, token)
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
