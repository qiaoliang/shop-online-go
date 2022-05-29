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

func (st *UserHandlerSuite) Should_login_with_admin() {

	data := url.Values{}
	data.Set("deviceId", "deviceId-7654321")
	data.Add("deviceName", "deviceName-7654321")
	data.Add("mobile", "13900007997")
	data.Add("pwd", "1234")

	body := utils.HttpRequest(st.router, data, "GET", "/v1/user/m/login")

	exp := "密码错误"
	st.Contains(string(body), exp)
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.GET("/user/m/login", Login)

	return router
}
