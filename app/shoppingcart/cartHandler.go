package cart

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PutIntoCart(c *gin.Context) {
	token := c.PostForm("token")
	gid := c.PostForm("goodsId")
	number := c.PostForm("number")
	vlm64, err2 := strconv.ParseUint(number, 10, 32)
	var result *CartInfo
	if err2 != nil {
		fmt.Println("error format of params for " + number)
		result = nil
	} else {
		result = GetCartsInstance().AddOrderIntoCart(token, gid, uint(vlm64))
	}
	fmt.Printf("PutIntoCart~~after~~~~token is :%v \n", result.Token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func UpdateShoppingCart(c *gin.Context) {

	token := c.PostForm("token")
	gid := c.PostForm("gid")
	number, _ := strconv.Atoi(c.PostForm("number"))
	result := GetCartsInstance().AddOrderIntoCart(token, gid, uint(number))

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
func GetShopingCart(c *gin.Context) {

	token := c.Param("token")
	cart := GetCartsInstance().GetCartByToken(token)
	var result interface{}
	if cart == nil {
		result = ""
	} else {
		result = cart
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
