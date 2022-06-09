package goods

import (
	"bookstore/app/configs"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

var lockSku = &sync.Mutex{}
var skuRepo SkuRepoIf

func GetSkuRepo() SkuRepoIf {
	lockSku.Lock()
	defer lockSku.Unlock()
	if skuRepo == nil {
		skuRepo = NewSkuRepo(configs.Cfg.Persistence)
	}
	return skuRepo
}

func NewSkuRepo(isPersistence bool) SkuRepoIf {
	if isPersistence {
		return getSkuRepoDB(configs.Cfg.GormDB())
	} else {
		return &SkuRepoDB{}
	}
}

func getSkuRepoDB(db *gorm.DB) SkuRepoIf {
	return SkuRepoDB{db}
}

type SkuRepoIf interface {
	Find(skuid string) *SKU
}
type SkuRepoDB struct {
	db *gorm.DB
}

func (s SkuRepoDB) Find(skuid string) *SKU {
	var sku SKU
	result := s.db.Where("sku_id = ?", skuid).First(&sku)
	if result.Error != nil {
		return nil
	}
	fmt.Printf("%v\n", sku)
	return &sku
}
