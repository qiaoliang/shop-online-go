package goods

import (
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CategoryRepoTestSuite struct {
	suite.Suite
}

func TestCategoryRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryRepoTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (st *CategoryRepoTestSuite) BeforeTest(suiteName, testName string) {

}

func (st *CategoryRepoTestSuite) AfterTest(suiteName, testName string) {}

func (st *CategoryRepoTestSuite) SetupSuite() {
}

func (st *CategoryRepoTestSuite) SetupTest() {
	cateRepo = nil
	cateRepo = NewCategorRepo()
}

func (st *CategoryRepoTestSuite) Test_should_initial_empty_goodRepo() {
	st.Equal(0, len(cateRepo.GetList()))
}

func (st *CategoryRepoTestSuite) Test_goods_load_from_() {
	cateRepo.loadCategory()
	st.Equal(2, len(cateRepo.GetList()))
}
