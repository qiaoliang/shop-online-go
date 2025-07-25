package cart

import (
	"gorm.io/gorm"
)

type CartRepoDB struct {
	db *gorm.DB
}

func NewCartRepoDB(db *gorm.DB) *CartRepoDB {
	return &CartRepoDB{db: db}
}

func (r *CartRepoDB) FindUserCartItemsBy(token string) []UserCartItem {
	var items []UserCartItem
	r.db.Where("Token = ?", token).Find(&items)
	return items
}

func (r *CartRepoDB) SaveUserCartItem(item UserCartItem) error {
	existing := r.GetUserCartItem(item)
	if existing != nil {
		existing.Quantity += item.Quantity
		return r.db.Model(&UserCartItem{}).Where("Token = ? AND sku_id = ?", item.Token, item.SkuId).Updates(existing).Error
	}
	return r.db.Create(&item).Error
}

func (r *CartRepoDB) DeleteUserCartItemsBy(token string) error {
	return r.db.Where("Token = ?", token).Delete(&UserCartItem{}).Error
}

func (r *CartRepoDB) UpdateUserCartItem(item *UserCartItem) error {
	return r.db.Model(&UserCartItem{}).Where("Token = ? AND sku_id = ?", item.Token, item.SkuId).Updates(item).Error
}

func (r *CartRepoDB) DeleteUserCartItem(item UserCartItem) error {
	return r.db.Where("Token = ? AND sku_id = ?", item.Token, item.SkuId).Delete(&UserCartItem{}).Error
}

func (r *CartRepoDB) GetUserCartItem(item UserCartItem) *UserCartItem {
	var result UserCartItem
	if err := r.db.Where("Token = ? AND sku_id = ?", item.Token, item.SkuId).First(&result).Error; err != nil {
		return nil
	}
	return &result
}