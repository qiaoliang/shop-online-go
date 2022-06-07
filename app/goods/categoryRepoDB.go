package goods

import (
	"bookstore/app/configs"
)

type CategoryRepoIf interface {
	LoadCategory() []Category
	GetList() []Category
}
type CategoryRepoDB struct {
	cates []Category
	db    *configs.DBConn
}

func GetCategoryRepoDB(db *configs.DBConn) CategoryRepoIf {
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
