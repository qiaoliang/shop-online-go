package controllers

import (
    "bookstore/app/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

func FetchGoodsList(c *gin.Context) {
	// params
	page := c.PostForm("page")
    pageSize := c.DefaultPostForm("pageSize", "10")

    // logic
	result := initGoodsList(page,pageSize)

	// response
	c.JSON(http.StatusOK, gin.H{
		"code":0,
		"data":gin.H{
			"totalRow":len(result),
			"result":&result,
		},
		"msg":"OK",
	})


}
func initGoodsList(page string,pagesize string) [] models.GoodsItem {
	items := []models.GoodsItem{}
	items1 := &models.GoodsItem{
		0,        			//id
		"持续交付 1.0 ",		//name
		"1", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225946.jpeg",//picURL
		"66.0",				//MinPrice
		"99.0",				//originalPrice
	}
	items2 := &models.GoodsItem{
		1,        			//id
		"持续交付 2.0",	//name
		"2", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225947.jpeg",//picURL
		"77.0",				//MinPrice
		"109.0",			//originalPrice
	}
	items3 := &models.GoodsItem{
		2,        			//id
		"DevOps 实战指南",	//name
		"3", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225948.jpeg",//picURL
		"60.0",				//MinPrice
		"89.0",			//originalPrice
	}
	items4 := &models.GoodsItem{
		3,        			//id
		"谷歌软件工程",	//name
		"4", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225949.jpeg",//picURL
		"80.0",				//MinPrice
		"129.0",			//originalPrice
	}
	items = append(items, *items1)
	items = append(items, *items2)
	items = append(items, *items3)
	items = append(items, *items4)
	return items
}