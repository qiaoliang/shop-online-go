package cart

import (
	"bookstore/app/configs"
	"bookstore/app/testutils"
	"bookstore/app/utils"
	"fmt"
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
	uci := NewUCIBuilder().token("test_save_UCI" + utils.RandomImpl{}.GenStr()).build()
	ret := s.repo.SaveUserCartItem(uci)
	s.Nil(ret)
	//clean up
	ret = s.repo.DeleteUserCartItem(uci)
	s.Nil(ret)
}

func (s *CartRepoDBTestSuite) Test_Update() {
	uci := NewUCIBuilder().token("test_Update_UCI" + utils.RandomImpl{}.GenStr()).build()
	ret := s.repo.SaveUserCartItem(uci)
	s.Nil(ret)
	uci.Name = "updated_name"

	ret = s.repo.UpdateUserCartItem(uci)
	s.Nil(ret)
	//clean up
	ret = s.repo.DeleteUserCartItem(uci)
	s.Nil(ret)

}
func (s *CartRepoDBTestSuite) Test_Get() {
	exp := NewUCIBuilder().token("test_get_UCI" + utils.RandomImpl{}.GenStr()).build()
	ret := s.repo.SaveUserCartItem(exp)
	s.Nil(ret)

	found := s.repo.GetUserCartItem(exp)
	s.NotNil(found)
	s.EqualValues(exp, *found)
	//clean up
	fmt.Printf("found token: %v\n", found.Token)
	ret = s.repo.DeleteUserCartItem(*found)
	s.Nil(ret)

}
