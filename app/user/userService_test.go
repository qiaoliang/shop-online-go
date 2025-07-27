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
	us   *UserService
	db   *gorm.DB
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

	// 确保admin用户存在于数据库中
	adminUser := s.repo.RetriveUserByMobile(AdminMobile)
	if adminUser == nil {
		// 如果admin用户不存在，手动创建
		s.repo.CreateAdmin(AdminMobile, AdminPwd)
	}
}

func (s *UserServiceTestSuite) TeardownSuite() {
	s.SupperSuite.TeardownSuite()
	s.us = nil
}

func (s *UserServiceTestSuite) SetupTest() {
	// 清理UserService的cache，确保每个测试都有干净的状态
	if s.us != nil {
		s.us.cache = make(map[string]string)
	}
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

func (s *UserServiceTestSuite) Test_login_with_invalid_credentials() {
	user, err := s.us.login("deviceid", "deviceName", "invalid_mobile", "invalid_pwd")
	s.Nil(user, "Should return nil user for invalid credentials")
	s.NotNil(err, "Should return error for invalid credentials")
	s.Contains(err.Error(), "can not find user", "Error message should indicate user not found")
}

func (s *UserServiceTestSuite) Test_login_with_empty_mobile() {
	user, err := s.us.login("deviceid", "deviceName", "", "any_pwd")
	s.Nil(user, "Should return nil user for empty mobile")
	s.NotNil(err, "Should return error for empty mobile")
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

	// 验证返回的用户nickname是admin（数据库初始化脚本设置的nickname）
	s.Equal("admin", user.Nickname, "返回用户的Nickname应为admin")
}

func (s *UserServiceTestSuite) Test_FindUserByToken_with_empty_token() {
	user := s.us.FindUserByToken("")
	s.Nil(user, "Should return nil for empty token")
}

func (s *UserServiceTestSuite) Test_FindUserByToken_with_nonexistent_user() {
	user := s.us.FindUserByToken("nonexistent_mobile")
	s.Nil(user, "Should return nil for nonexistent user")
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

func (s *UserServiceTestSuite) Test_RegisterNewUser_with_duplicate_mobile() {
	// 先注册一个用户
	mobile := "duplicate_test_mobile" + utils.RandomImpl{}.GenStr()
	pwd := "test_pwd"
	nickname := "测试用户"
	autoLogin := "1"

	user1, err1 := s.us.RegisterNewUser(mobile, pwd, nickname, autoLogin)
	s.Nil(err1, "First registration should succeed")
	s.NotNil(user1, "First user should be created")

	// 尝试用相同的手机号注册
	user2, err2 := s.us.RegisterNewUser(mobile, "different_pwd", "different_nick", "0")
	s.Nil(user2, "Second registration should fail")
	s.NotNil(err2, "Should return error for duplicate mobile")
	s.Contains(err2.Error(), "该用户已注册", "Error message should indicate user already registered")
}

func (s *UserServiceTestSuite) Test_RegisterNewUser_with_empty_mobile() {
	user, err := s.us.RegisterNewUser("", "pwd", "nick", "1")
	s.Nil(user, "Should return nil for empty mobile")
	s.NotNil(err, "Should return error for empty mobile")
	s.Contains(err.Error(), "手机号不能为空", "Error message should indicate mobile cannot be empty")
}

func (s *UserServiceTestSuite) Test_RegisterNewUser_with_empty_pwd() {
	mobile := "empty_pwd_test" + utils.RandomImpl{}.GenStr()
	user, err := s.us.RegisterNewUser(mobile, "", "nick", "1")
	s.NotNil(user, "Database implementation allows empty password")
	s.Nil(err, "Should not return error for empty password")
	s.Equal("", user.Pwd, "Password field should be empty")
}

func (s *UserServiceTestSuite) Test_logout() {
	// 先登录一个用户
	mobile := "logout_test_mobile" + utils.RandomImpl{}.GenStr()
	s.us.RegisterNewUser(mobile, "pwd", "nick", "1")
	s.True(s.us.isOnline(mobile), "User should be online after registration")

	// 登出用户
	s.us.logout(mobile)
	s.False(s.us.isOnline(mobile), "User should be offline after logout")
}

func (s *UserServiceTestSuite) Test_logout_nonexistent_token() {
	// 登出不存在的token应该不会报错
	s.us.logout("nonexistent_token")
	// 验证没有副作用
	s.False(s.us.isOnline("nonexistent_token"))
}

func (s *UserServiceTestSuite) Test_isOnline() {
	// 测试未登录用户
	s.False(s.us.isOnline("nonexistent_token"), "Nonexistent token should be offline")

	// 测试已登录用户
	mobile := "online_test_mobile" + utils.RandomImpl{}.GenStr()
	s.us.RegisterNewUser(mobile, "pwd", "nick", "1")
	s.True(s.us.isOnline(mobile), "Registered user should be online")
}

func (s *UserServiceTestSuite) Test_findUser() {
	// 先注册一个用户
	mobile := "find_test_mobile" + utils.RandomImpl{}.GenStr()
	pwd := "test_pwd"
	nickname := "测试用户"
	autoLogin := "1"

	s.us.RegisterNewUser(mobile, pwd, nickname, autoLogin)

	// 测试正确的用户名密码
	user := s.us.findUser(mobile, pwd)
	s.NotNil(user, "Should find user with correct credentials")
	s.Equal(mobile, user.Mobile, "Should return correct user")

	// 测试错误的密码
	user = s.us.findUser(mobile, "wrong_pwd")
	s.Nil(user, "Should not find user with wrong password")

	// 测试不存在的用户
	user = s.us.findUser("nonexistent_mobile", "any_pwd")
	s.Nil(user, "Should not find nonexistent user")
}

func (s *UserServiceTestSuite) Test_GetDeliveryAddressesFor() {
	// 这个方法目前返回nil，测试其行为
	addresses := s.us.GetDeliveryAddressesFor("any_token")
	s.Nil(addresses, "Should return nil for unimplemented method")
}

func (s *UserServiceTestSuite) Test_GetDefaultDeliveryAddress() {
	// 这个方法目前返回nil，测试其行为
	addresses := s.us.GetDefaultDeliveryAddress("any_token")
	s.Nil(addresses, "Should return nil for unimplemented method")
}

// 测试UpdateUserByToken功能
func (s *UserServiceTestSuite) Test_UpdateUserByToken() {
	// 注册一个测试用户，使用随机手机号避免冲突
	mobile := "13900001234" + utils.RandomImpl{}.GenStr()
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

func (s *UserServiceTestSuite) Test_UpdateUserByToken_with_invalid_token() {
	// 测试无效token
	updateData := User{
		Nickname: "新昵称",
		City:     "新城市",
	}

	updatedUser := s.us.UpdateUserByToken("invalid_token", updateData)
	s.Nil(updatedUser, "Should return nil for invalid token")
}

func (s *UserServiceTestSuite) Test_UpdateUserByToken_with_empty_token() {
	// 测试空token
	updateData := User{
		Nickname: "新昵称",
		City:     "新城市",
	}

	updatedUser := s.us.UpdateUserByToken("", updateData)
	s.Nil(updatedUser, "Should return nil for empty token")
}

func (s *UserServiceTestSuite) Test_UpdateUserByToken_partial_update() {
	// 注册一个测试用户
	mobile := "partial_update_test" + utils.RandomImpl{}.GenStr()
	pwd := "password123"
	nickname := "原始昵称"
	autoLogin := "1"

	user, err := s.us.RegisterNewUser(mobile, pwd, nickname, autoLogin)
	s.Nil(err, "注册用户不应该返回错误")
	s.NotNil(user, "用户创建失败")

	token := mobile

	// 只更新昵称
	updateData := User{
		Nickname: "只更新昵称",
	}

	updatedUser := s.us.UpdateUserByToken(token, updateData)

	// 验证只有昵称被更新
	s.NotNil(updatedUser, "更新用户信息失败")
	s.Equal("只更新昵称", updatedUser.Nickname, "昵称应该被更新")
	s.Equal("未知", updatedUser.Province, "省份应该保持原值")
	s.Equal("未知", updatedUser.City, "城市应该保持原值")
}

func (s *UserServiceTestSuite) loginAsAdmin() {
	s.us.login("diviceid", "deviceName", AdminMobile, AdminPwd)
}

