package user

import (
	"testing"

	"bookstore/app/configs"
	"bookstore/app/testutils"
	"bookstore/app/utils"

	"gorm.io/gorm"

	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	testutils.SupperSuite
	us *UserService
	db *gorm.DB
	repo UserRepo
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (s *UserServiceTestSuite) BeforeTest(suiteName, testName string) {}

// This will run after test finishes
// and receives the suite and test names as input
func (s *UserServiceTestSuite) AfterTest(suiteName, testName string) {}

// This will run before before the tests in the suite are run
func (s *UserServiceTestSuite) SetupSuite() {
	s.SupperSuite.SetupSuite()
	s.db = configs.Cfg.DBConnection()
	s.repo = NewUserRepoDB(s.db)
	s.us = NewUserServiceWithRepo(s.repo)
	// 已有 admin 用户由 migration 脚本插入，无需重复注册
}
func (s *UserServiceTestSuite) TeardownSuite() {
	s.SupperSuite.TeardownSuite()
	s.us = nil
}

func (s *UserServiceTestSuite) SetupTest() {
}

var AdminMobile = "13900007997"
var AdminPwd = "1234"

func (s *UserServiceTestSuite) Test_admin_login() {

	user, _ := s.us.login("diviceid", "deviceName", AdminMobile, "1234")
	s.NotNil(user)
	s.Equal(AdminMobile, user.Mobile, "Should found Admin directly.")
	s.Equal(AdminPwd, user.Pwd, "Should get Default pwd "+AdminPwd+" for Admin .")
	s.Equal(AdminMobile, s.us.cache[AdminMobile], "should find admin online")
}
func (s *UserServiceTestSuite) Test_findUserByMobile() {
	s.loginAsAdmin()
	s.True(s.us.isOnline(AdminMobile))
	user := s.us.FindUserByToken(AdminMobile)
	s.Equal(AdminMobile, user.Mobile, "Should found Admin directly.")
	s.Equal(AdminPwd, user.Pwd, "Should get Default pwd "+AdminPwd+" for Admin .")
	offlineuser := s.us.FindUserByToken("offlineUser")
	s.True(offlineuser == nil)
}

func (s *UserServiceTestSuite) Test_FindUserByToken_AdminUser() {
	// 确保admin用户在缓存中
	s.loginAsAdmin()

	// 当token为13900007997时，应返回admin用户
	user := s.us.FindUserByToken("13900007997")

	// 验证返回的用户不为空
	s.NotNil(user, "应该找到admin用户")

	// 验证返回的用户mobile是13900007997
	s.Equal(AdminMobile, user.Mobile, "返回用户的Mobile应为13900007997")

	// 验证返回的用户nickname是admin
	s.Equal("admin", user.Nickname, "返回用户的Nickname应为admin")
}

func (s *UserServiceTestSuite) Test_is_Online_after_register() {
	expMobile := "newMobile" + utils.RandomImpl{}.GenStr()
	n, err := s.us.RegisterNewUser(expMobile, "pwd", "nickString", "0")
	s.Nil(err)
	s.Equal(expMobile, n.Mobile)
	s.True(s.us.isOnline(expMobile))
	r := s.us.findUser(expMobile, "pwd")
	s.Equal(expMobile, r.Mobile)
	s.Contains(r.AvatarUrl, ".jpeg")
	s.True(s.us.isOnline(expMobile))
}

func (s *UserServiceTestSuite) loginAsAdmin() {
	s.us.login("diviceid", "deviceName", AdminMobile, AdminPwd)
}

// 测试UpdateUserByToken功能
func (s *UserServiceTestSuite) Test_UpdateUserByToken() {
	// 注册一个测试用户
	mobile := "13900001234"
	pwd := "password123"
	nickname := "测试用户"
	autoLogin := "1"

	user, err := s.us.RegisterNewUser(mobile, pwd, nickname, autoLogin)
	s.Nil(err, "注册用户不应该返回错误")
	s.NotNil(user, "用户创建失败")

	// 使用手机号作为token（根据当前实现）
	token := mobile

	// 准备更新数据
	newNickname := "更新后的昵称"
	newAvatarUrl := "http://example.com/new-avatar.jpg"
	newProvince := "广东省"
	newCity := "深圳市"

	updateData := User{
		Nickname:  newNickname,
		AvatarUrl: newAvatarUrl,
		Province:  newProvince,
		City:      newCity,
	}

	// 执行更新
	updatedUser := s.us.UpdateUserByToken(token, updateData)

	// 验证更新结果
	s.NotNil(updatedUser, "更新用户信息失败，返回了nil")
	s.Equal(newNickname, updatedUser.Nickname, "昵称更新失败")
	s.Equal(newAvatarUrl, updatedUser.AvatarUrl, "头像URL更新失败")
	s.Equal(newProvince, updatedUser.Province, "省份更新失败")
	s.Equal(newCity, updatedUser.City, "城市更新失败")

	// 再次获取用户，确认更新已持久化
	fetchedUser := s.us.FindUserByToken(token)
	s.NotNil(fetchedUser, "无法通过token获取更新后的用户")
	s.Equal(newNickname, fetchedUser.Nickname, "持久化后昵称不匹配")
}

