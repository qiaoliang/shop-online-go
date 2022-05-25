package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func GetShopingCart(c *gin.Context) {

	token := c.Param("token")
	result := cartInfo(token)
	c.JSON(http.StatusOK, gin.H{"code":0,"data": &result,"msg":"OK"})
}

func cartInfo(token string) interface{} {
  	return map[string]string{"token":"iamToken","cartInfo":"cartInfo 0"};

}