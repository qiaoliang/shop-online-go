package user

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var AdminMobile = "13900007997"
var AdminPwd = "1234"

type UserServiceTestSuite struct {
	suite.Suite
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (ur *UserServiceTestSuite) BeforeTest(suiteName, testName string) {}

// This will run after test finishes
// and receives the suite and test names as input
func (ur *UserServiceTestSuite) AfterTest(suiteName, testName string) {}

// This will run before before the tests in the suite are run
func (ur *UserServiceTestSuite) SetupSuite() {}

// This will run before each test in the suite
func (ur *UserServiceTestSuite) SetupTest() {
	userService = nil
	userService = GetUserService()
}

func (suite *UserServiceTestSuite) Test_admin_login() {

	user := GetUserService().login("diviceid", "deviceName", AdminMobile, "1234")
	suite.Equal(AdminMobile, user.Mobile, "Should found Admin directly.")
	suite.Equal(AdminPwd, user.Password, "Should get Default pwd "+AdminPwd+" for Admin .")
	suite.Equal(AdminMobile, userService.userOnline[AdminMobile], "should find admin online")
}
func (suite *UserServiceTestSuite) Test_findUserByMobile() {
	suite.loginAsAdmin()
	suite.True(GetUserService().isOnline(AdminMobile))
	user := GetUserService().FindUserByToken(AdminMobile)
	suite.Equal(AdminMobile, user.Mobile, "Should found Admin directly.")
	suite.Equal(AdminPwd, user.Password, "Should get Default pwd "+AdminPwd+" for Admin .")
	offlineuser := GetUserService().FindUserByToken("offlineUser")
	suite.True(offlineuser == nil)
}

func (suite *UserServiceTestSuite) Test_is_Online_after_register() {
	n := GetUserService().RegisterNewUser("newMobile", "pwd", "nickString")
	suite.Equal("newMobile", n.Mobile)
	suite.True(GetUserService().isOnline("newMobile"))
	r := GetUserService().findUser("newMobile", "pwd")
	suite.Equal("newMobile", r.Mobile)
	suite.Contains(r.AvatarUrl, ".jpeg")
	suite.True(GetUserService().isOnline("newMobile"))
}

func (suite *UserServiceTestSuite) loginAsAdmin() {
	GetUserService().login("diviceid", "deviceName", AdminMobile, AdminPwd)
}
