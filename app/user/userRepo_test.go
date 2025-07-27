package user

import (
	"testing"

	"bookstore/app/configs"
	"bookstore/app/testutils"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserRepoTestSuite struct {
	testutils.SupperSuite
	db *gorm.DB
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestUserReppTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (ur *UserRepoTestSuite) BeforeTest(suiteName, testName string) {}

// This will run after test finishes
// and receives the suite and test names as input
func (ur *UserRepoTestSuite) AfterTest(suiteName, testName string) {}

// This will run before before the tests in the suite are run
func (ur *UserRepoTestSuite) SetupSuite() {
	ur.SupperSuite.SetupSuite()
	ur.db = configs.Cfg.DBConnection()
}

// This will run before each test in the suite
func (ur *UserRepoTestSuite) SetupTest() {
	// 清理测试数据
	ur.db.Where("mobile LIKE ?", "%test%").Delete(&User{})
}

// 数据库实现的测试
func (suite *UserRepoTestSuite) Test_DB_CreateUser() {
	userRepo := NewUserRepoDB(suite.db)

	// 清理测试数据
	suite.db.Where("mobile = ?", "db_test_mobile").Delete(&User{})

	user, err := userRepo.CreateUser("db_test_mobile", "pwd", "nickname", "1", genUId)
	suite.NotNil(user)
	suite.Nil(err)
	suite.Equal("db_test_mobile", user.Mobile)
	suite.Equal("pwd", user.Pwd)
	suite.Equal("nickname", user.Nickname)
	suite.Equal(uint(1), user.AutoLogin)

	// 清理测试数据
	suite.db.Where("mobile = ?", "db_test_mobile").Delete(&User{})
}

func (suite *UserRepoTestSuite) Test_DB_CreateUser_with_duplicate_mobile() {
	userRepo := NewUserRepoDB(suite.db)

	// 清理测试数据
	mobile := "db_duplicate_test_mobile"
	suite.db.Where("mobile = ?", mobile).Delete(&User{})

	// 第一次创建用户
	user1, err1 := userRepo.CreateUser(mobile, "pwd1", "nickname1", "1", genUId)
	suite.NotNil(user1)
	suite.Nil(err1)

	// 尝试用相同手机号创建用户
	user2, err2 := userRepo.CreateUser(mobile, "pwd2", "nickname2", "0", genUId)
	suite.Nil(user2)
	suite.NotNil(err2)
	suite.Contains(err2.Error(), "hello,error")

	// 清理测试数据
	suite.db.Where("mobile = ?", mobile).Delete(&User{})
}

func (suite *UserRepoTestSuite) Test_DB_RetriveUserByMobile() {
	userRepo := NewUserRepoDB(suite.db)

	// 清理测试数据
	mobile := "db_retrieve_test_mobile"
	suite.db.Where("mobile = ?", mobile).Delete(&User{})

	// 创建用户
	user, _ := userRepo.CreateUser(mobile, "pwd", "nickname", "1", genUId)
	suite.NotNil(user)

	// 检索用户
	result := userRepo.RetriveUserByMobile(mobile)
	suite.NotNil(result)
	suite.Equal(mobile, result.Mobile)
	suite.Equal("pwd", result.Pwd)
	suite.Equal("nickname", result.Nickname)

	// 测试不存在的用户
	result = userRepo.RetriveUserByMobile("nonexistent_mobile")
	suite.Nil(result)

	// 清理测试数据
	suite.db.Where("mobile = ?", mobile).Delete(&User{})
}

func (suite *UserRepoTestSuite) Test_DB_RetriveUserByID() {
	userRepo := NewUserRepoDB(suite.db)

	// 清理测试数据
	mobile := "db_id_test_mobile"
	suite.db.Where("mobile = ?", mobile).Delete(&User{})

	// 创建用户
	user, _ := userRepo.CreateUser(mobile, "pwd", "nickname", "1", genUId)
	suite.NotNil(user)

	// 通过ID检索用户
	result := userRepo.RetriveUserByID(user.Id)
	suite.NotNil(result)
	suite.Equal(user.Id, result.Id)
	suite.Equal(mobile, result.Mobile)

	// 测试不存在的ID
	result = userRepo.RetriveUserByID("nonexistent_id")
	suite.Nil(result)

	// 清理测试数据
	suite.db.Where("mobile = ?", mobile).Delete(&User{})
}

func (suite *UserRepoTestSuite) Test_DB_findUser() {
	userRepo := NewUserRepoDB(suite.db)

	// 清理测试数据
	mobile := "db_find_test_mobile"
	suite.db.Where("mobile = ?", mobile).Delete(&User{})

	// 创建用户
	user, _ := userRepo.CreateUser(mobile, "pwd", "nickname", "1", genUId)
	suite.NotNil(user)

	// 正确的用户名密码
	result := userRepo.findUser(mobile, "pwd")
	suite.NotNil(result)
	suite.Equal(mobile, result.Mobile)

	// 错误的密码
	result = userRepo.findUser(mobile, "wrong_pwd")
	suite.Nil(result)

	// 不存在的用户
	result = userRepo.findUser("nonexistent_mobile", "any_pwd")
	suite.Nil(result)

	// 清理测试数据
	suite.db.Where("mobile = ?", mobile).Delete(&User{})
}

func (suite *UserRepoTestSuite) Test_DB_TotalUsers() {
	userRepo := NewUserRepoDB(suite.db)

	// 获取初始用户数量
	initialCount := userRepo.TotalUsers()

	// 创建测试用户
	mobile := "db_count_test_mobile"
	suite.db.Where("mobile = ?", mobile).Delete(&User{})

	user, _ := userRepo.CreateUser(mobile, "pwd", "nickname", "1", genUId)
	suite.NotNil(user)

	// 验证用户数量增加
	newCount := userRepo.TotalUsers()
	suite.Equal(initialCount+1, newCount)

	// 清理测试数据
	suite.db.Where("mobile = ?", mobile).Delete(&User{})
}

func (suite *UserRepoTestSuite) Test_DB_DeleteByMobile() {
	userRepo := NewUserRepoDB(suite.db)

	// 清理测试数据
	mobile := "db_delete_test_mobile"
	suite.db.Where("mobile = ?", mobile).Delete(&User{})

	// 创建用户
	user, _ := userRepo.CreateUser(mobile, "pwd", "nickname", "1", genUId)
	suite.NotNil(user)

	// 验证用户存在
	result := userRepo.RetriveUserByMobile(mobile)
	suite.NotNil(result)

	// 删除用户
	userRepo.DeleteByMobile(mobile)

	// 验证用户已被删除
	result = userRepo.RetriveUserByMobile(mobile)
	suite.Nil(result)
}

func (suite *UserRepoTestSuite) Test_DB_CreateAdmin() {
	userRepo := NewUserRepoDB(suite.db)

	// 清理测试数据
	mobile := "db_admin_test_mobile"
	suite.db.Where("mobile = ?", mobile).Delete(&User{})

	userRepo.CreateAdmin(mobile, "admin_pwd")

	user := userRepo.RetriveUserByMobile(mobile)
	suite.NotNil(user)
	suite.Equal(mobile, user.Mobile)
	suite.Equal("admin_pwd", user.Pwd)
	suite.Equal("超级塞亚人", user.Nickname)
	suite.Equal(uint(1), user.AutoLogin)

	// 清理测试数据
	suite.db.Where("mobile = ?", mobile).Delete(&User{})
}

func (suite *UserRepoTestSuite) Test_DB_updateUser() {
	userRepo := NewUserRepoDB(suite.db)

	// 清理测试数据
	mobile := "db_update_test_mobile"
	suite.db.Where("mobile = ?", mobile).Delete(&User{})

	// 创建用户
	user, _ := userRepo.CreateUser(mobile, "pwd", "nickname", "1", genUId)
	suite.NotNil(user)

	// 更新用户信息
	user.Nickname = "updated_nickname"
	user.Province = "广东省"
	user.City = "深圳市"
	userRepo.updateUser(user)

	// 验证更新
	updatedUser := userRepo.RetriveUserByMobile(mobile)
	suite.NotNil(updatedUser)
	suite.Equal("updated_nickname", updatedUser.Nickname)
	suite.Equal("广东省", updatedUser.Province)
	suite.Equal("深圳市", updatedUser.City)

	// 清理测试数据
	suite.db.Where("mobile = ?", mobile).Delete(&User{})
}

func (suite *UserRepoTestSuite) Test_DB_updateUser_with_nil_user() {
	userRepo := NewUserRepoDB(suite.db)
	// 不应该panic
	userRepo.updateUser(nil)
}

func (suite *UserRepoTestSuite) Test_DB_updateUser_with_empty_mobile() {
	userRepo := NewUserRepoDB(suite.db)
	user := &User{
		Mobile: "",
		Nickname: "test",
	}
	// 不应该panic
	userRepo.updateUser(user)
}

func (suite *UserRepoTestSuite) Test_genUId() {
	// 测试ID生成器
	id1 := genUId()
	id2 := genUId()

	suite.NotEmpty(id1)
	suite.NotEmpty(id2)
	suite.NotEqual(id1, id2, "Generated IDs should be unique")
	suite.Contains(id1, "userId", "Generated ID should contain 'userId' prefix")
}
