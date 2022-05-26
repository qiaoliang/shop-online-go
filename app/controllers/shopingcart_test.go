package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestOne(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")
	v1.POST("/shopping-cart/add", PutIntoCart)

	//构建参数
	data := url.Values{}
	data.Set("token", "anyToken")
	data.Add("goodsid", "1")
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

	exp := "{\"code\":0,\"data\":null,\"msg\":\"OK\"}"

	if exp != string(body) {
		t.Fatalf("resp.StatusCode=%v, Content-Type=%v, expected=%v actual=%v\n", resp.StatusCode, resp.Header.Get("Content-Type"), exp, string(body))
	}
}
