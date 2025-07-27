package user

import (
	"fmt"
	"net/url"
	"testing"

	"bookstore/app/configs"
	"bookstore/app/testutils"
	"bookstore/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserHandlerSuite struct {
	testutils.SupperSuite
	router *gin.Engine
	db     *gorm.DB
	repo   UserRepo
	us     *UserService
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerSuite))
}



func (st *UserHandlerSuite) SetupSuite() {
	st.SupperSuite.SetupSuite()
	st.db = configs.Cfg.DBConnection()
	st.repo = NewUserRepoDB(st.db)
	st.us = NewUserServiceWithRepo(st.repo)
	userHandler := NewUserHandler(st.us)
	st.router = setupTestRouter(userHandler)
	// 调试：打印 users 表所有内容
	var users []User
	st.db.Find(&users)
	for _, u := range users {
		fmt.Printf("[DEBUG] user: id=%s, mobile=%s, pwd=%s, nick=%s\n", u.Id, u.Mobile, u.Pwd, u.Nickname)
	}
	// 已有 admin 用户由 migration 脚本插入，无需重复注册
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

func (st *UserHandlerSuite) Test_GetUserDetail() {
	// 确保已在setupTestRouter中已添加GetUserDetail路由
	// 发送请求，token为13900007997
	params := map[string]string{"token": "13900007997"}
	body := testutils.HttpGet("/v1/user/detail", params, st.router)

	// 验证响应
	st.Contains(body, `"code":0`, "应返回成功状态码")
	st.Contains(body, `"nick":"admin"`, "当token为13900007997时，用户名应为admin")
	st.Contains(body, `"mobile":"13900007997"`, "应返回正确的手机号")
}

func setupTestRouter(userHandler *UserHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")
	v1.POST("/user/m/login", userHandler.Login)
	v1.POST("/user/m/register", userHandler.Register)
	v1.GET("/user/detail", userHandler.GetUserDetail) // 添加GetUserDetail路由
	return router
}
