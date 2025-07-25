package goods

import (
	"log"
	"sync"

	"gorm.io/gorm"
)

type GoodsRepoIf interface {
}

var lockSku = &sync.Mutex{}
var skuRepo SkuRepoIf = NewSkuRepo()

func GetSkuRepo() SkuRepoIf {
	lockSku.Lock()
	defer lockSku.Unlock()
	if skuRepo == nil {
		skuRepo = NewSkuRepo()
	}
	return skuRepo
}

func NewSkuRepo() SkuRepoIf {
	return &SkuRepoDB{}
}

func NewSkuRepoDB(db *gorm.DB) *SkuRepoDB {
	return &SkuRepoDB{db: db}
}

type SkuRepoIf interface {
	FindAll() []SKU
	First(skuid string) *SKU
	FindWithCarouselPics(skuid string) *SKU
	Create(sku SKU) error
	Delete(sku SKU) error
	Update(skuid string, sku *SKU) error
}
type SkuRepoDB struct {
	db *gorm.DB
}

func (s SkuRepoDB) Update(skuid string, sku *SKU) error {
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
	var sku SKU
	ret := s.db.Preload("SkuCarouPictures").Where("sku_id = ?", skuid).First(&sku)
	if ret.Error != nil {
		log.Println(ret.Error)
		return nil
	}
	log.Printf("sku is %v\n", sku)
	return &sku
}
