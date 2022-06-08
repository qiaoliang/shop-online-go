package goods

import (
	"gorm.io/gorm"
)

type CategoryRepoIf interface {
	LoadCategory() []Category
	GetList() []Category
}
type CategoryRepoDB struct {
	cates []Category
	db    *gorm.DB
}

func GetCategoryRepoDB(db *gorm.DB) CategoryRepoIf {
	if cateRepo == nil {
		cateRepo = &CategoryRepoDB{[]Category{}, db}
	}
	return cateRepo
}
func (cr *CategoryRepoDB) LoadCategory() []Category {
	cr.db.Find(&cr.cates)
	return cr.cates
}
func (cr *CategoryRepoDB) GetList() []Category {
	if len(cr.cates) == 0 {
		cr.LoadCategory()
	}
	return cr.cates
}
