package testutils

import (
	"os"
	"testing"

	"bookstore/app/configs"

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
	configs.Cfg.DBConnection()
}
func (suite *SupperSuite) TeardownSuite() {
	// 测试套件结束后清理test.db
	os.Remove("./test.db")
}

func (suite *SupperSuite) SetupTest() {}
