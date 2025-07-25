package goods

import (
	"net/url"
	"testing"

	"bookstore/app/configs"
	"bookstore/app/testutils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type GoodItemJson struct {
	Data *GoodsItem `json:"data"`
	testutils.JsonResult
}

type GoodsHandlerSuite struct {
	suite.Suite
	router *gin.Engine
}

func TestGoodsHandlerSuite(t *testing.T) {
	suite.Run(t, new(GoodsHandlerSuite))
}

func (st *GoodsHandlerSuite) SetupSuite() {
	configs.GetConfigInstance(testutils.GetConfigFileForTest())
	db := configs.Cfg.DBConnection()
	skuRepo := NewSkuRepoDB(db)
	cateRepo := &CategoryRepo{}
	goodsService := NewGoodsService(skuRepo, cateRepo)
	goodsHandler := NewGoodsHandler(goodsService)
	st.router = st.setupTestRouter(goodsHandler)
}

func (st *GoodsHandlerSuite) Test_get_GoodsDetail() {
	//构建参数
	params := map[string]string{
		"token": "anythingIsOKByNow",
		"id":    "g7225946",
	}
	//构建返回值
	//调用请求接口
	exp := `{"code":0,"data":{"id":"g7225946","name":"持续交付1.0","pics":[{"id":"g7225946-01","pic":"http://localhost:9090/pic/goods/g7225946-01.jpeg"},{"id":"g7225946-02","pic":"http://localhost:9090/pic/goods/g7225946-02.jpeg"}],"goodsId":0,"stores":10,"unit":"本","logistics":"1","content":"这是第一本 DevOps 的书","status":0,"statusStr":"在售","pic":"http://localhost:9090/pic/goods/g7225946.jpeg","minPrice":"66.0","originalPrice":"99.0","afterSale":"1"},"msg":"OK"}`

	body := testutils.HttpGet("/v1/shop/goods/detail", params, st.router)
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
	exp := `{"code":0,"data":{"result":[{"id":"g7225946","name":"持续交付1.0","catalogueId":0,"recommendStatus":"1","pic":"http://localhost:9090/pic/goods/g7225946.jpeg","minPrice":"66.0","originalPrice":"99.0","goodsDetail":{"id":"g7225946","name":"持续交付1.0","pics":[{"id":"g7225946-01","pic":"http://localhost:9090/pic/goods/g7225946-01.jpeg"},{"id":"g7225946-02","pic":"http://localhost:9090/pic/goods/g7225946-02.jpeg"}],"goodsId":0,"stores":10,"unit":"本","logistics":"1","content":"这是第一本 DevOps 的书","status":0,"statusStr":"在售","pic":"http://localhost:9090/pic/goods/g7225946.jpeg","minPrice":"66.0","originalPrice":"99.0","afterSale":"1"}},{"id":"g7225947","name":"持续交付2.0","catalogueId":0,"recommendStatus":"1","pic":"http://localhost:9090/pic/goods/g7225947.jpeg","minPrice":"88.0","originalPrice":"120.0","goodsDetail":{"id":"g7225947","name":"持续交付2.0","pics":[{"id":"g7225947-01","pic":"http://localhost:9090/pic/goods/g7225947-01.jpeg"},{"id":"g7225947-02","pic":"http://localhost:9090/pic/goods/g7225947-02.jpeg"}],"goodsId":0,"stores":20,"unit":"本","logistics":"1","content":"这是第二本 DevOps 的书","status":0,"statusStr":"在售","pic":"http://localhost:9090/pic/goods/g7225947.jpeg","minPrice":"88.0","originalPrice":"120.0","afterSale":"1"}}],"totalRow":2},"msg":"OK"}`

	body := testutils.HttpPost(st.router, data, "/v1/goods/list")
	st.Equal(exp, body, "should same.")

	data = url.Values{}
	data.Set("page", "")
	data.Add("pageSize", "anythingIsOKByNow")
	data.Add("categoryId", "888")
	body = testutils.HttpPost(st.router, data, "/v1/goods/list")
	st.NotEqual(exp, body, "should not same.")

}
func (st *GoodsHandlerSuite) setupTestRouter(handler *GoodsHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")
	v1.GET("/shop/goods/detail", handler.GetGoodsDetail)
	v1.POST("/goods/list", handler.FetchGoodsList)
	return router
}
