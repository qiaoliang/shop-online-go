package goods

import (
	"fmt"
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
	var input CreateBookInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Create Book
	book := Book{Title: input.Title, Author: input.Author}
	configs.Cfg.DBConnection().Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	var book Book

	//Validate Data
	bookid := c.Param("id")
	fmt.Printf("id = %v", bookid)
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
	if err := configs.Cfg.DBConnection().Where("id = ?", c.Param("id")).First(&book).Errors(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	//Validate Input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	configs.DB.Model(book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book Book

	//Validate Data
	if err := configs.DB.Where("id = ? ", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	configs.DB.Delete(book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
