package goods

import (
	"bookstore/app/configs"
)

var cateRepoDB *CategoryRepoDB

type CategoryRepoIf interface {
	loadCategory() []Category
	GetList() []Category
}
type CategoryRepoDB struct {
	cates []Category
	db    *configs.DBConn
}

func GetCategoryRepoDB(db *configs.DBConn) *CategoryRepoDB {
	if cateRepoDB == nil {
		cateRepoDB = &CategoryRepoDB{[]Category{}, db}
	}
	return cateRepoDB
}
func (cr *CategoryRepoDB) loadCategory() []Category {
	cr.db.Find(&cr.cates)
	return cr.cates
}
func (cr *CategoryRepoDB) GetList() []Category {
	if len(cr.cates) == 0 {
		cr.loadCategory()
	}
	return cr.cates
}
