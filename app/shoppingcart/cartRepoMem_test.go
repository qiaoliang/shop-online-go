package cart

import (
	"bookstore/app/goods"
	"bookstore/app/testutils"
	"bookstore/app/utils"
	"fmt"
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CartRepositoryTestSuite struct {
	testutils.SupperSuite
	repo  CartRepoIf
	gRepo *goods.GoodsRepoMem
}

func (st *CartRepositoryTestSuite) TestExample() {
	st.Equal(true, true)
}

func TestCartRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CartRepositoryTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (st *CartRepositoryTestSuite) BeforeTest(suiteName, testName string) {

}

// This will run after test finishes
// and receives the suite and test names as input
func (st *CartRepositoryTestSuite) AfterTest(suiteName, testName string) {
}

// This will run before before the tests in the suite are run
func (st *CartRepositoryTestSuite) SetupSuite() {
	cartRepo = nil
	st.SupperSuite.SetupSuite()
	st.repo = newCartsRepo(false)
	st.gRepo = goods.GetGoodsRepo()
	st.gRepo.LoadGoods()
}

// This will run before each test in the suite
func (st *CartRepositoryTestSuite) SetupTest() {
	cartRepo = nil
	cartRepo = newCartsRepo(false)
}

func (s *CartRepositoryTestSuite) Test_Save() {
	uci := NewUCIBuilder().token("test_save_UCI" + utils.RandomImpl{}.GenStr()).build()
	ret := s.repo.SaveUserCartItem(uci)
	s.Nil(ret)
	//clean up
	ret = s.repo.DeleteUserCartItem(uci)
	s.Nil(ret)
}

func (s *CartRepositoryTestSuite) Test_Update() {
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
func (s *CartRepositoryTestSuite) Test_Get() {
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
func (s *CartRepositoryTestSuite) Test_FindMore() {
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
func (s *CartRepositoryTestSuite) Test_Get_CartRepo_DB() {
	cartRepo = nil
	cr := newCartsRepo(true)
	_, isDB := cr.(*CartRepoDB)
	s.True(isDB)
	cartRepo = nil
	cr = newCartsRepo(false)
	_, isMem := cr.(*CartRepoMem)
	s.True(isMem)
}
