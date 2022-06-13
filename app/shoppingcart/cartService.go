package cart

import (
	"bookstore/app/configs"
	"bookstore/app/goods"
	"fmt"
	"log"
	"strings"
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
	ret := cs.cached.get(token)
	if ret == nil {
		fmt.Printf("～～没有找到 token：%v", token)
	}
	ret.Modify(goodsDetail, quantity)
	ret.caculateRedDot()
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
func (cs *CartService) UserCartItemToVM(uci UserCartItem) (CartItemVM, ItemPairVM) {
	civm := CartItemVM{
		uci.SkuId,
		uci.FullPicStr(),
		uci.Status,
		uci.Name,
		strings.Split(uci.SkuStrs, ","),
		uci.Price,
		uci.Quantity,
		uci.Selected,
		uci.OptionValueName,
	}
	ipvm := ItemPairVM{
		uci.SkuId,
		uci.Quantity,
	}
	return civm, ipvm
}

func (cs *CartService) CreateCartInfoFor(token string, prod *goods.GoodsDetail, quantity uint) *CartInfo {

	return cs.saveCartInfo(token, prod, quantity)

}
func (cs *CartService) saveCartInfo(token string, prod *goods.GoodsDetail, quantity uint) *CartInfo {
	uci := UserCartItem{
		Token:           token,
		SkuId:           prod.Gid,
		Pic:             prod.PicUrl,
		Status:          prod.Status,
		Name:            prod.Name,
		SkuStrs:         "sku1,sku3",
		Price:           prod.MinPrice,
		Quantity:        quantity,
		Selected:        "1",
		OptionValueName: "OptionValueName",
	}
	err := cs.cr.SaveUserCartItem(uci)
	if err != nil {
		log.Fatalf("%v \n save db has error:\n%v\n", uci, err)
		return nil
	}
	ci := &CartInfo{token, 0, []CartItemVM{}, []ItemPairVM{}}
	item, ip := cs.UserCartItemToVM(uci)
	ci.Items = append(ci.Items, item)
	ci.Pairs = append(ci.Pairs, ip)
	ci.RedDot = uint(len(ci.Items))
	cs.cached.update(token, ci)
	return ci
}
