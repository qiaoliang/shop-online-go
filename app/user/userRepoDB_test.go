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
	userRepo = nil
	ur.SupperSuite.SetupSuite()
	ur.repo = GetUserRepoDB(configs.Cfg.DBConnection())
}
func (ur *UserRepoDBTestSuite) TeardownSuite() {
	ur.SupperSuite.TeardownSuite()
	ur.repo = nil
}

func (ur *UserRepoDBTestSuite) SetupTest() {
}

func (ur *UserRepoDBTestSuite) Test_total_users() {
	ur.Equal(1, ur.repo.TotalUsers())
}
func (ur *UserRepoDBTestSuite) Should_Create_users() {
	exp := ur.repo.TotalUsers()
	ur.repo.CreateUser("mymoble", "mypwd", "nickname")
	ur.Equal(exp+1, ur.repo.TotalUsers())
}

func (ur *UserRepoDBTestSuite) should_find_user_by_mobile_and_pwd() {
	result := ur.repo.findUser("13900007997", "1234")
	ur.NotEmpty(result)
	pattern := configs.Cfg.AvatarPicPrefix() + "[a-l]\\.jpeg$"
	reg, _ := regexp.Compile(pattern)
	reg.MatchString(result.AvatarUrl)
	ur.Equal("13900007997", result.Mobile)
}
func (ur *UserRepoDBTestSuite) Test_retriveUserByMobile() {
	result := ur.repo.retriveUserByMobile("13900007997")
	ur.NotEmpty(result)
	//result = userRepo.retriveUserByMobile("noexistedUser")
	//ur.True(result == nil)
}
