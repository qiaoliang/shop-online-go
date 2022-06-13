package cart

import (
	"bookstore/app/configs"
	"bookstore/app/goods"
	"fmt"
)

var cartService *CartService

type CartService struct {
	cached CachedCart
	gs     *goods.GoodsService
	cr     CartRepoIf
}

type CachedCart map[string]*CartInfo

func (c CachedCart) get(token string) *CartInfo {
	if _, ok := c[token]; !ok {
		return nil
	}
	return c[token]
}
func (c CachedCart) update(token string, cart *CartInfo) {
	c[token] = cart
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
		cartService.init()
	}
	return cartService
}
func (cs *CartService) init() {
	//TODO: Load from persistance
}
func (cs *CartService) PutItemsInCart(token string, gid string, quantity uint) *CartInfo {
	goodsDetail := cs.gs.GetItemDetail(gid)
	if goodsDetail == nil {
		fmt.Println("sku is not found.")
		return nil
	}
	ci := cs.cached.get(token)
	if ci == nil {
		ci = cs.CreateCartInfoFor(token, goodsDetail, quantity)
		ci.caculateRedDot()
		cs.cached.update(token, ci)
		return ci
	}
	ci.AddMore(goodsDetail, quantity)
	ci.caculateRedDot()
	cs.cached.update(token, ci)
	return ci
}
func (cs *CartService) ModifyQuantityOfGoodsInCate(token string, gid string, quantity uint) *CartInfo {

	goodsDetail := cs.gs.GetItemDetail(gid)
	if goodsDetail == nil {
		fmt.Printf("～～没有找到 Gid是 %v 的goodsDetail", gid)
	}
	if _, ok := cs.cached[token]; !ok {
		fmt.Printf("～～没有找到 token：%v", token)
	}
	cs.cached[token].Modify(goodsDetail, quantity)
	cs.cached[token].caculateRedDot()
	return cs.cached.get(token)
}

func (cs *CartService) GetCartByToken(token string) *CartInfo {
	ret := cs.cached.get(token)
	if ret == nil {
		return nil
	}
	ret.caculateRedDot()
	cs.cached.update(token, ret)
	return ret
}
func (cs *CartService) CreateCartInfoFor(token string, prod *goods.GoodsDetail, quantity uint) *CartInfo {
	//TODO: Persistance
	items := make([]CartItem, 0)
	ips := make([]ItemPair, 0)
	ci := &CartInfo{token, quantity, items, ips}
	ci.AddMore(prod, quantity)
	ci.caculateRedDot()
	cs.cached.update(token, ci)
	return ci

}
