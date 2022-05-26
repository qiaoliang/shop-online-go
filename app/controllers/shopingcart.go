package controllers

import (
	"bookstore/app/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PutIntoCart(c *gin.Context) {

	token := c.Param("token")
	goodsId := c.Param("goodsId")
	number := c.Param("number")
	id64, err1 := strconv.ParseUint(goodsId, 10, 32)
	vlm64, err2 := strconv.ParseUint(number, 10, 32)
	var result *repository.Cart
	if err1 == nil && err2 == nil {
		result = repository.GetCartsInstance().AddOrderIntoCart(token, uint(id64), uint(vlm64))
	} else {
		fmt.Println("error format of params for " + number)
		result = nil
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": result, "msg": "OK"})
}

func GetShopingCart(c *gin.Context) {

	token := c.Param("token")
	result := cartInfo(token, "", "")
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func cartInfo(token string, goodsId string, number string) interface{} {
	return map[string]string{"token": "iamToken", "cartInfo": "cartInfo 0"}

}
