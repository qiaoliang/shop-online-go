package user

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

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
	user := userService.login("diviceid", "deviceName", "13900007997", "1234")
	suite.Equal("13900007997", user.Mobile, "Should found Admin directly.")
	suite.Equal("1234", user.Password, "Should get Default pwd 1234 for Admin .")
	suite.Equal("13900007997", userService.userOnline["token"], "should find admin online")
}
