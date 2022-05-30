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
	configs.GetConfigInstance(utils.GetConfigFileForTest())
}

// This will run right before the test starts
// and receives the suite and test names as input
func (ur *UserHandlerSuite) BeforeTest(suiteName, testName string) {}

// This will run after test finishes
// and receives the suite and test names as input
func (ur *UserHandlerSuite) AfterTest(suiteName, testName string) {}

// This will run before each test in the suite
func (ur *UserHandlerSuite) SetupTest() {
	userService = nil
	userRepo = nil
	userService = GetUserService()
	userRepo = GetUserRepoInstance()

}

func (st *UserHandlerSuite) Test_login_with_admin() {

	data := url.Values{}
	data.Set("deviceId", "deviceId-7654321")
	data.Add("deviceName", "deviceName-7654321")
	data.Add("mobile", "13900007997")
	data.Add("pwd", "1234")

	body := utils.HttpRequest(st.router, data, "POST", "/v1/user/m/login")

	st.Contains(string(body), "13900007997", "should return admin")
}
func (st *UserHandlerSuite) Test_Register_User() {
	data := url.Values{}
	data.Set("mobile", "newUser")
	data.Add("pwd", "secret")
	data.Add("nick", "天下无贼")
	data.Add("autoLogin", "1")
	data.Add("code", "5678")

	body := utils.HttpRequest(st.router, data, "POST", "/v1/user/m/register")
	exp := `{"code":0,"data":{"token":"newUser","base":{"id":"userIdnl8x7lipma","pwd":"天下无贼","mobile":"newUser","nick":"secret","avatarUrl":"http://localhost:9090/pic/avatar/d.jpeg","province":"未知","city":"未知","autoLogin":0,"userInfo":"FakeUserInfo","userLevel":{"id":1,"name":"Green"}},"userLevel":{"id":1,"name":"Green"}},"msg":"OK"}`
	st.Equal(exp, string(body), "should rigister successfully")

}
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.POST("/user/m/login", Login)
	v1.POST("/user/m/register", Register)
	return router
}
