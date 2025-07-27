package user

import (
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
	userRepo := newUserRepo()
	_ = userRepo // 保证无未使用变量
}

func (suite *UserRepoTestSuite) Test_create_user() {
	userRepo := newUserRepo()
	suite.Equal(0, userRepo.TotalUsers())
	userRepo.CreateUser("mobile1", "pwd1", "nickname2", "1", genUId)
	suite.Equal(1, userRepo.TotalUsers())
	userRepo.CreateUser("mobile2", "pwd2", "nickname2", "1", genUId)
	suite.Equal(2, userRepo.TotalUsers())
}

func (suite *UserRepoTestSuite) Test_find_user_by_mobile_and_pwd() {
	userRepo := newUserRepo()
	userRepo.CreateUser("mobile", "pwd", "nickname", "1", genUId)
	result := userRepo.findUser("mobile", "pwd")
	suite.NotNil(result)
}
func (suite *UserRepoTestSuite) Test_retriveUserByMobile() {
	userRepo := newUserRepo()
	userRepo.CreateUser("mobile", "pwd", "nickname", "1", genUId)
	result := userRepo.RetriveUserByMobile("mobile")
	suite.NotNil(result)
	result = userRepo.RetriveUserByMobile("noexistedUser")
	suite.Nil(result)
}
