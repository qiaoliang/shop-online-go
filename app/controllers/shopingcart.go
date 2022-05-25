package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func PutIntoCart(c *gin.Context) {

	token := c.Param("token")
	goodsId := c.Param("goodsId")
	number := c.Param("number")


	result := cartInfo(token,goodsId,number)

	c.JSON(http.StatusOK, gin.H{"code":0,"data": &result,"msg":"OK"})
}


func GetShopingCart(c *gin.Context) {

	token := c.Param("token")
	result := cartInfo(token,"","")
	c.JSON(http.StatusOK, gin.H{"code":0,"data": &result,"msg":"OK"})
}

func cartInfo(token string,goodsId string, number string) interface{} {
  	return map[string]string{"token":"iamToken","cartInfo":"cartInfo 0"};

}