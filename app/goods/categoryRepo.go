package goods

import (
	"sync"

	"gorm.io/gorm"
)

var lockCR = &sync.Mutex{}
var cateRepo CategoryRepoInterface

// CategoryRepoInterface 定义分类仓库接口
type CategoryRepoInterface interface {
	LoadCategory() []Category
	GetList() []Category
}

// CategoryRepo 内存实现
type CategoryRepo struct {
	cates []Category
}

// CategoryRepoDB 数据库实现
type CategoryRepoDB struct {
	db *gorm.DB
}

// GetCategoryRepo 获取分类仓库实例
func GetCategoryRepo() CategoryRepoInterface {
	if cateRepo == nil {
		lockCR.Lock()
		defer lockCR.Unlock()
		if cateRepo == nil {
			cateRepo = NewCategoryRepo()
		}
	}
	return cateRepo
}

// NewCategoryRepo 创建新的分类仓库实例（内存实现）
func NewCategoryRepo() CategoryRepoInterface {
	return &CategoryRepo{}
}

// NewCategoryRepoDB 创建新的分类仓库实例（数据库实现）
func NewCategoryRepoDB(db *gorm.DB) CategoryRepoInterface {
	return &CategoryRepoDB{db: db}
}

// LoadCategory 加载分类（内存实现）
func (cr *CategoryRepo) LoadCategory() []Category {
	cate1 := &Category{0, "DevOps"}
	cate2 := &Category{1, "大数据"}
	cr.cates = append(cr.cates, *cate1)
	cr.cates = append(cr.cates, *cate2)

	return cr.cates
}

// GetList 获取分类列表（内存实现）
func (cr *CategoryRepo) GetList() []Category {
	return cr.cates
}

// LoadCategory 加载分类（数据库实现）
func (cr *CategoryRepoDB) LoadCategory() []Category {
	var categories []Category
	cr.db.Find(&categories)
	return categories
}

// GetList 获取分类列表（数据库实现）
func (cr *CategoryRepoDB) GetList() []Category {
	var categories []Category
	cr.db.Find(&categories)
	return categories
}
