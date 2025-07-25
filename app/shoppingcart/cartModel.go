package cart

import (
	"bookstore/app/configs"
	"bookstore/app/goods"
)

type UserCart struct {
	Token string
	Items []UserCartItem
}

type UserCartItem struct {
	ID              int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Token           string `gorm:"column:token"`
	SkuId           string `gorm:"column:sku_id"`
	Pic             string `gorm:"column:pic"`
	Status          uint   `gorm:"column:status"`
	Name            string `gorm:"column:name"`
	SkuStrs         string `gorm:"column:sku_strs"`
	Price           string `gorm:"column:price"`
	Quantity        uint   `gorm:"column:quantity"`
	Selected        string `gorm:"column:selected"`
	OptionValueName string `gorm:"column:option_value_name"`
}

func (ci UserCartItem) FullPicStr() string {
	return configs.Cfg.GoodsPicPrefix() + ci.Pic
}

func (ci *UserCart) NewUserCartItem(token string, sku goods.SKU, quantity uint) UserCartItem {

	skuStr := "sku1, sku3"
	item := UserCartItem{0, token, sku.SkuId, configs.Cfg.GoodsPicPrefix() + sku.PicStr, 0, sku.Name, skuStr, sku.MinPrice, quantity, "1", "optionValueName"}
	return item
}
func (ci *UserCart) findUserCartItem(skuid string) (*UserCartItem, int) {
	items := ci.Items
	for i := 0; i < len(items); i++ {
		if items[i].SkuId == skuid {
			return &items[i], i
		}
	}
	return nil, -1
}
