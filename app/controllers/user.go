package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	result := initUserData()
	c.JSON(http.StatusOK, gin.H{"data": &result})
}

func initUserData() string {
	return "ok";
}
