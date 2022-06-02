package cart

import (
	"bookstore/app/configs"
	"bookstore/app/goods"
)

type CartInfo struct {
	Token           string     `json:"token"`
	Infos           string     `json:"cartInfo"`
	NewItemQuantity uint       `json:"number"`
	Items           []CartItem `json:"items"`
	Pairs           []ItemPair `json:"goods"`
}

type CartItem struct {
	Gid             string   `json:"key"`
	Pic             string   `json:"pic"`
	Status          uint     `json:"status"` // === 1 【失效】
	Name            string   `json:"name"`
	Sku             []string `json:"sku"`
	Price           string   `json:"price"`
	Quantity        uint     `json:"number"`
	Selected        string   `json:"selected"`
	OptionValueName string   `json:"optionValueName"`
}

type ItemPair struct {
	GoodsId string `json:"goodsId"`
	Volume  uint   `json:"number"`
}

func (ci *CartInfo) NewCartItem(gid string, quantity uint) CartItem {
	sku := []string{"sku1", "sku3"}
	item := CartItem{gid, configs.Cfg.GoodsPicPrefix() + gid + "-01.jpeg", 0, "CD1.0", sku, "66.0", quantity, "1", "valueName"}
	return item
}
func (ci *CartInfo) getToken() string {
	return ci.Token
}

func (ci *CartInfo) AddMore(prod goods.GoodsDetail, quantity uint) {
	ci.NewItemQuantity = quantity
	if ci.Update(prod, quantity) {
		return
	}
	item := ci.createCartItem(prod, quantity)
	ip := ItemPair{prod.Gid, quantity}
	ci.Items = append(ci.Items, *item)
	ci.Pairs = append(ci.Pairs, ip)
}

func (ci *CartInfo) Update(prod goods.GoodsDetail, quantity uint) bool {
	for i := range ci.Items {
		it := &ci.Items[i]
		if it.Gid == prod.Gid {
			it.Quantity = it.Quantity + quantity
			ci.Pairs[i] = ItemPair{prod.Gid, it.Quantity}
			return true
		}
	}
	return false
}

func (c *CartInfo) getVolumeById(gid string) uint {
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

func (ci *CartInfo) createCartItem(prod goods.GoodsDetail, quantity uint) *CartItem {
	sku := []string{"sku1", "sku3"}
	selected := "1"
	optionValue := "optionValueName"

	item := CartItem{prod.Gid,
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
