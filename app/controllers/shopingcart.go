package controllers

import (
	"bookstore/app/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PutIntoCart(c *gin.Context) {
	token := c.PostForm("token")
	goodsId := c.PostForm("goodsId")
	number := c.PostForm("number")
	id64, err1 := strconv.ParseUint(goodsId, 10, 32)
	vlm64, err2 := strconv.ParseUint(number, 10, 32)
	var result *repository.CartInfo
	if err1 != nil || err2 != nil {
		fmt.Println("error format of params for " + number)
		result = nil
	} else {
		result = repository.GetCartsInstance().AddOrderIntoCart(token, uint(id64), uint(vlm64))
	}
	fmt.Printf("PutIntoCart~~after~~~~token is :%v \n", result.Token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func UpdateShoppingCart(c *gin.Context) {

	token := c.PostForm("token")
	id, _ := strconv.Atoi(c.PostForm("key"))
	number, _ := strconv.Atoi(c.PostForm("number"))
	result := repository.GetCartsInstance().AddOrderIntoCart(token, uint(id), uint(number))

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
func GetShopingCart(c *gin.Context) {

	token := c.Query("token")
	cart := repository.GetCartsInstance().GetCartByToken(token)
	var result interface{}
	if cart == nil {
		result = ""
	} else {
		result = cart
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
