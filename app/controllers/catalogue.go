package controllers

import (
	"net/http"
	"bookstore/app/models"
	"github.com/gin-gonic/gin"
)


func FetchCatalogues(c *gin.Context) {
	token := c.Param("token")

	result := initCataloguesData(token)
	c.JSON(http.StatusOK, gin.H{"code":0,"data": &result,"msg":"OK"})

}
func initCataloguesData(token string) []models.Category {
	cates := []models.Category{}
	cate1 := &models.Category{0,"DevOps"}
	cate2 := &models.Category{1,"大数据"}
	cates = append(cates, *cate1)
	cates = append(cates, *cate2)
	return cates
}
