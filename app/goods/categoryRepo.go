package goods

import (
	"bookstore/app/configs"
	"sync"
)

var lockCR = &sync.Mutex{}
var cateRepo CategoryRepoIf

type CategoryRepo struct {
	cates []Category
}

func GetCategoryRepo() CategoryRepoIf {
	return NewCategoryRepo(configs.Cfg.Persistence)
}
func NewCategoryRepo(persistence bool) CategoryRepoIf {
	lockCR.Lock()
	defer lockCR.Unlock()

	if cateRepo == nil {
		if persistence {

			cateRepo = GetCategoryRepoDB(configs.Cfg.DBConnection())

		} else {

			cateRepo = &CategoryRepo{}
		}
	}
	return cateRepo
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
