package cart

import (
	"bookstore/app/testutils"
	"bookstore/app/utils"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type CartJson struct {
	Data *CartInfoVM `json:"data"`
	testutils.JsonResult
}

type ShoppingCartHandlerSuite struct {
	testutils.SupperSuite
	router *gin.Engine
}

func TestShoppingCartHandlerSuite(t *testing.T) {
	suite.Run(t, new(ShoppingCartHandlerSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input

// This will run before each test in the suite
func (st *ShoppingCartHandlerSuite) SetupTest() {
}
func (st *ShoppingCartHandlerSuite) SetupSuite() {
	st.SupperSuite.SetupSuite()
	st.router = setupTestRouter()
}

const (
	SECOND_SKU = "g7225947"
)

func (st *ShoppingCartHandlerSuite) Test_add_one_item_to_shoppingcart_for_a_token() {
	//goods.GetGoodsRepo().LoadGoods()
	data := url.Values{}
	token := "an_token_" + utils.RandomImpl{}.GenStr()
	data.Set("token", token)
	data.Add("goodsId", EXISTED_SKU_ONE)
	data.Add("number", "5")

	body := testutils.HttpPost(st.router, data, "/v1/shopping-cart/add")

	exp := fmt.Sprintf(`{"code":0,"data":{"token":"%v","number":1,"items":[{"key":"%v","pic":"http://localhost:9090/pic/goods/%v.jpeg","status":0,"name":"持续交付1.0","sku":["sku1","sku3"],"price":"66.0","number":5,"selected":"1","optionValueName":"OptionValueName"}],"goods":[{"goodsId":"%v","number":5}]},"msg":"OK"}`,
		token, EXISTED_SKU_ONE, EXISTED_SKU_ONE, EXISTED_SKU_ONE)
	st.Equal(exp, string(body), "should same.")
}
func (st *ShoppingCartHandlerSuite) Test_add_item_in_cart_for_another_token() {
	st.Test_add_one_item_to_shoppingcart_for_a_token()
	gid := EXISTED_SKU_ONE
	token := "second_token_" + utils.RandomImpl{}.GenStr()
	quantity := "13"
	data := url.Values{}
	data.Set("token", token)
	data.Add("goodsId", gid)
	data.Add("number", quantity)

	body := string(testutils.HttpPost(st.router, data, "/v1/shopping-cart/add"))

	exp := fmt.Sprintf(`{"code":0,"data":{"token":"%v","number":1,"items":[{"key":"%v","pic":"http://localhost:9090/pic/goods/%v.jpeg","status":0,"name":"持续交付1.0","sku":["sku1","sku3"],"price":"66.0","number":%v,"selected":"1","optionValueName":"OptionValueName"}],"goods":[{"goodsId":"%v","number":%v}]},"msg":"OK"}`,
		token, gid, gid, quantity, gid, quantity)
	st.Equal(exp, string(body))

}
func (st *ShoppingCartHandlerSuite) Test_add_more_quntity_for_same_goods_in_shoppingcart_for_same_token() {
	gid := EXISTED_SKU_ONE
	token := "same_token_add_more" + utils.RandomImpl{}.GenStr()
	alreadyHave := uint(46)
	GetCartsService().PutItemsInCart(token, gid, alreadyHave)
	log.Println("prepare a existed sku in cart for the token.")
	extraQuantity := uint(10)
	data := url.Values{}
	data.Set("token", token)
	data.Add("goodsId", gid)
	data.Add("number", strconv.Itoa(int(extraQuantity)))
	totalQuantity := alreadyHave + extraQuantity
	//
	body := string(testutils.HttpPost(st.router, data, "/v1/shopping-cart/add"))

	exp := fmt.Sprintf(`{"code":0,"data":{"token":"%v","number":1,"items":[{"key":"%v","pic":"http://localhost:9090/pic/goods/%v.jpeg","status":0,"name":"持续交付1.0","sku":["sku1","sku3"],"price":"66.0","number":%v,"selected":"1","optionValueName":"OptionValueName"}],"goods":[{"goodsId":"%v","number":%v}]},"msg":"OK"}`,
		token, gid, gid, totalQuantity, gid, totalQuantity)
	st.Equal(exp, string(body))
}
func (st *ShoppingCartHandlerSuite) Test_Modify_number_of_item_in_shoppingcart_for_a_token() {
	gid := EXISTED_SKU_ONE
	token := "same_token_" + utils.RandomImpl{}.GenStr()
	initquantity := uint(10)
	GetCartsService().PutItemsInCart(token, gid, initquantity)

	newQuan := "11"
	data := url.Values{}
	data.Set("token", token)
	data.Add("key", gid)
	data.Add("number", newQuan)
	totalQuantity := uint(11)

	body := string(testutils.HttpPost(st.router, data, "/v1/shopping-cart/modifyNumber"))
	exp := fmt.Sprintf(`{"code":0,"data":{"token":"%v","number":1,"items":[{"key":"%v","pic":"http://localhost:9090/pic/goods/%v.jpeg","status":0,"name":"持续交付1.0","sku":["sku1","sku3"],"price":"66.0","number":%v,"selected":"1","optionValueName":"OptionValueName"}],"goods":[{"goodsId":"%v","number":%v}]},"msg":"OK"}`,
		token, gid, gid, totalQuantity, gid, totalQuantity)
	st.Equal(exp, string(body))
}

func (st *ShoppingCartHandlerSuite) Test_get_cart_for_unexisted_token() {

	exp := `{"code":0,"data":"","msg":"OK"}`
	//构建参数
	data := map[string]string{
		"token": "UnexistedToken",
	}
	body := testutils.HttpGet("/v1/shopping-cart/info", data, st.router)

	st.Equal(exp, string(body))
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.GET("/shopping-cart/info", GetShopingCart)
	v1.POST("/shopping-cart/add", PutIntoCart)
	v1.POST("/shopping-cart/modifyNumber", ModifyNumberOfGoodsInCart)
	return router
}
