package goods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bookstore/app/configs"

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
	err := configs.Cfg.DBConnection().Where("id = ?", c.Param("id")).First(&book).Errors()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book Book
	//Validate Data
	if err := configs.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	//Validate Input
	defer c.Request.Body.Close()
	body, _ := ioutil.ReadAll(c.Request.Body)
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(body), &result)

	configs.DB.Model(&book).Updates(result)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book Book
	fmt.Printf("Delete id =%s\n", c.Param("id"))
	err := configs.Cfg.DBConnection().Where("id = ?", c.Param("id")).First(&book).Errors()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	configs.DB.Delete(book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
