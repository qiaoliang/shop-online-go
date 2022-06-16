package user

import (
	"bookstore/app/testutils"
	"bookstore/app/utils"
	"fmt"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type UserHandlerSuite struct {
	testutils.SupperSuite
	router *gin.Engine
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerSuite))
}

func (st *UserHandlerSuite) SetupSuite() {
	st.SupperSuite.SetupSuite()
	st.router = setupTestRouter()
}

func (ur *UserHandlerSuite) BeforeTest(suiteName, testName string) {}

func (ur *UserHandlerSuite) AfterTest(suiteName, testName string) {}

func (ur *UserHandlerSuite) SetupTest() {
}

func (st *UserHandlerSuite) Test_login_with_admin() {

	data := url.Values{}
	data.Set("deviceId", "any")
	data.Add("deviceName", "any")
	data.Add("mobile", "13900007997")
	data.Add("pwd", "1234")

	body := testutils.HttpPost(st.router, data, "/v1/user/m/login")

	st.Contains(string(body), "13900007997", "should return admin")
}

func (st *UserHandlerSuite) Test_Register_User() {

	mobile := "AnyMobile" + utils.RandomImpl{}.GenStr()
	token := mobile
	pwd := "un_encryption"
	autologin := "1"
	nick := "天下无贼"
	code := "unimplemented"
	data := url.Values{}
	data.Set("mobile", mobile)
	data.Add("pwd", pwd)
	data.Add("nick", nick)
	data.Add("autoLogin", autologin)
	data.Add("code", code)

	body := testutils.HttpPost(st.router, data, "/v1/user/m/register")
	exp1 := fmt.Sprintf(`{"code":0,"data":{"token":"%v","base":{"id":`, token)
	exp2 := fmt.Sprintf(`"pwd":"%v","mobile":"%v","nick":"%v","avatarUrl":"http://localhost:9090/pic/avatar/`,
		pwd, mobile, nick)
	exp3 := fmt.Sprintf(`.jpeg","province":"未知","city":"未知","autoLogin":%v,"userInfo":"FakeUserInfo","userLevel":{"id":1,"name":"Green"}},"userLevel":{"id":1,"name":"Green"}},"msg":"OK"}`,
		autologin)

	st.Contains(string(body), exp1, "should contain token:"+token)
	st.Containsf(string(body), exp2, "should contain pwd :%v, mobile:%v, nick:  %v\n", pwd, mobile, nick)
	st.Contains(string(body), exp3, "should contain autologin:%v", autologin)
}
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.POST("/user/m/login", Login)
	v1.POST("/user/m/register", Register)
	return router
}
