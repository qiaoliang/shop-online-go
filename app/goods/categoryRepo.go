package goods

import (
	"sync"
)

var lockCR = &sync.Mutex{}
var cateRepo *CategoryRepo

type CategoryRepo struct {
	cates []Category
}

func GetCategoryRepo() *CategoryRepo {
	return NewCategoryRepo()
}
func NewCategoryRepo() *CategoryRepo {
	return &CategoryRepo{}
}
func (cr *CategoryRepo) LoadCategory() []Category {
	cate1 := &Category{0, "DevOps"}
	cate2 := &Category{1, "大数据"}
	cr.cates = append(cr.cates, *cate1)
	cr.cates = append(cr.cates, *cate2)

	return cr.cates
}
func (cr *CategoryRepo) GetList() []Category {
	return cr.cates
}
