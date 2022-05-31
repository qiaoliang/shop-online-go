package goods

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type GoodItemJson struct {
	Data *GoodsItem `json:"data"`
	utils.JsonResult
}

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
	data.Add("id", "g7225946")
	params := map[string]string{
		"token": "anythingIsOKByNow",
		"id":    "g7225946",
	}
	//构建返回值
	//调用请求接口
	exp := `{"code":0,"data":{"id":"g7225946","name":"持续交付1.0","pics":[{"id":"g7225946-01","pic":"http://localhost:9090/pic/goods/g7225946-01.jpeg"},{"id":"g7225946-02","pic":"http://localhost:9090/pic/goods/g7225946-02.jpeg"}],"goodsId":0,"stores":10,"unit":"册","logistics":"0","content":"一本DevOps的经典书。","status":2,"statusStr":"在售","pic":"http://localhost:9090/pic/goods/g7225946.jpeg","minPrice":"66.0","originalPrice":"99.0","afterSale":"1"},"msg":"OK"}`

	body := utils.HttpGet("/v1/shop/goods/detail", params, st.router)
	st.Equal(exp, body, "should same.")
}
func (st *GoodsHandlerSuite) Test_fetch_GoodsList() {
	//构建参数
	data := url.Values{}
	data.Set("page", "")
	data.Add("pageSize", "anythingIsOKByNow")
	data.Add("categoryId", "0")

	//构建返回值
	//调用请求接口
	exp := `{"code":0,"data":{"result":[{"id":"g7225946","name":"持续交付1.0","catalogueId":0,"recommendStatus":"1","pic":"http://localhost:9090/pic/goods/g7225946.jpeg","minPrice":"66.0","originalPrice":"99.0","goodsDetail":{"id":"g7225946","name":"持续交付1.0","pics":[{"id":"g7225946-01","pic":"http://localhost:9090/pic/goods/g7225946-01.jpeg"},{"id":"g7225946-02","pic":"http://localhost:9090/pic/goods/g7225946-02.jpeg"}],"goodsId":0,"stores":10,"unit":"册","logistics":"0","content":"一本DevOps的经典书。","status":2,"statusStr":"在售","pic":"http://localhost:9090/pic/goods/g7225946.jpeg","minPrice":"66.0","originalPrice":"99.0","afterSale":"1"}},{"id":"g7225947","name":"持续交付2.0","catalogueId":0,"recommendStatus":"1","pic":"http://localhost:9090/pic/goods/g7225947.jpeg","minPrice":"99.0","originalPrice":"129.0","goodsDetail":{"id":"g7225947","name":"持续交付2.0","pics":[{"id":"g7225947-01","pic":"http://localhost:9090/pic/goods/g7225947-01.jpeg"},{"id":"g7225947-02","pic":"http://localhost:9090/pic/goods/g7225947-02.jpeg"}],"goodsId":0,"stores":20,"unit":"册","logistics":"0","content":"另一本DevOps的经典书。","status":2,"statusStr":"在售","pic":"http://localhost:9090/pic/goods/g7225947.jpeg","minPrice":"99.0","originalPrice":"129.0","afterSale":"1"}},{"id":"g7225948","name":"DevOps实战指南","catalogueId":0,"recommendStatus":"1","pic":"http://localhost:9090/pic/goods/g7225948.jpeg","minPrice":"55.0","originalPrice":"85.0","goodsDetail":{"id":"g7225948","name":"DevOps实战指南","pics":[{"id":"g7225948-01","pic":"http://localhost:9090/pic/goods/g7225948-01.jpeg"},{"id":"g7225948-02","pic":"http://localhost:9090/pic/goods/g7225948-02.jpeg"}],"goodsId":0,"stores":2,"unit":"册","logistics":"0","content":"DevOps 黄皮书。","status":2,"statusStr":"在售","pic":"http://localhost:9090/pic/goods/g7225948.jpeg","minPrice":"55.0","originalPrice":"85.0","afterSale":"1"}},{"id":"g7225949","name":"谷歌软件工程","catalogueId":0,"recommendStatus":"1","pic":"http://localhost:9090/pic/goods/g7225949.jpeg","minPrice":"77.0","originalPrice":"107.0","goodsDetail":{"id":"g7225949","name":"谷歌软件工程","pics":[{"id":"g7225949-01","pic":"http://localhost:9090/pic/goods/g7225949-01.jpeg"},{"id":"g7225949-02","pic":"http://localhost:9090/pic/goods/g7225949-02.jpeg"}],"goodsId":0,"stores":5,"unit":"册","logistics":"0","content":"解密硅谷头部互联网企业 如何打造软件工程文化。","status":2,"statusStr":"在售","pic":"http://localhost:9090/pic/goods/g7225949.jpeg","minPrice":"77.0","originalPrice":"107.0","afterSale":"1"}}],"totalRow":4},"msg":"OK"}`

	body := utils.HttpPost(st.router, data, "/v1/goods/list")
	st.Equal(exp, body, "should same.")

	data = url.Values{}
	data.Set("page", "")
	data.Add("pageSize", "anythingIsOKByNow")
	data.Add("categoryId", "1")
	body = utils.HttpPost(st.router, data, "/v1/goods/list")
	st.NotEqual(exp, body, "should not same.")

}
func (st *GoodsHandlerSuite) setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.GET("/shop/goods/detail", GetGoodsDetail)
	v1.POST("/goods/list", FetchGoodsList)
	return router
}
