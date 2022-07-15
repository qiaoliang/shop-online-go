package goods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/example/project/app/configs"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FindBooks(c *gin.Context) {

	var books []Book
	configs.Cfg.DBConnection().Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": &books})
}

func CreateBook(c *gin.Context) {
	//Validate Input

	title := c.PostForm("title")
	author := c.PostForm("author")

	//Create Book
	book := Book{Title: title, Author: author}
	configs.Cfg.DBConnection().Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	var book Book
	result := configs.Cfg.DBConnection().Where("id = ?", c.Param("id")).First(&book)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book Book
	//Validate Data
	if err := configs.Cfg.DBConnection().Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	//Validate Input
	defer c.Request.Body.Close()
	body, _ := ioutil.ReadAll(c.Request.Body)
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(body), &result)

	configs.Cfg.DBConnection().Model(&book).Updates(result)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book Book
	fmt.Printf("Delete id =%s\n", c.Param("id"))
	result := configs.Cfg.DBConnection().Where("id = ?", c.Param("id")).First(&book)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	result = configs.Cfg.DBConnection().Delete(book)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can Not Delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
