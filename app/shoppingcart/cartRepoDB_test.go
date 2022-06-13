package cart

import (
	"bookstore/app/configs"
	"bookstore/app/testutils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CartRepoDBTestSuite struct {
	testutils.SupperSuite
	repo CartRepoIf
}

func TestCartRepoDBTestSuite(t *testing.T) {
	suite.Run(t, new(CartRepoDBTestSuite))

}

func (s *CartRepoDBTestSuite) BeforeTest(suiteName, testName string) {}

func (s *CartRepoDBTestSuite) AfterTest(suiteName, testName string) {}

func (s *CartRepoDBTestSuite) SetupSuite() {
	cartRepo = nil
	s.SupperSuite.SetupSuite()
	s.repo = NewCartsRepo(configs.Cfg.Persistence)
}
func (s *CartRepoDBTestSuite) TeardownSuite() {
	s.SupperSuite.TeardownSuite()
	s.repo = nil
}

func (s *CartRepoDBTestSuite) SetupTest() {
}

func (s *CartRepoDBTestSuite) Test_Save() {
	uci := NewUCIBuilder().token("test_save_token").build()
	ret := s.repo.SaveUserCartItem(uci)
	s.Nil(ret)
	//clean up
	ret = s.repo.DeleteUserCartItem(uci)
	s.Nil(ret)

}
