package user

import (
	"math/rand"
	"regexp"
	"testing"

	"github.com/example/project/app/configs"
	"github.com/example/project/app/testutils"

	"github.com/stretchr/testify/suite"
)

type UserRepoDBTestSuite struct {
	testutils.SupperSuite
	repo UserRepoIf
}

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

func (ur *UserRepoDBTestSuite) Test_generate_String() {
	str := GenStr(4)
	ur.Equal(4, len(str))
	str2 := GenStr(5)
	ur.Equal(5, len(str2))
	result := make(map[string]string, 100)
	for i := 0; i < 100; i++ {
		str = GenStr(10)
		result[str] = str
	}
	ur.Equal(100, len(result))
}

func (ur *UserRepoDBTestSuite) Test_get_DB_REPO_Instance() {
	result := GetUserRepo()
	_, ok := result.(*UserRepoDB)
	ur.True(ok)
}

func (ur *UserRepoDBTestSuite) Test_total_users() {
	ur.True(ur.repo.TotalUsers() >= 1)
}
func (ur *UserRepoDBTestSuite) Test_Create_users() {
	ur.repo.CreateUser("mymobile", "mypwd", "nickname", "1", genUId)
	result := ur.repo.retriveUserByMobile("mymobile")
	ur.Equal("mymobile", result.Mobile)
	ur.NotContains(result.AvatarUrl, "http://localhost:9090/pic/avatar/")

	//cleanup
	ur.repo.DeleteByMobile("mymobile")

	result = userRepo.retriveUserByMobile("mymobile")
	ur.True(result == nil)
}

func (ur *UserRepoDBTestSuite) Test_find_user_by_mobile_and_pwd() {
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
	ur.Equal(GREENTYPE, result.UserLevelId)
	ur.Equal("这是UserInfo", result.UserInfo)
	ur.Equal("a.jpeg", result.AvatarUrl)
	result = userRepo.retriveUserByMobile("noexistedUser")
	ur.True(result == nil)

}

func GenStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
