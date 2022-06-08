package user

import (
	"bookstore/app/configs"
	"bookstore/app/testutils"
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepoDBTestSuite struct {
	testutils.SupperSuite
	repo UserRepoIf
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestUserRepoDBTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoDBTestSuite))

}

func (ur *UserRepoDBTestSuite) BeforeTest(suiteName, testName string) {}

func (ur *UserRepoDBTestSuite) AfterTest(suiteName, testName string) {}

func (ur *UserRepoDBTestSuite) SetupSuite() {
	ur.SupperSuite.SetupSuite()
	ur.repo = GetUserRepoDB(configs.Cfg.DBConnection())
}
func (ur *UserRepoDBTestSuite) TeardownSuite() {
	ur.SupperSuite.TeardownSuite()
	ur.repo = nil
}

// This will run before each test in the suite
func (ur *UserRepoDBTestSuite) SetupTest() {
}
func (ur *UserRepoDBTestSuite) Should_create_user() {
	ur.Equal(1, ur.repo.TotalUsers())
	user, _ := ur.repo.CreateUser("mobile1", "pwd1", "nickname2")

	ur.Equal(2, ur.repo.TotalUsers())
	ur.Equal("mobile1", user.Mobile)
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
