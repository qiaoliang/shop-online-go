package cart

import (
	"bookstore/app/goods"
)

type CartRepo struct {
	cartInfos map[string]*CartInfo
}

var cartRepo *CartRepo

func init() {
	GetCartsInstance()
}
func GetCartsInstance() *CartRepo {
	if cartRepo == nil {
		cartRepo = &CartRepo{make(map[string]*CartInfo, 0)}
	}
	return cartRepo
}

func (cs *CartRepo) AddOrderIntoCart(token string, gid string, quantity uint) *CartInfo {
	goodsDetail := goods.GetGoodsRepo().GetItemDetail(gid)
	if _, ok := cs.cartInfos[token]; !ok {
		cs.cartInfos[token] = cs.CreateCartInfoFor(token, goodsDetail, quantity)
		return cs.cartInfos[token]
	}
	cs.cartInfos[token].AddMore(goodsDetail, quantity)
	return cs.cartInfos[token]
}
func (cs *CartRepo) UpdateQuantityOfGoodsInCate(token string, gid string, quantity uint) *CartInfo {
	goodsDetail := goods.GetGoodsRepo().GetItemDetail(gid)
	cs.cartInfos[token].Update(goodsDetail, quantity)
	return cs.cartInfos[token]
}

func (cs *CartRepo) GetCartByToken(token string) *CartInfo {
	if _, ok := cs.cartInfos[token]; !ok {
		return nil
	}
	return cs.cartInfos[token]
}
func (cs *CartRepo) CreateCartInfoFor(token string, prod goods.GoodsDetail, quantity uint) *CartInfo {
	items := make([]CartItem, 0)
	ips := make([]ItemPair, 0)
	cs.cartInfos[token] = &CartInfo{token, quantity, items, ips}
	cs.cartInfos[token].AddMore(prod, quantity)
	return cs.cartInfos[token]

}
