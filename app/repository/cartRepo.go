package repository

import "fmt"

type CartInfo struct {
	Token           string     `json:"token"`
	Infos           string     `json:"cartInfo"`
	NewItemQuantity uint       `json:"number"`
	Items           []CartItem `json:"items"`
	Pairs           []ItemPair `json:"goods"`
}

type CartItem struct {
	Key             uint     `json:"key"`
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
	item := CartItem{key, "http://localhost:9090/pic/goods/g7227946-01.jpeg", 0, "CD1.0", sku, 66.0, quantity, "1", "valueName"}
	return item
}
func (ci *CartInfo) getToken() string {
	return ci.Token
}
func (ci *CartInfo) Update(key uint, quantity uint) {
	for i, _ := range ci.Items {
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

type CartRepo struct {
	cartInfos map[string]*CartInfo
}

var cartRepo *CartRepo

func GetCartsInstance() *CartRepo {
	if cartRepo == nil {
		cartRepo = &CartRepo{make(map[string]*CartInfo, 0)}
	}
	return cartRepo
}

func (cs *CartRepo) AddOrderIntoCart(token string, goodsId uint, quantity uint) *CartInfo {

	if _, ok := cs.cartInfos[token]; !ok {
		cs.cartInfos[token] = cs.createCartInfo(token, goodsId, quantity)
		return cs.cartInfos[token]
	}
	cs.cartInfos[token].Update(goodsId, quantity)
	return cs.cartInfos[token]
}
func (cs *CartRepo) UpdateQuantityOfGoodsInCate(token string, goodsId uint, quantity uint) *CartInfo {

	if _, ok := cs.cartInfos[token]; !ok {
		cs.cartInfos[token] = cs.createCartInfo(token, goodsId, quantity)
		return cs.cartInfos[token]
	}
	cs.cartInfos[token].Update(goodsId, quantity)
	return cs.cartInfos[token]
}

func (cs *CartRepo) GetCartByToken(token string) *CartInfo {
	if _, ok := cs.cartInfos[token]; !ok {
		return nil
	}
	return cs.cartInfos[token]
}

func (cs *CartRepo) createCartInfo(token string, key uint, quantity uint) *CartInfo {
	sku := []string{"sku1", "sku3"}
	item := CartItem{key, "http://localhost:9090/pic/goods/g7227946-01.jpeg", 0, "CD1.0", sku, 66.0, quantity, "1", "valueName"}
	items := make([]CartItem, 0)
	items = append(items, item)
	ip := make([]ItemPair, 0)
	ip = append(ip, ItemPair{key, quantity})
	cartInfo := CartInfo{token, "iamInfos", quantity, items, ip}
	cs.cartInfos[token] = &cartInfo
	return &cartInfo
}
