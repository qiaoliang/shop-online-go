package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	result := initUserData()
	c.JSON(http.StatusOK, gin.H{"code":0,"data": &result,"msg":"OK"})
}

func initUserData() interface{} {
	return map[string]string{"token":"iamToken",}
}
