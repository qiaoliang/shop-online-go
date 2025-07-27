package testutils

import (
	"os"
	"path/filepath"
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
	// 在初始化配置和数据库连接之前，先删除所有的 test.db 文件
	// 这样升级脚本就会重新创建它
	suite.cleanupTestDBFiles()

	configs.GetConfigInstance(GetConfigFileForTest())
	configs.Cfg.DBConnection()
}

func (suite *SupperSuite) TeardownSuite() {
	// 测试套件结束后清理所有 test.db 文件
	suite.cleanupTestDBFiles()
}

func (suite *SupperSuite) cleanupTestDBFiles() {
	// 直接删除当前目录下的 test.db 文件
	testDBPath := "./test.db"
	if _, err := os.Stat(testDBPath); err == nil {
		os.Remove(testDBPath)
	}

	// 递归查找并删除所有 test.db 文件
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == "test.db" {
			os.Remove(path)
		}
		return nil
	})
	if err != nil {
		// 忽略错误，因为清理失败不应该影响测试结果
	}
}

func (suite *SupperSuite) SetupTest() {}
