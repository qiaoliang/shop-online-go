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
	//TODO: Get goods from repo.
	sku := []string{"sku1", "sku3"}
	item := CartItem{gid, configs.Cfg.GoodsPicPrefix() + gid + "-01.jpeg", 0, "CD1.0", sku, "66.0", quantity, "1", "valueName"}
	return item
}
func (ci *CartInfo) getToken() string {
	return ci.Token
}

func (ci *CartInfo) AddMore(gid string, quantity uint) {
	ci.NewItemQuantity = quantity
	for i := range ci.Items {
		it := &ci.Items[i]
		if it.Gid == gid {
			it.Quantity = it.Quantity + quantity
			ci.Pairs[i] = ItemPair{gid, it.Quantity}
			return
		}
	}
	fmt.Println("error. should not be here")
}

func (ci *CartInfo) Update(gid string, quantity uint) {
	for i := range ci.Items {
		it := &ci.Items[i]
		if it.Gid == gid {
			it.Quantity = quantity
			ci.Pairs[i] = ItemPair{gid, quantity}
			return
		}
	}
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
