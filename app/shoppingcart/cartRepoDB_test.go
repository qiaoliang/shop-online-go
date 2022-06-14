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
	s.repo = newCartsRepo(configs.Cfg.Persistence)
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

	ret = s.repo.UpdateUserCartItem(&uci)
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
func (s *CartRepoDBTestSuite) Test_FindMore() {
	//arrange
	token := "test_Find_More_UCI" + utils.RandomImpl{}.GenStr()
	exp1 := NewUCIBuilder().token(token).skuId("new-Sku-id2").build()
	ret := s.repo.SaveUserCartItem(exp1)
	s.Nil(ret)
	exp2 := NewUCIBuilder().token(token).skuId("new-Sku-id1").build()
	ret = s.repo.SaveUserCartItem(exp2)
	s.Nil(ret)
	//act
	found := s.repo.FindUserCartItemsBy(token)
	//assert
	s.NotNil(found)
	s.Equal(2, len(found))
	s.EqualValues(exp1, found[0])
	s.EqualValues(exp2, found[1])
	//clean up
	ret = s.repo.DeleteUserCartItemsBy(token)
	s.Nil(ret)
}
func (s *CartRepoDBTestSuite) Test_Get_CartRepo_DB() {
	cartRepo = nil
	cr := newCartsRepo(true)
	_, isDB := cr.(*CartRepoDB)
	s.True(isDB)
	cartRepo = nil
	cr = newCartsRepo(false)
	_, isMem := cr.(*CartRepo)
	s.True(isMem)
}
