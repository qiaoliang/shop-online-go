package cart

import (
	"bookstore/app/configs"
	"bookstore/app/goods"
	"log"
	"strings"
)

var cartService *CartService

type CartService struct {
	cached CachedCart
	gs     *goods.GoodsService
	cr     CartRepoIf
}

type CachedCart map[string]*CartInfoVM

func (c CachedCart) get(token string) *CartInfoVM {
	if _, ok := c[token]; !ok {
		return nil
	}
	return c[token]
}
func (c CachedCart) update(token string, cart *CartInfoVM) {
	c[token] = cart
}
func GetCartsService() CartRepoIf {
	if cartRepo == nil {
		cartService = newCartService(configs.Cfg.Persistence)
	}
	return cartRepo
}
func newCartService(persistance bool) *CartService {
	return &CartService{make(map[string]*CartInfoVM, 0), goods.GetGoodsService(), newCartsRepo(persistance)}
}

func (cs *CartService) PutItemsInCart(token string, skuId string, quantity uint) *CartInfoVM {
	goodsDetail := cs.gs.GetItemDetail(skuId)
	if goodsDetail == nil {
		log.Printf("～～没有找到 skuId 是 %v 的goodsDetail", skuId)
	}
	ci := cs.cached.get(token)
	if ci == nil {
		ci = cs.fetchCartItemsFromPersistance(token)
		if ci == nil {
			ci = cs.CreateCartInfoFor(token, goodsDetail, 0)
			return ci
		}
	}
	item, ip := ci.FindBy(goodsDetail.Gid)
	item.AddMore(quantity)
	ip.AddMore(quantity)
	cs.cr.UpdateUserCartItem(cs.fromVMToUCI(token, item))
	ci.caculateRedDot()
	cs.cached.update(token, ci)
	return ci
}
func (cs *CartService) ModifyQuantityOfGoodsInCate(token string, skuId string, quantity uint) *CartInfoVM {

	goodsDetail := cs.gs.GetItemDetail(skuId)
	if goodsDetail == nil {
		log.Printf("～～没有找到 skuId 是 %v 的goodsDetail", skuId)
	}
	ci := cs.cached.get(token)
	if ci == nil {
		ci = cs.fetchCartItemsFromPersistance(token)
		if ci == nil {
			return nil
		}
	}
	item := ci.Modify(goodsDetail, quantity)
	cs.cr.UpdateUserCartItem(cs.fromVMToUCI(token, *item))
	ci.caculateRedDot()
	return ci
}

func (cs *CartService) GetCartByToken(token string) *CartInfoVM {
	ret := cs.cached.get(token)
	if ret == nil {
		return cs.fetchCartItemsFromPersistance(token)
	}
	return ret
}
func (cs *CartService) fetchCartItemsFromPersistance(token string) *CartInfoVM {
	found := cs.cr.FindUserCartItemsBy(token)
	if len(found) == 0 {
		return nil
	}
	ci := cs.fromUCIsToVM(token, found)
	cs.cached.update(token, ci)
	return ci
}

func (cs *CartService) fromUCIsToVM(token string, found []UserCartItem) *CartInfoVM {
	ci := &CartInfoVM{token, 0, []CartItemVM{}, []ItemPairVM{}}
	for _, v := range found {
		item, ip := cs.fromUCIToVM(v)
		ci.Items = append(ci.Items, item)
		ci.Pairs = append(ci.Pairs, ip)
	}
	ci.caculateRedDot()
	return ci
}

func (cs *CartService) fromUCIToVM(uci UserCartItem) (CartItemVM, ItemPairVM) {
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
func (cs *CartService) fromVMToUCI(token string, ci CartItemVM) UserCartItem {
	uci := UserCartItem{
		token,
		ci.Gid,
		ci.RetrivePicStr(),
		ci.Status,
		ci.Name,
		strings.Join(ci.Sku, ","),
		ci.Price,
		ci.Quantity,
		ci.Selected,
		ci.OptionValueName,
	}
	return uci
}

func (cs *CartService) CreateCartInfoFor(token string, prod *goods.GoodsDetail, quantity uint) *CartInfoVM {
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
	ci := &CartInfoVM{token, 0, []CartItemVM{}, []ItemPairVM{}}
	item, ip := cs.fromUCIToVM(uci)
	ci.Items = append(ci.Items, item)
	ci.Pairs = append(ci.Pairs, ip)
	ci.RedDot = uint(len(ci.Items))
	cs.cached.update(token, ci)
	return ci
}
