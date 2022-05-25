package controllers

import (
    "bookstore/app/repository"
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

    result:=getItemDetail(id,token)


    // response
	c.JSON(http.StatusOK, gin.H{
		"code":0,
		"data":&result,
		"msg":"OK",
	})
}

func getItemDetail(id uint,token string) models.GoodsDetail {
	goods:=repository.InitGoodsList();

	for _, item := range goods {
		if sameAs(id,item){
			return item.GoodsDetail;
		}
	}
	return models.GoodsDetail{}
}

func sameAs(id uint, item models.GoodsItem) bool{
	return id == item.Id
}

func getGoods(page string,pageSize string,catalogueId uint) [] models.GoodsItem {
	goods:=repository.InitGoodsList();
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