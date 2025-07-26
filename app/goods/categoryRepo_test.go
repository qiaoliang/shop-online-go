package goods

import (
	"testing"

	"bookstore/app/configs"

	"github.com/stretchr/testify/suite"
)

type CategoryRepoTestSuite struct {
	suite.Suite
	repo CategoryRepoInterface
}

func TestCategoryRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryRepoTestSuite))
}

// BeforeTest 在每个测试前运行
func (st *CategoryRepoTestSuite) BeforeTest(suiteName, testName string) {
}

// AfterTest 在每个测试后运行
func (st *CategoryRepoTestSuite) AfterTest(suiteName, testName string) {
}

// SetupSuite 在测试套件开始前设置
func (st *CategoryRepoTestSuite) SetupSuite() {
	// 使用测试配置文件
	configs.GetConfigInstance("../../config-test.yaml")
	// 获取数据库连接
	db := configs.Cfg.DBConnection()
	// 创建仓库实例
	st.repo = NewCategoryRepoDB(db)
	st.NotNil(st.repo)
}

// TeardownSuite 在测试套件结束后清理
func (st *CategoryRepoTestSuite) TeardownSuite() {
	st.repo = nil
}

// SetupTest 在每个测试前设置
func (st *CategoryRepoTestSuite) SetupTest() {
	cateRepo = nil
	cateRepo = st.repo
}

// Test_goods_load_from_db 测试从数据库加载分类
func (st *CategoryRepoTestSuite) Test_goods_load_from_db() {
	// 执行测试
	categories := st.repo.LoadCategory()

	// 验证结果
	st.Equal(2, len(categories))
	st.Equal("DevOps", categories[0].Name)
	st.Equal("大数据", categories[1].Name)
}

// Test_goods_get_list 测试获取分类列表
func (st *CategoryRepoTestSuite) Test_get_category_list() {
	// 执行测试
	categories := st.repo.GetList()

	// 验证结果
	st.Equal(2, len(categories))
	st.Equal(uint(0), categories[0].Id)
	st.Equal("DevOps", categories[0].Name)
	st.Equal(uint(1), categories[1].Id)
	st.Equal("大数据", categories[1].Name)
}
