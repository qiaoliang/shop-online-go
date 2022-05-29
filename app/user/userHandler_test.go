package user

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type UserHandlerSuite struct {
	suite.Suite
	router *gin.Engine
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerSuite))
}

func (st *UserHandlerSuite) SetupSuite() {
	st.router = setupTestRouter()
	configs.NewConfig(utils.GetConfigFileForTest())
}

func (st *UserHandlerSuite) Test_login_with_admin() {

	data := url.Values{}
	data.Set("deviceId", "deviceId-7654321")
	data.Add("deviceName", "deviceName-7654321")
	data.Add("mobile", "13900007997")
	data.Add("pwd", "1234")

	body := utils.HttpRequest(st.router, data, "GET", "/v1/user/m/login")

	exp := `{"code":0,"data":{"token":"iamTestToken7896554","cartInfo":"iamInfos","number":5,"items":[{"key":1,"pic":"http://localhost:9090/pic/goods/g7227946-01.jpeg","status":0,"name":"CD1.0","sku":["sku1","sku3"],"price":66,"number":5,"selected":"1","optionValueName":"valueName"}],"goods":[{"goodsId":1,"number":5}]},"msg":"OK"}`
	st.Equal(exp, string(body), "should same.")
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.GET("/user/m/login", Login)

	return router
}
