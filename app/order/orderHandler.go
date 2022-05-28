package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderStatistics(c *gin.Context) {

	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "noToken", "msg": "OK"})
		return
	}
	result := fetchOrderStatistics(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func DiscountStatistics(c *gin.Context) {

	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "noToken", "msg": "OK"})
		return
	}
	result := fetchDiscount(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func Coupons(c *gin.Context) {
	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "noToken", "msg": "OK"})
		return
	}
	result := fetchDiscount(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func fetchOrderStatistics(token string) interface{} {
	return map[string]string{"token": token, "orderStatistics": "0"}
}

func fetchDiscount(token string) interface{} {
	return map[string]string{"token": token, "discount": "0"}
}
