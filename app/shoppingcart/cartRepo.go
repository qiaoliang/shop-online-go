package cart

import (
	"bookstore/app/configs"
)

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
	item := CartItem{key, configs.Cfg.StaticPicURI() + "/goods/g7227946-01.jpeg", 0, "CD1.0", sku, 66.0, quantity, "1", "valueName"}
	items := make([]CartItem, 0)
	items = append(items, item)
	ip := make([]ItemPair, 0)
	ip = append(ip, ItemPair{key, quantity})
	cartInfo := CartInfo{token, "iamInfos", quantity, items, ip}
	cs.cartInfos[token] = &cartInfo
	return &cartInfo
}
