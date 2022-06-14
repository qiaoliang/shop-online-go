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
	Token           string
	SkuId           string `grom:"column:sku_id"`
	Pic             string
	Status          uint // === 1 【失效】
	Name            string
	SkuStrs         string `grom:"column:sku_strs"`
	Price           string ``
	Quantity        uint
	Selected        string
	OptionValueName string `grom:"column:Option_value_name"`
}

func (ci UserCartItem) FullPicStr() string {
	return configs.Cfg.GoodsPicPrefix() + ci.Pic
}

func (ci *UserCart) NewUserCartItem(token string, sku goods.SKU, quantity uint) UserCartItem {

	skuStr := "sku1, sku3"
	item := UserCartItem{token, sku.SkuId, configs.Cfg.GoodsPicPrefix() + sku.PicStr, 0, sku.Name, skuStr, sku.MinPrice, quantity, "1", "optionValueName"}
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
