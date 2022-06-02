package cart

import (
	"bookstore/app/configs"
	"bookstore/app/goods"
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
	goods.GetGoodsRepo().LoadGoods()
	data := url.Values{}
	data.Set("token", "13900007997")
	data.Add("goodsId", "g7225946")
	data.Add("number", "5")

	body := utils.HttpPost(st.router, data, "/v1/shopping-cart/add")

	exp := `{"code":0,"data":{"token":"13900007997","number":1,"items":[{"key":"g7225946","pic":"http://localhost:9090/pic/goods/g7225946.jpeg","status":0,"name":"持续交付1.0","sku":["sku1","sku3"],"price":"66.0","number":5,"selected":"1","optionValueName":"optionValueName"}],"goods":[{"goodsId":"g7225946","number":5}]},"msg":"OK"}`
	st.Equal(exp, string(body), "should same.")
}
func (st *ShoppingCartHandlerSuite) Test_add_item_in_cart_for_another_token() {
	st.Test_add_one_item_to_shoppingcart_for_a_token()
	gid := "g7225947"
	token := "13900007996"
	quantity := "10"
	data := url.Values{}
	data.Set("token", token)
	data.Add("goodsId", gid)
	data.Add("number", quantity)

	body := string(utils.HttpPost(st.router, data, "/v1/shopping-cart/add"))

	exp := `{"code":0,"data":{"token":"13900007996","number":1,"items":[{"key":"g7225947","pic":"http://localhost:9090/pic/goods/g7225947.jpeg","status":0,"name":"持续交付2.0","sku":["sku1","sku3"],"price":"99.0","number":10,"selected":"1","optionValueName":"optionValueName"}],"goods":[{"goodsId":"g7225947","number":10}]},"msg":"OK"}`
	st.Equal(exp, string(body))

}
func (st *ShoppingCartHandlerSuite) Test_add_more_items_in_shoppingcart_for_same_token() {
	st.Test_add_one_item_to_shoppingcart_for_a_token()
	gid := "g7225946"
	token := "13900007997"
	moreQuantity := "10"
	data := url.Values{}
	data.Set("token", token)
	data.Add("goodsId", gid)
	data.Add("number", moreQuantity)

	body := string(utils.HttpPost(st.router, data, "/v1/shopping-cart/add"))

	exp := `{"code":0,"data":{"token":"13900007997","number":1,"items":[{"key":"g7225946","pic":"http://localhost:9090/pic/goods/g7225946.jpeg","status":0,"name":"持续交付1.0","sku":["sku1","sku3"],"price":"66.0","number":15,"selected":"1","optionValueName":"optionValueName"}],"goods":[{"goodsId":"g7225946","number":15}]},"msg":"OK"}`
	st.Equal(exp, string(body))
}
func (st *ShoppingCartHandlerSuite) Test_update_volume_of_item_in_shoppingcart_for_a_token() {

	initquantity := uint(10)
	gid := "g7225946" //构建参数
	token := "13900007997"
	cartRepo.AddOrderIntoCart(token, gid, initquantity)

	data := url.Values{}
	data.Set("token", token)
	data.Add("gid", gid)
	data.Add("number", "10")

	exp := `{"code":0,"data":{"token":"13900007997","number":1,"items":[{"key":"g7225946","pic":"http://localhost:9090/pic/goods/g7225946.jpeg","status":0,"name":"持续交付1.0","sku":["sku1","sku3"],"price":"66.0","number":20,"selected":"1","optionValueName":"optionValueName"}],"goods":[{"goodsId":"g7225946","number":20}]},"msg":"OK"}`
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
