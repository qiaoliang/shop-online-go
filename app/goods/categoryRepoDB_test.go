package goods

import (
	"bookstore/app/configs"
	"testing"

	"github.com/stretchr/testify/suite"
)

// We'll be able to store suite-wide
// variables and add methods to this
// test suite struct
type CategoryRepoDBTestSuite struct {
	suite.Suite
	db CategoryRepoIf
}

// This is an example test that will always succeed
func (s *CategoryRepoDBTestSuite) Test_get_categories_from_db() {
	categories := s.db.GetList()
	s.Equal(2, len(categories))
	s.Equal(uint(0), categories[0].Id)
	s.Equal("DevOps", categories[0].Name)
	s.Equal(uint(1), categories[1].Id)
	s.Equal("大数据", categories[1].Name)
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryRepoDBTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (s *CategoryRepoDBTestSuite) BeforeTest(suiteName, testName string) {

}

// This will run after test finishes
// and receives the suite and test names as input
func (s *CategoryRepoDBTestSuite) AfterTest(suiteName, testName string) {
}

// This will run before before the tests in the suite are run
func (s *CategoryRepoDBTestSuite) SetupSuite() {

	s.db = GetCategoryRepoDB(configs.Cfg.DBConnection())
	s.NotNil(s.db)
}

func (suite *CategoryRepoDBTestSuite) TeardownSuite() {
	configs.Cfg.Downgrade()

}
