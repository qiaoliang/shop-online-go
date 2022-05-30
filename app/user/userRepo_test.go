package user

import (
	"bookstore/app/configs"
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
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
func (ur *UserRepoTestSuite) SetupSuite() {}

// This will run before each test in the suite
func (ur *UserRepoTestSuite) SetupTest() {
	userRepo = nil
	userRepo = GetUserRepoInstance()
}

func (suite *UserRepoTestSuite) Test_create_user() {
	suite.Equal(1, userRepo.TotalUsers())
	user := userRepo.CreateUser("mobile1", "pwd1", "nickname2")

	suite.Equal(2, userRepo.TotalUsers())
	suite.Equal("mobile1", user.Mobile)
	userRepo.CreateUser("mobile2", "pwd2", "nickname2")
	suite.Equal(3, userRepo.TotalUsers())
}

func (suite *UserRepoTestSuite) Test_find_user() {
	userRepo.CreateUser("mobile", "pwd", "nickname")
	result := userRepo.findUser("mobile", "pwd")
	suite.NotEmpty(result)
	pattern := configs.Cfg.AvatarPicPrefix() + "[a-l]\\.jpeg$"
	reg, _ := regexp.Compile(pattern)
	reg.MatchString(result.AvatarUrl)
	suite.Equal("mobile", result.Mobile)
}
