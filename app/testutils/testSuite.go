package testutils

import (
	"bookstore/app/configs"
	"testing"

	"github.com/stretchr/testify/suite"
)

// We'll be able to store suite-wide
// variables and add methods to this
// test suite struct
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
		configs.Cfg.MysqlDBConn()
	}
}
func (suite *SupperSuite) TeardownSuite() {
	configs.GetConfigInstance(GetConfigFileForTest())
	if configs.Cfg.Persistence {
		configs.Cfg.Downgrade()
		configs.Cfg.DBDisconnect()
	}
}

// This will run before each test in the suite
func (suite *SupperSuite) SetupTest() {}
