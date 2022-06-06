package goods

import (
	"bookstore/app/configs"
	"sync"
)

var lockCRD = &sync.Mutex{}
var cateRepoDB *CategoryRepoDB

type CategoryRepoIf interface {
	loadCategory() []Category
	GetList() []Category
}
type CategoryRepoDB struct {
	cates []Category
}

func GetCategoryRepoDB() *CategoryRepoDB {
	lockCRD.Lock()
	defer lockCR.Unlock()
	if cateRepo == nil {
		cateRepoDB = &CategoryRepoDB{}
	}
	return cateRepoDB
}
func (cr *CategoryRepoDB) loadCategory() []Category {
	var categories []Category
	configs.DB.Find(&categories)
	cr.cates = categories
	return cr.cates
}
func (cr *CategoryRepoDB) GetList() []Category {
	if cr.cates == nil {
		cr.loadCategory()
	}
	return cr.cates
}
