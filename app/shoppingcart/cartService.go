package cart

import (
	"bookstore/app/configs"
	"bookstore/app/goods"
	"fmt"
)

var cartService *CartService

type CartService struct {
	cartInfos map[string]*CartInfo
	gs        *goods.GoodsService
	cr        CartRepoIf
}

func GetCartsService() CartRepoIf {
	if cartRepo == nil {
		cartService = NewCartService(configs.Cfg.Persistence)
	}
	return cartRepo
}
func NewCartService(persistance bool) *CartService {
	if cartService == nil {
		if persistance {
			cartService = &CartService{make(map[string]*CartInfo, 0), goods.GetGoodsService(), NewCartsRepo(persistance)}
		} else {
			cartService = &CartService{make(map[string]*CartInfo, 0), nil, nil}
		}
	}
	return cartService
}

func (cs *CartService) PutItemsInCart(token string, gid string, quantity uint) *CartInfo {
	goodsDetail := cs.gs.GetItemDetail(gid)
	if goodsDetail == nil {
		fmt.Println("goodsDetail is nil")
		return nil
	}
	if _, ok := cs.cartInfos[token]; !ok {
		cs.cartInfos[token] = cs.CreateCartInfoFor(token, goodsDetail, quantity)
		cs.cartInfos[token].caculateRedDot()
		return cs.cartInfos[token]
	}
	cs.cartInfos[token].AddMore(goodsDetail, quantity)
	cs.cartInfos[token].caculateRedDot()
	return cs.cartInfos[token]
}
func (cs *CartService) ModifyQuantityOfGoodsInCate(token string, gid string, quantity uint) *CartInfo {

	goodsDetail := cs.gs.GetItemDetail(gid)
	if goodsDetail == nil {
		fmt.Printf("～～没有找到 Gid是 %v 的goodsDetail", gid)
	}
	if _, ok := cs.cartInfos[token]; !ok {
		fmt.Printf("～～没有找到 token：%v", token)
	}
	cs.cartInfos[token].Modify(goodsDetail, quantity)
	cs.cartInfos[token].caculateRedDot()
	return cs.cartInfos[token]
}

func (cs *CartService) GetCartByToken(token string) *CartInfo {
	if _, ok := cs.cartInfos[token]; !ok {
		return nil
	}
	cs.cartInfos[token].caculateRedDot()
	return cs.cartInfos[token]
}
func (cs *CartService) CreateCartInfoFor(token string, prod *goods.GoodsDetail, quantity uint) *CartInfo {
	items := make([]CartItem, 0)
	ips := make([]ItemPair, 0)
	ci := &CartInfo{token, quantity, items, ips}
	ci.AddMore(prod, quantity)
	ci.caculateRedDot()
	cs.cartInfos[token] = ci
	return ci

}
