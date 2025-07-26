package goods

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewCategoryHandler(cateRepo *CategoryRepoDB) *CategoryHandler {
	return &CategoryHandler{cateRepo: cateRepo}
}

type CategoryHandler struct {
	cateRepo *CategoryRepoDB
}

func  (cr *CategoryHandler) FetchCatalogues(c *gin.Context) {
	result := cr.cateRepo.LoadCategory()
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
