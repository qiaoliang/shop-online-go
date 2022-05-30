package goods

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type GoodsHandlerSuite struct {
	suite.Suite
	router *gin.Engine
}

func TestGoodsHandlerSuite(t *testing.T) {
	suite.Run(t, new(GoodsHandlerSuite))
}

func (st *GoodsHandlerSuite) SetupSuite() {
	st.router = st.setupTestRouter()
	configs.GetConfigInstance(utils.GetConfigFileForTest())
}

func (st *GoodsHandlerSuite) Test_get_GoodsDetail() {
	//构建参数
	data := url.Values{}
	data.Set("token", "anythingIsOKByNow")
	data.Add("id", "1111111")
	//构建返回值
	//调用请求接口

	body := utils.HttpRequest(st.router, data, "GET", "/v1/shop/goods/detail")

	exp := `{"code":0,"data":{"id":0,"gid":"g7225946","name":"持续交付1.0","pics":[{"id":"g7225946-01","pic":"http://localhost:9090/pic/goods/g7225946-01.jpeg"},{"id":"g7225946-02","pic":"http://localhost:9090/pic/goods/g7225946-02.jpeg"}],"goodsId":0,"stores":10,"unit":"册","logistics":"0","content":"一本DevOps的经典书。","status":2,"statusStr":"在售","pic":"http://localhost:9090/pic/goods/g7225946.jpeg","minPrice":"66.0","originalPrice":"99.0","afterSale":"1"},"msg":"OK"}`
	st.Equal(exp, string(body), "should same.")
}

func (st *GoodsHandlerSuite) setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.GET("/shop/goods/detail", GetGoodsDetail)
	return router
}
