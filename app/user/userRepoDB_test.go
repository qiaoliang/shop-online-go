package user

import (
	"bookstore/app/configs"
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepoDBTestSuite struct {
	suite.Suite
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestUserRepoDBTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoDBTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (ur *UserRepoDBTestSuite) BeforeTest(suiteName, testName string) {}

// This will run after test finishes
// and receives the suite and test names as input
func (ur *UserRepoDBTestSuite) AfterTest(suiteName, testName string) {}

// This will run before before the tests in the suite are run
func (ur *UserRepoDBTestSuite) SetupSuite() {}

// This will run before each test in the suite
func (ur *UserRepoDBTestSuite) SetupTest() {
	userRepo = nil
	userRepo = GetUserRepoInstance()
}

func (suite *UserRepoDBTestSuite) Should_create_user() {
	suite.Equal(1, userRepo.TotalUsers())
	user, _ := userRepo.CreateUser("mobile1", "pwd1", "nickname2")

	suite.Equal(2, userRepo.TotalUsers())
	suite.Equal("mobile1", user.Mobile)
	userRepo.CreateUser("mobile2", "pwd2", "nickname2")
	suite.Equal(3, userRepo.TotalUsers())
}

func (suite *UserRepoDBTestSuite) Should_find_user_by_mobile_and_pwd() {
	userRepo.CreateUser("mobile", "pwd", "nickname")
	result := userRepo.findUser("mobile", "pwd")
	suite.NotEmpty(result)
	pattern := configs.Cfg.AvatarPicPrefix() + "[a-l]\\.jpeg$"
	reg, _ := regexp.Compile(pattern)
	reg.MatchString(result.AvatarUrl)
	suite.Equal("mobile", result.Mobile)
}
func (suite *UserRepoDBTestSuite) Should_retriveUserByMobile() {
	userRepo.CreateUser("mobile", "pwd", "nickname")
	result := userRepo.retriveUserByMobile("mobile")
	suite.NotEmpty(result)
	result = userRepo.retriveUserByMobile("noexistedUser")
	suite.True(result == nil)
}
