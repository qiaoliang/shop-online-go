package testutils

import (
	"bookstore/app/configs"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SupperSuite struct {
	suite.Suite
}

func TestBookstoreTestSuite(t *testing.T) {
	suite.Run(t, new(SupperSuite))
}

func (suite *SupperSuite) BeforeTest(suiteName, testName string) {}

func (suite *SupperSuite) AfterTest(suiteName, testName string) {}

func (suite *SupperSuite) SetupSuite() {
	configs.GetConfigInstance(GetConfigFileForTest())
	if configs.Cfg.Persistence {
		configs.Cfg.Upgrade()
		configs.Cfg.DBConnection()
	}
}
func (suite *SupperSuite) TeardownSuite() {}

func (suite *SupperSuite) SetupTest() {}
