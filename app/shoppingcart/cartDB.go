package cart

import (
	"bookstore/app/configs"
	"bookstore/app/goods"
)

type UserCart struct {
	Token string       `json:"token"`
	Items []CartItemVM `json:"items"`
	Pairs []ItemPairVM `json:"goods"`
}

type UserCartItem struct {
	Token           string
	SkuId           string
	Pic             string
	Status          uint `` // === 1 【失效】
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

func (ci *UserCart) RedDot() uint {
	if ci.Items != nil {
		return uint(len(ci.Items))
	}
	return 0
}
func (ci *UserCart) NewUserCartItem(token string, sku goods.SKU, quantity uint) UserCartItem {

	skuStr := "sku1, sku3"
	item := UserCartItem{token, sku.SkuId, configs.Cfg.GoodsPicPrefix() + sku.PicStr, 0, sku.Name, skuStr, sku.MinPrice, quantity, "1", "optionValueName"}
	return item
}

func (ci *UserCart) findItemByGid(gid string) *CartItemVM {
	for i := range ci.Items {
		it := &ci.Items[i]
		if it.Gid == gid {
			return it
		}
	}
	return nil
}
func (ci *UserCart) findPairByGid(gid string) *ItemPairVM {
	for i := range ci.Pairs {
		it := &ci.Pairs[i]
		if it.GoodsId == gid {
			return it
		}
	}
	return nil
}

func (ci *UserCart) AddMore(prod *goods.GoodsDetail, quantity uint) {
	item := ci.findItemByGid(prod.Gid)
	if item != nil {
		updatedQuantity := item.Quantity + quantity
		item.Quantity = updatedQuantity
		ip := ci.findPairByGid(prod.Gid)
		ip.Volume = item.Quantity
		return
	}
	item = ci.createCartItem(prod, quantity)
	ip := ItemPairVM{prod.Gid, quantity}
	ci.Items = append(ci.Items, *item)
	ci.Pairs = append(ci.Pairs, ip)
}

func (ci *UserCart) Modify(prod *goods.GoodsDetail, quantity uint) bool {
	item := ci.findItemByGid(prod.Gid)
	if item != nil {
		item.Quantity = quantity
		ip := ci.findPairByGid(prod.Gid)
		ip.Volume = item.Quantity
		return true
	}
	return false
}

func (c *UserCart) getVolumeById(gid string) uint {
	if len(c.Pairs) == 0 {
		return uint(0)
	}

	for _, item := range c.Pairs {
		v := item
		if v.GoodsId == gid {
			return item.Volume
		}
	}
	return uint(0)
}

func (ci *UserCart) createCartItem(prod *goods.GoodsDetail, quantity uint) *CartItemVM {
	sku := []string{"sku1", "sku3"}
	selected := "1"
	optionValue := "optionValueName"

	item := CartItemVM{prod.Gid,
		configs.Cfg.GoodsPicPrefix() + prod.Gid + ".jpeg",
		0,
		prod.Name,
		sku,
		prod.MinPrice,
		quantity,
		selected,
		optionValue,
	}
	return &item
}
