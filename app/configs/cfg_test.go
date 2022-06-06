package configs

import (
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
func (s *ConfigTestSuite) Test_BannerPath() {
	s.EqualValues("http://localhost:9090/pic/banners/", s.cfg.BannerPicPrefix())
}

func (s *ConfigTestSuite) Test_AvatarPath() {
	s.EqualValues("http://localhost:9090/pic/avatar/", s.cfg.AvatarPicPrefix())
}
func (s *ConfigTestSuite) Test_GoodsPath() {
	s.EqualValues("http://localhost:9090/pic/goods/", s.cfg.GoodsPicPrefix())
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
