package cart

import (
	"bookstore/app/configs"
	"fmt"
)

type CartInfo struct {
	Token           string     `json:"token"`
	Infos           string     `json:"cartInfo"`
	NewItemQuantity uint       `json:"number"`
	Items           []CartItem `json:"items"`
	Pairs           []ItemPair `json:"goods"`
}

type CartItem struct {
	Key             uint     `json:"key"`
	Gid             string   `json:"gid"`
	Pic             string   `json:"pic"`
	Status          uint     `json:"status"` // === 1 【失效】
	Name            string   `json:"name"`
	Sku             []string `json:"sku"`
	Price           uint     `json:"price"`
	Quantity        uint     `json:"number"`
	Selected        string   `json:"selected"`
	OptionValueName string   `json:"optionValueName"`
}

type ItemPair struct {
	GoodsId uint `json:"goodsId"`
	Volume  uint `json:"number"`
}

func (ci *CartInfo) NewCartItem(key uint, quantity uint) CartItem {
	sku := []string{"sku1", "sku3"}
	item := CartItem{key, "gid", configs.Cfg.GoodsPicPrefix() + "7225946-01.jpeg", 0, "CD1.0", sku, 66.0, quantity, "1", "valueName"}
	return item
}
func (ci *CartInfo) getToken() string {
	return ci.Token
}
func (ci *CartInfo) Update(key uint, quantity uint) {
	for i := range ci.Items {
		it := &ci.Items[i]
		if it.Key == key {
			fmt.Printf("find same token and key, %v and %v\n", key, quantity)
			it.Quantity = quantity
			ci.Pairs[i] = ItemPair{key, quantity}
			return
		}
	}
	ci.Items = append(ci.Items, ci.NewCartItem(key, quantity))
	ci.Pairs = append(ci.Pairs, ItemPair{key, quantity})
}

func (c *CartInfo) getVolumeById(id uint) uint {
	if len(c.Pairs) == 0 {
		return uint(0)
	}

	for _, item := range c.Pairs {
		v := item
		if v.GoodsId == id {
			return item.Volume
		}
	}
	return uint(0)
}
