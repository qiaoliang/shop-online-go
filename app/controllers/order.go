package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func GetOrderStatistics(c *gin.Context) {

	token := c.Param("token")
	result := fetchOrderStatistics(token)
	c.JSON(http.StatusOK, gin.H{"code":0,"data": &result,"msg":"OK"})
}


func DiscountStatistics(c *gin.Context) {

	token := c.Param("token")
	result := fetchDiscount(token)
	c.JSON(http.StatusOK, gin.H{"code":0,"data": &result,"msg":"OK"})
}

func Coupons(c *gin.Context) {
	token := c.Param("token")
	result := fetchDiscount(token)
	c.JSON(http.StatusOK, gin.H{"code":0,"data": &result,"msg":"OK"})
}

func fetchOrderStatistics(token string)interface{}{
	return map[string]string{"token":"iamToken","orderStatistics":"amount 0"};
}

func fetchDiscount(token string)interface{}{
	return map[string]string{"token":"iamToken","discount":"amount 0"};
}