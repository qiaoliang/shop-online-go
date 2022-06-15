package goods

import (
	"bookstore/app/testutils"
	"net/http"
	"strconv"

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
	cateId := uint(catalogueId)

	// logic
	result := getGoods(page, pageSize, cateId)

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
	gid, _ := c.GetQuery("id")
	result := getItemDetail(gid)

	// response
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": &result,
		"msg":  "OK",
	})
}

func getItemDetail(gid string) *GoodsDetail {
	gs := GetGoodsService()
	gs.LoadGoods()
	return gs.GetItemDetail(gid)
}

func getGoods(page string, pageSize string, catalogueId uint) []GoodsItem {
	gs := GetGoodsService()
	gs.LoadGoods()
	return gs.GetCategory(catalogueId)
}
