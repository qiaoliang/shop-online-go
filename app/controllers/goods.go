package controllers

import (
    "bookstore/app/models"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func FetchGoodsList(c *gin.Context) {
	// params
	page := c.PostForm("page")
    pageSize := c.DefaultPostForm("pageSize", "10")
    categoryId := c.PostForm("categoryId")
    catalogueId,_:= strconv.Atoi(categoryId)
    id := uint(catalogueId)

    // logic
	result := getGoods(page,pageSize,id)

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

func GetGoodsDetail(c *gin.Context) {

	// params
	token := c.PostForm("token")
    idStr := c.PostForm("id")
    idUint,_:= strconv.Atoi(idStr)
    id := uint(idUint)

    result:=getItem(id,token)


    // response
	c.JSON(http.StatusOK, gin.H{
		"code":0,
		"data":&result,
		"msg":"OK",
	})
}

func getItem(id uint,token string) models.GoodsItem {
	goods:=initGoodsList();
	for _, item := range goods {
		if sameAs(id,item){
			return item;
		}
	}
	return models.GoodsItem{}
}

func sameAs(id uint, item models.GoodsItem) bool{
	return id == item.Id
}

func getGoods(page string,pageSize string,catalogueId uint) [] models.GoodsItem {
	goods:=initGoodsList();
	result:=goods[:0] //我们利用传过来的slice重新创建一个slice，底层不会重新创建数组
	for _, item := range goods {
		if isA(catalogueId,item) {
			result = append(result, item)
		}
	}
	return result
}

func isA(catalogueId uint,item models.GoodsItem) bool {
	return catalogueId == item.CatalogueId
}

func initGoodsList() [] models.GoodsItem {
	items := []models.GoodsItem{}
	items1 := &models.GoodsItem{
		0,        			//id
		"持续交付 1.0 ",		//name
		0,        			//catalogueId
		"1", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225946.jpeg",//picURL
		"66.0",				//MinPrice
		"99.0",				//originalPrice
	}
	items2 := &models.GoodsItem{
		1,        			//id
		"持续交付 2.0",	//name
		0,        			//catalogueId
		"2", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225947.jpeg",//picURL
		"77.0",				//MinPrice
		"109.0",			//originalPrice
	}
	items3 := &models.GoodsItem{
		2,        			//id
		"DevOps 实战指南",	//name
		0,        			//catalogueId
		"3", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225948.jpeg",//picURL
		"60.0",				//MinPrice
		"89.0",			//originalPrice
	}
	items4 := &models.GoodsItem{
		3,        			//id
		"谷歌软件工程",	//name
		0,        			//catalogueId
		"4", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225949.jpeg",//picURL
		"80.0",				//MinPrice
		"129.0",			//originalPrice
	}
	items = append(items, *items1)
	items = append(items, *items2)
	items = append(items, *items3)
	items = append(items, *items4)


	// add the second catalouge

	items5 := &models.GoodsItem{
		10,        			//id
		"驾驭大数据",		//name
		1,        			//catalogueId
		"1", 				//recommandStatus
		"http://localhost:9090/pic/goods/g1872110.jpeg",//picURL
		"50.0",				//MinPrice
		"75.0",				//originalPrice
	}
	items6 := &models.GoodsItem{
		11,        			//id
		"数据分析变革",		//name
		1,        			//catalogueId
		"2", 				//recommandStatus
		"http://localhost:9090/pic/goods/g1872111.jpeg",//picURL
		"41.0",				//MinPrice
		"65.0",			//originalPrice
	}
	items7 := &models.GoodsItem{
		12,        			//id
		"大数据测试技术与实践",	//name
		1,        			//catalogueId
		"3", 				//recommandStatus
		"http://localhost:9090/pic/goods/g1872112.jpeg",//picURL
		"60.0",				//MinPrice
		"89.0",			//originalPrice
	}
	items8 := &models.GoodsItem{
		13,        			//id
		"图解Spark 大数据快速分析实战",	//name
		1,        			//catalogueId
		"4", 				//recommandStatus
		"http://localhost:9090/pic/goods/g1872113.jpeg",//picURL
		"80.0",				//MinPrice
		"129.0",			//originalPrice
	}
	items = append(items, *items5)
	items = append(items, *items6)
	items = append(items, *items7)
	items = append(items, *items8)
	return items
}