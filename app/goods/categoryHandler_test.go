package goods

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
	"net/url"
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
	configs.NewConfig(utils.GetConfigFileForTest())
}

func (st *CategoryHandlerSuite) Test_get_category_list() {
	//构建参数
	data := url.Values{}
	data.Set("token", "anythingIsOKByNow")
	//构建返回值
	//调用请求接口
	body := utils.HttpRequest(st.router, data, "GET", "/v1/shop/goods/category/all")

	exp := `{"code":0,"data":[{"id":0,"name":"DevOps"},{"id":1,"name":"大数据"}],"msg":"OK"}`
	st.Equal(exp, string(body), "should same.")
}

func (st *CategoryHandlerSuite) setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.GET("/shop/goods/category/all", FetchCatalogues)
	return router
}
