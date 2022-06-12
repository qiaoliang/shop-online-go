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
		return getSkuRepoDB(configs.Cfg.DBConnection())
	} else {
		return &SkuRepoDB{}
	}
}

func getSkuRepoDB(db *gorm.DB) SkuRepoIf {
	return SkuRepoDB{db}
}

type SkuRepoIf interface {
	FindAll() []SKU
	First(skuid string) *SKU
	FindWithCarouselPics(skuid string) *SKU
	Create(sku SKU) error
	Delete(sku SKU) error
	Update(skuid string, sku SKU) error
}
type SkuRepoDB struct {
	db *gorm.DB
}

func (s SkuRepoDB) Update(skuid string, sku SKU) error {
	oSku := SKU{SkuId: skuid}
	ret := s.db.Model(&oSku).Updates(sku)
	return ret.Error
}
func (s SkuRepoDB) FindAll() []SKU {
	var skus []SKU
	ret := s.db.Preload("SkuCarouPictures").Find(&skus)
	if ret.Error != nil {
		return nil
	}
	return skus
}
func (s SkuRepoDB) Create(sku SKU) error {
	err := s.db.Create(&sku).Error
	if err != nil {
		return err
	}
	return s.db.Save(&sku).Error
}
func (s SkuRepoDB) Delete(sku SKU) error {
	r := s.db.Model(&sku).Delete(sku)
	return r.Error
}
func (s SkuRepoDB) First(skuid string) *SKU {
	var sku SKU
	result := s.db.Where("sku_id = ?", skuid).First(&sku)
	if result.Error != nil {
		return nil
	}
	return &sku
}
func (s SkuRepoDB) FindWithCarouselPics(skuid string) *SKU {
	sku := SKU{SkuId: skuid}
	err := s.db.Model(&sku).Association("SkuCarouPictures").Error

	if err != nil {
		fmt.Println(err)
		return nil
	}
	ret := s.db.Preload("SkuCarouPictures").First(&sku)
	log.Printf("sku is %v\n", sku)
	if ret == nil {
		log.Println("Can not find Pics for " + sku.SkuId)
		return nil
	}
	if ret.Error != nil {
		log.Println(ret.Error)
		return nil
	}
	return &sku
}
