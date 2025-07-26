package goods

import (
	"gorm.io/gorm"
)


func NewCategoryRepoDB(db *gorm.DB) *CategoryRepoDB {
	return &CategoryRepoDB{db: db}
}

// CategoryRepoDB 数据库实现
type CategoryRepoDB struct {
	db *gorm.DB
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
