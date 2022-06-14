package cart

import (
	"bookstore/app/configs"
	"bookstore/app/goods"
	"bookstore/app/testutils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CartServiceTestSuite struct {
	testutils.SupperSuite
	cs *CartService
}

func TestSkuRepoDBTestSuite(t *testing.T) {
	suite.Run(t, new(CartServiceTestSuite))
}

func (s *CartServiceTestSuite) BeforeTest(suiteName, testName string) {}

func (s *CartServiceTestSuite) AfterTest(suiteName, testName string) {}

func (s *CartServiceTestSuite) SetupSuite() {
	s.SupperSuite.SetupSuite()
	s.cs = newCartService(configs.Cfg.Persistence)
}
func (s *CartServiceTestSuite) TeardownSuite() {
	s.SupperSuite.TeardownSuite()
	s.cs = nil
	cartRepo = nil
	cartService = nil
}
func (s *CartServiceTestSuite) SetupTest() {}

func (s *CartServiceTestSuite) Test_GetPersistance() {
	cs := newCartService(false)
	_, isok := cs.cr.(*CartRepo)
	s.True(isok)
	cs = newCartService(true)
	_, ok := cs.cr.(*CartRepoDB)
	s.True(ok)
}
func (s *CartServiceTestSuite) Test_CreateCartInfoFor() {
	token := "create_cartInfo_token"
	sku_id := "g7225946"
	prod := goods.GetGoodsRepo().GetItemDetail(sku_id)
	quantity := uint(10)
	ci := s.cs.CreateCartInfoFor(token, prod, quantity)
	s.NotNil(ci)
}
