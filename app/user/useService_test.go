package user

import (
	"bookstore/app/testutils"
	"bookstore/app/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	testutils.SupperSuite
	us *UserService
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
	s.us = newUserService(true)
}
func (s *UserServiceTestSuite) TeardownSuite() {
	s.SupperSuite.TeardownSuite()
	s.us = nil
}

func (s *UserServiceTestSuite) SetupTest() {
}

var AdminMobile = "13900007997"
var AdminPwd = "1234"

func (s *UserServiceTestSuite) Test_assertUserRepoPersistance() {
}

func (s *UserServiceTestSuite) Test_admin_login() {

	user, _ := s.us.login("diviceid", "deviceName", AdminMobile, "1234")
	s.NotNil(user)
	s.Equal(AdminMobile, user.Mobile, "Should found Admin directly.")
	s.Equal(AdminPwd, user.Password, "Should get Default pwd "+AdminPwd+" for Admin .")
	s.Equal(AdminMobile, s.us.userOnline[AdminMobile], "should find admin online")
}
func (s *UserServiceTestSuite) Test_findUserByMobile() {
	s.loginAsAdmin()
	s.True(s.us.isOnline(AdminMobile))
	user := s.us.FindUserByToken(AdminMobile)
	s.Equal(AdminMobile, user.Mobile, "Should found Admin directly.")
	s.Equal(AdminPwd, user.Password, "Should get Default pwd "+AdminPwd+" for Admin .")
	offlineuser := s.us.FindUserByToken("offlineUser")
	s.True(offlineuser == nil)
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

func (s *UserServiceTestSuite) loginAsAdmin() {
	s.us.login("diviceid", "deviceName", AdminMobile, AdminPwd)
}
