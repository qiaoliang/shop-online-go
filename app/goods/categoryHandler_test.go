package goods

import (
	"bookstore/app/configs"
	"bookstore/app/testutils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type CategoryHandlerSuite struct {
	suite.Suite
	router *gin.Engine
}

func TestCategoryHandlerSuite(t *testing.T) {
	suite.Run(t, new(CategoryHandlerSuite))
}

func (st *CategoryHandlerSuite) SetupSuite() {
	st.router = st.setupTestRouter()
	configs.GetConfigInstance(testutils.GetConfigFileForTest())
	configs.Cfg.Upgrade()
}

func (st *CategoryHandlerSuite) Test_get_category_list() {
	//构建参数
	params := map[string]string{
		"token": "anythingIsOKByNow",
	}
	body := testutils.HttpGet("/v1/shop/goods/category/all", params, st.router)

	exp := `{"code":0,"data":[{"id":1,"name":"DevOps"},{"id":2,"name":"大数据"}],"msg":"OK"}`
	st.Equal(exp, string(body), "should same.")
}

func (st *CategoryHandlerSuite) setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.GET("/shop/goods/category/all", FetchCatalogues)
	return router
}
