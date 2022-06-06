package goods

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

// We'll be able to store suite-wide
// variables and add methods to this
// test suite struct
type CategoryRepoDBTestSuite struct {
	suite.Suite
	db *CategoryRepoDB
}

// This is an example test that will always succeed
func (s *CategoryRepoDBTestSuite) Test_get_categories_from_db() {
	categories := s.db.GetList()
	s.Equal(2, len(categories))
	s.Equal(uint(1), s.db.cates[0].Id)
	s.Equal("DevOps", s.db.cates[0].Name)
	s.Equal(uint(2), s.db.cates[1].Id)
	s.Equal("大数据", s.db.cates[1].Name)
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryRepoDBTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (s *CategoryRepoDBTestSuite) BeforeTest(suiteName, testName string) {
	configs.GetConfigInstance(utils.GetConfigFileForTest())
	configs.Cfg.Upgrade()
	configs.Cfg.MysqlDBConn()
	s.db = GetCategoryRepoDB(configs.Cfg.DBConnection())
	s.NotNil(s.db)
}

// This will run after test finishes
// and receives the suite and test names as input
func (s *CategoryRepoDBTestSuite) AfterTest(suiteName, testName string) {
	configs.Cfg.Downgrade()
}

// This will run before before the tests in the suite are run
func (suite *CategoryRepoDBTestSuite) SetupSuite() {
}

// This will run before each test in the suite
func (suite *CategoryRepoDBTestSuite) SetupTest() {
	configs.GetConfigInstance(utils.GetConfigFileForTest())
	configs.Cfg.Upgrade()

}
