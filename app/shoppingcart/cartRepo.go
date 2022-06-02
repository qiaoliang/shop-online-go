package cart

import (
	"bookstore/app/goods"
	"fmt"
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
		cs.cartInfos[token].caculateRedDot()
		return cs.cartInfos[token]
	}
	cs.cartInfos[token].AddMore(goodsDetail, quantity)
	cs.cartInfos[token].caculateRedDot()
	return cs.cartInfos[token]
}
func (cs *CartRepo) UpdateQuantityOfGoodsInCate(token string, gid string, quantity uint) *CartInfo {

	goodsDetail := goods.GetGoodsRepo().GetItemDetail(gid)
	if goodsDetail == nil {
		fmt.Printf("～～没有找到 Gid是 %v 的goodsDetail", gid)
	}
	if _, ok := cs.cartInfos[token]; !ok {
		fmt.Printf("～～没有找到 token：%v", token)
	}
	fmt.Printf("～～ token：%s, gid=%s\n", token, gid)
	for i, v := range cs.cartInfos[token].Items {
		fmt.Printf(" 打印已有的 Item %s:%d\n", v.Gid, v.Quantity)
		fmt.Printf(" 打印已有的 pair %s:%d\n", cs.cartInfos[token].Pairs[i].GoodsId, cs.cartInfos[token].Pairs[i].Volume)
	}
	cs.cartInfos[token].Update(goodsDetail, quantity)
	cs.cartInfos[token].caculateRedDot()
	return cs.cartInfos[token]
}

func (cs *CartRepo) GetCartByToken(token string) *CartInfo {
	if _, ok := cs.cartInfos[token]; !ok {
		return nil
	}
	cs.cartInfos[token].caculateRedDot()
	return cs.cartInfos[token]
}
func (cs *CartRepo) CreateCartInfoFor(token string, prod *goods.GoodsDetail, quantity uint) *CartInfo {
	items := make([]CartItem, 0)
	ips := make([]ItemPair, 0)
	ci := &CartInfo{token, quantity, items, ips}
	ci.AddMore(prod, quantity)
	ci.caculateRedDot()
	cs.cartInfos[token] = ci
	return ci

}
