package cart

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type CartJson struct {
	Data *CartInfo `json:"data"`
	utils.JsonResult
}

type ShoppingCartHandlerSuite struct {
	suite.Suite
	router *gin.Engine
}

func TestShoppingCartHandlerSuite(t *testing.T) {
	suite.Run(t, new(ShoppingCartHandlerSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input

// This will run before each test in the suite
func (st *ShoppingCartHandlerSuite) SetupTest() {
	cartRepo = nil
	cartRepo = GetCartsInstance()
}
func (st *ShoppingCartHandlerSuite) SetupSuite() {
	st.router = setupTestRouter()
	configs.GetConfigInstance(utils.GetConfigFileForTest())
}

func (st *ShoppingCartHandlerSuite) Test_add_one_item_to_shoppingcart_for_a_token() {

	data := url.Values{}
	data.Set("token", "iamTestToken7896554")
	data.Add("goodsId", "1")
	data.Add("gid", "gid")
	data.Add("number", "5")

	body := utils.HttpPost(st.router, data, "/v1/shopping-cart/add")

	exp := `{"code":0,"data":{"token":"iamTestToken7896554","cartInfo":"iamInfos","number":5,"items":[{"key":"1","pic":"http://localhost:9090/pic/goods/g7225946-01.jpeg","status":0,"name":"CD1.0","sku":["sku1","sku3"],"price":66,"number":5,"selected":"1","optionValueName":"valueName"}],"goods":[{"goodsId":"1","number":5}]},"msg":"OK"}`
	st.Equal(exp, string(body), "should same.")
}

func (st *ShoppingCartHandlerSuite) Test_update_volume_of_item_in_shoppingcart_for_a_token() {
	//构建参数
	data := url.Values{}
	data.Set("token", "iamTestToken7896554")
	data.Add("key", "1")
	data.Add("number", "5")
	data.Add("gid", "gid")

	exp := `{"code":0,"data":{"token":"iamTestToken7896554","cartInfo":"iamInfos","number":5,"items":[{"key":"gid","pic":"http://localhost:9090/pic/goods/g7225946-01.jpeg","status":0,"name":"CD1.0","sku":["sku1","sku3"],"price":66,"number":5,"selected":"1","optionValueName":"valueName"}],"goods":[{"goodsId":"gid","number":5}]},"msg":"OK"}`
	body := string(utils.HttpPost(st.router, data, "/v1/shopping-cart/modifyNumber"))
	st.Equal(exp, string(body))
}

func (st *ShoppingCartHandlerSuite) Test_get_cart_for_unexisted_token() {

	exp := `{"code":0,"data":"","msg":"OK"}`
	//构建参数
	data := map[string]string{
		"token": "UnexistedToken",
	}
	body := utils.HttpGet("/v1/shopping-cart/info", data, st.router)

	st.Equal(exp, string(body))
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.GET("/shopping-cart/info", GetShopingCart)
	v1.POST("/shopping-cart/add", PutIntoCart)
	v1.POST("/shopping-cart/modifyNumber", UpdateShoppingCart)
	return router
}
