package goods

import (
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GoodsRepositoryTestSuite struct {
	suite.Suite
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(GoodsRepositoryTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (st *GoodsRepositoryTestSuite) BeforeTest(suiteName, testName string) {

}

func (st *GoodsRepositoryTestSuite) AfterTest(suiteName, testName string) {}

func (st *GoodsRepositoryTestSuite) SetupSuite() {
}

func (st *GoodsRepositoryTestSuite) SetupTest() {
	goodsRepo = nil
	GetGoodsRepo()
}

func (st *GoodsRepositoryTestSuite) Test_should_initial_empty_goodRepo() {
	st.Equal(0, len(goodsRepo.GetGoodsList()))
}

func (st *GoodsRepositoryTestSuite) Test_goods_load_from_() {
	goodsRepo.loadGoods()
	st.Equal(8, len(goodsRepo.GetGoodsList()))
}
