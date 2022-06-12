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
	if token == "" || gid == "" || number == "" || err2 != nil {
		result := fmt.Sprintf("token =%v, gid = %v, number =%v\n ", token, gid, number)
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": result, "msg": "OK"})
		return
	}
	result := GetCartsRepo().PutItemsInCart(token, gid, uint(vlm64))
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func ModifyNumberOfGoodsInCart(c *gin.Context) {

	token := c.PostForm("token")
	gid := c.PostForm("key") //只有这里用了Key，其它都用了gid，或 goodsId
	numStr := c.PostForm("number")

	number, _ := strconv.Atoi(numStr)

	GetCartsRepo().ModifyQuantityOfGoodsInCate(token, gid, uint(number))
	result := GetCartsRepo().GetCartByToken(token)

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
func GetShopingCart(c *gin.Context) {

	token, err := c.GetQuery("token")
	if !err {
		fmt.Println("can not Parse token。")
	}
	cart := GetCartsRepo().GetCartByToken(token)
	var result interface{}
	if cart == nil {
		fmt.Println("没有找到 token 为 " + token + " 的购物车")
		result = ""
	} else {
		result = cart
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
