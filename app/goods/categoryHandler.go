package goods

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchCatalogues(c *gin.Context) {
	cr := GetCategoryRepo()
	result := cr.loadCategory()
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})

}
