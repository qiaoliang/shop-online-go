package ad

import (
	"gorm.io/gorm"
)

type BannerRepoDB struct {
	db *gorm.DB
}

type Banner struct {
	BusinessId uint   `gorm:"column:businessId" json:"businessId"`
	DateAdd    string `gorm:"column:dateAdd" json:"dateAdd"`
	Id         uint   `gorm:"column:id;primaryKey" json:"id"`
	LinkUrl    string `gorm:"column:linkUrl" json:"linkUrl"`
	Paixu      uint   `gorm:"column:paixu" json:"paixu"`
	PicUrl     string `gorm:"column:picUrl" json:"picUrl"`
	Remark     string `gorm:"column:remark" json:"remark"`
	Status     uint   `gorm:"column:status" json:"status"`
	StatusStr  string `gorm:"column:statusStr" json:"statusStr"`
	Title      string `gorm:"column:title" json:"title"`
	Type       string `gorm:"column:type" json:"type"`
	UserId     uint   `gorm:"column:userId" json:"userId"`
}

func NewBannerRepoDB(db *gorm.DB) *BannerRepoDB {
	return &BannerRepoDB{db: db}
}

func (r *BannerRepoDB) FindAllBanners() []Banner {
	var banners []Banner
	r.db.Find(&banners)
	return banners
}