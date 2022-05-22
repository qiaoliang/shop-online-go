package controllers

import (
	"bookstore/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchCategorys(c *gin.Context) {
	result := initTestData()
	c.JSON(http.StatusOK, gin.H{"data": &result})
}
func initTestData() []models.Category {
	cates := []models.Category{}
	cate1 := new(models.Category)
	cate1.Id = 1
	cate1.Name = "DevOps"
	subCate1 := models.Category{
		Id:   100,
		Name: "持续交付 1.0",
	}
	subCate2 := models.Category{
		Id:   101,
		Name: "持续交付 2.0",
	}
	cate1.Children = append(cate1.Children, subCate1)
	cate1.Children = append(cate1.Children, subCate2)
	cates = append(cates, *cate1)
	return cates
}
