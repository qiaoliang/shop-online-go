package goods

import (
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GoodsRepositoryTestSuite struct {
	suite.Suite
}

func (st *GoodsRepositoryTestSuite) TestExample() {
	st.Equal(true, true)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(GoodsRepositoryTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (st *GoodsRepositoryTestSuite) BeforeTest(suiteName, testName string) {

}

// This will run after test finishes
// and receives the suite and test names as input
func (st *GoodsRepositoryTestSuite) AfterTest(suiteName, testName string) {}

// This will run before before the tests in the suite are run
func (st *GoodsRepositoryTestSuite) SetupSuite() {
}

// This will run before each test in the suite
func (st *GoodsRepositoryTestSuite) SetupTest() {
	NewGoodsRepo()
}

func (st *GoodsRepositoryTestSuite) Test_should_initial_empty_goodRepo() {
	st.Equal(0, len(goodsRepo.GetGoodsList()))
}

func (st *GoodsRepositoryTestSuite) Test_create_init_goods() {
	goodsRepo.creatData()
	st.Equal(8, len(goodsRepo.GetGoodsList()))
}
