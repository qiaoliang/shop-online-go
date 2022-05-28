package testdemo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// We'll be able to store suite-wide
// variables and add methods to this
// test suite struct
type ExampleTestSuite struct {
	suite.Suite
}

// This is an example test that will always succeed
func (suite *ExampleTestSuite) TestExample() {
	suite.Equal(true, true)
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (suite *ExampleTestSuite) BeforeTest(suiteName, testName string) {}

// This will run after test finishes
// and receives the suite and test names as input
func (suite *ExampleTestSuite) AfterTest(suiteName, testName string) {}

// This will run before before the tests in the suite are run
func (suite *ExampleTestSuite) SetupSuite() {}

// This will run before each test in the suite
func (suite *ExampleTestSuite) SetupTest() {}
