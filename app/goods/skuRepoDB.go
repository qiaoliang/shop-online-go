package goods

import (
	"bookstore/app/configs"
	"fmt"
	"log"
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
	FindWithCarouselPics(skuid string) *SKU
	Create(sku SKU) error
}
type SkuRepoDB struct {
	db *gorm.DB
}

func (s SkuRepoDB) Create(sku SKU) error {
	err := s.db.Create(&sku).Error
	if err != nil {
		return err
	}
	return s.db.Save(&sku).Error
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
func (s SkuRepoDB) FindWithCarouselPics(skuid string) *SKU {
	sku := SKU{SkuId: skuid}
	err := s.db.Model(&sku).Association("SkuCarouPictures").Error

	if err != nil {
		fmt.Println(err)
		return nil
	}
	result := s.db.Preload("SkuCarouPictures").First(&sku)
	log.Printf("sku is %v\n", sku)
	if result == nil {
		log.Println("Can not find Pics for " + sku.SkuId)
		return nil
	}
	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}
	return &sku
}
