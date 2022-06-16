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
	sr     goods.SkuRepoIf
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
func GetCartsService() *CartService {
	if cartService == nil {
		cartService = newCartService(configs.Cfg.Persistence)
	}
	return cartService
}
func newCartService(persistance bool) *CartService {
	return &CartService{make(map[string]*CartInfoVM, 0), goods.NewSkuRepo(persistance), newCartsRepo(persistance)}
}

func (cs *CartService) PutItemsInCart(token string, skuId string, quantity uint) *CartInfoVM {
	sku := cs.sr.First(skuId)
	if sku == nil {
		log.Printf("～～没有找到 skuId 是 %v 的goodsDetail", skuId)
		return nil
	}
	ci := cs.findCartInfoFor(token)
	if ci == nil {
		return cs.CreateCartItemFor(token, sku, quantity)
	}

	item, ip := ci.FindDetailBy(skuId)
	if item == nil {
		return cs.CreateCartItemFor(token, sku, quantity)
	}

	item.AddMore(quantity)
	ip.AddMore(quantity)
	cs.cr.UpdateUserCartItem(cs.fromVMToUCI(token, item))

	ci.caculateRedDot()
	cs.cached.update(token, ci)
	return ci
}

func (cs *CartService) findCartInfoFor(token string) *CartInfoVM {
	ci := cs.cached.get(token)
	if ci == nil {
		ci = cs.fetchCartItems(token)
	}
	return ci
}
func (cs *CartService) ModifyQuantityOfGoodsInCate(token string, gid string, quantity uint) *CartInfoVM {

	gd := cs.sr.First(gid)
	if gd == nil {
		log.Printf("～～没有找到 skuId 是 %v 的goodsDetail", gid)
		return nil
	}
	ci := cs.findCartInfoFor(token)
	if ci == nil {
		return nil
	}
	item := ci.Modify(gid, quantity)
	cs.cr.UpdateUserCartItem(cs.fromVMToUCI(token, item))
	ci.caculateRedDot()
	return ci
}

func (cs *CartService) GetCartByToken(token string) *CartInfoVM {
	ret := cs.cached.get(token)
	if ret == nil {
		return cs.fetchCartItems(token)
	}
	return ret
}
func (cs *CartService) fetchCartItems(token string) *CartInfoVM {
	found := cs.cr.FindUserCartItemsBy(token)
	if len(found) == 0 {
		return nil
	}
	ci := cs.fromUCIlistToVM(token, found)
	cs.cached.update(token, ci)
	return ci
}

func (cs *CartService) fromUCIlistToVM(token string, found []UserCartItem) *CartInfoVM {
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
func (cs *CartService) fromVMToUCI(token string, ci *CartItemVM) *UserCartItem {
	return &UserCartItem{
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
}

func (cs *CartService) CreateCartItemFor(token string, sku *goods.SKU, quantity uint) *CartInfoVM {
	uci := UserCartItem{
		Token:           token,
		SkuId:           sku.SkuId,
		Pic:             sku.PicStr,
		Status:          uint(sku.Status),
		Name:            sku.Name,
		SkuStrs:         "sku1,sku3",
		Price:           sku.MinPrice,
		Quantity:        quantity,
		Selected:        "1",
		OptionValueName: "OptionValueName",
	}
	err := cs.cr.SaveUserCartItem(uci)
	if err != nil {
		log.Fatalf("%v \n save db has error:\n%v\n", uci, err)
		return nil
	}
	ci := cs.cached.get(token)
	if ci == nil {
		ci = &CartInfoVM{token, 0, []CartItemVM{}, []ItemPairVM{}}
		cs.cached.update(token, ci)
	}
	item, ip := cs.fromUCIToVM(uci)
	ci.Items = append(ci.Items, item)
	ci.Pairs = append(ci.Pairs, ip)
	ci.RedDot = uint(len(ci.Items))
	return ci
}
