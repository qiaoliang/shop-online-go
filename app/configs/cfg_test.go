package configs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

// We'll be able to store suite-wide
// variables and add methods to this
// test suite struct
type ConfigTestSuite struct {
	suite.Suite
	cfg *Config
}

func (s *ConfigTestSuite) Test_StaicPicPath() {
	s.EqualValues("http://localhost:9090/pic", s.cfg.StaticPic)
}

func (s *ConfigTestSuite) Test_Config_Load() {
	s.NotNil(s.cfg)
	// 只测试 SQLite 相关配置
	s.NotEmpty(s.cfg.SQLiteDBFile)
}

func (s *ConfigTestSuite) Test_BannerPath() {
	s.EqualValues("http://localhost:9090/pic/banners/", s.cfg.BannerPicPrefix())
}

func (s *ConfigTestSuite) Test_AvatarPath() {
	s.EqualValues("http://localhost:9090/pic/avatar/", s.cfg.AvatarPicPrefix())
}
func (s *ConfigTestSuite) Test_GoodsPath() {
	s.EqualValues("http://localhost:9090/pic/goods/", s.cfg.GoodsPicPrefix())
}

func (s *ConfigTestSuite) Test_DBMigration() {
	s.NotPanics(func() {
		s.cfg.runMigrations()
	})
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (s *ConfigTestSuite) BeforeTest(suiteName, testName string) {}

// This will run after test finishes
// and receives the suite and test names as input
func (s *ConfigTestSuite) AfterTest(suiteName, testName string) {}

// This will run before before the tests in the suite are run
func (s *ConfigTestSuite) SetupSuite() {
	s.cfg = GetConfigInstance("../../config-test.yaml")
}

// This will run before each test in the suite
func (s *ConfigTestSuite) SetupTest() {}

// This will run after all tests in the suite are run
func (s *ConfigTestSuite) TeardownSuite() {
	// 清理 test.db 文件
	s.cleanupTestDBFiles()
}

func (s *ConfigTestSuite) cleanupTestDBFiles() {
	// 递归查找并删除所有 test.db 文件
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == "test.db" {
			os.Remove(path)
		}
		return nil
	})
}
