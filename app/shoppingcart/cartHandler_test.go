package cart

import (
	"bookstore/app/configs"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ShoppingCartHandlerSuite struct {
	suite.Suite
	router *gin.Engine
}

func TestShoppingCartHandlerSuite(t *testing.T) {
	suite.Run(t, new(ShoppingCartHandlerSuite))
}

func (st *ShoppingCartHandlerSuite) SetupSuite() {
	st.router = setupTestRouter()

	_, filename, _, _ := runtime.Caller(0)
	path, _ := filepath.Abs(filename)
	path = filepath.Dir(path)
	configs.NewConfig(path + "/../../config.yaml")
}

func (st *ShoppingCartHandlerSuite) Test_add_one_item_to_shoppingcart_for_a_token() {

	data := url.Values{}
	data.Set("token", "iamTestToken7896554")
	data.Add("goodsId", "1")
	data.Add("number", "5")

	req, _ := http.NewRequest("POST", "/v1/shopping-cart/add", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//构建返回值
	w := httptest.NewRecorder()

	//调用请求接口
	st.router.ServeHTTP(w, req)
	st.True(w.Code == http.StatusOK, "should return Http OK.")

	resp := w.Result()
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)

	//exp := `{"code":0,"data":{"token":"iamTestToken7896554","cartInfo":"iamInfos","number":5,"items":[{"key":1,"pic":"http://localhost:9090/pic/goods/g7227946-01.jpeg","status":0,"name":"CD1.0","sku":["sku1","sku3"],"price":66,"number":5,"selected":"1","optionValueName":"valueName"}],"goods":[{"goodsId":1,"number":5}]},"msg":"OK"}`
	//st.Equal(exp, string(body), "should same.")
}

func Test_add_one_item_to_shoppingcart_for_a_token(t *testing.T) {

	router := setupTestRouter()

	//构建参数
	data := url.Values{}
	data.Set("token", "iamTestToken7896554")
	data.Add("goodsId", "1")
	data.Add("number", "5")

	req, _ := http.NewRequest("POST", "/v1/shopping-cart/add", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//构建返回值
	w := httptest.NewRecorder()

	//调用请求接口
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	resp := w.Result()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	exp := `{"code":0,"data":{"token":"iamTestToken7896554","cartInfo":"iamInfos","number":5,"items":[{"key":1,"pic":"http://localhost:9090/pic/goods/g7227946-01.jpeg","status":0,"name":"CD1.0","sku":["sku1","sku3"],"price":66,"number":5,"selected":"1","optionValueName":"valueName"}],"goods":[{"goodsId":1,"number":5}]},"msg":"OK"}`

	if exp != string(body) {
		t.Fatalf("exp=%v, actual = %v", exp, string(body))
	}
}

func Test_update_volume_of_item_in_shoppingcart_for_a_token(t *testing.T) {
	router := setupTestRouter()

	//构建参数
	data := url.Values{}
	data.Set("token", "iamTestToken7896554")
	data.Add("key", "1")
	data.Add("number", "5")

	req, _ := http.NewRequest("POST", "/v1/shopping-cart/modifyNumber", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//构建返回值
	w := httptest.NewRecorder()

	//调用请求接口
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	resp := w.Result()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	exp := `{"code":0,"data":{"token":"iamTestToken7896554","cartInfo":"iamInfos","number":5,"items":[{"key":1,"pic":"http://localhost:9090/pic/goods/g7227946-01.jpeg","status":0,"name":"CD1.0","sku":["sku1","sku3"],"price":66,"number":5,"selected":"1","optionValueName":"valueName"}],"goods":[{"goodsId":1,"number":5}]},"msg":"OK"}`

	if exp != string(body) {
		t.Fatalf("exp=%v, actual = %v", exp, string(body))
	}
}

func Test_get_cart_for_unexisted_token(t *testing.T) {
	router := setupTestRouter()

	//构建参数
	data := url.Values{}
	data.Set("token", "UnexistedToken")

	req, _ := http.NewRequest("GET", "/v1/shopping-cart/info", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	//构建返回值
	w := httptest.NewRecorder()

	//调用请求接口
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	resp := w.Result()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	exp := `{"code":0,"data":"","msg":"OK"}`

	if exp != string(body) {
		t.Fatalf("exp=%v, actual = %v", exp, string(body))
	}
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
