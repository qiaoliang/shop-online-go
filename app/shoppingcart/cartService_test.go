package cart

import (
	"bookstore/app/configs"
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
	s.cs = NewCartService(configs.Cfg.Persistence)
}
func (s *CartServiceTestSuite) TeardownSuite() {
	s.SupperSuite.TeardownSuite()
	s.cs = nil
}
func (s *CartServiceTestSuite) SetupTest() {}

func (s *CartServiceTestSuite) Test_CreateCartInfoFor() {

}