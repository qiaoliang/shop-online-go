package cart

import (
	"log"
	"strings"

	"bookstore/app/goods"
)

var cartService *CartService

type CartService struct {
	goodsRepo goods.SkuRepoIf
	cr     *CartRepoDB
}

func NewCartService(goodsRepo goods.SkuRepoIf, cr *CartRepoDB) *CartService {
	return &CartService{goodsRepo, cr}
}

func (cs *CartService) PutItemsInCart(token string, skuId string, quantity uint) *CartInfoVM {
	sku := cs.goodsRepo.First(skuId)
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
	return ci
}

func (cs *CartService) findCartInfoFor(token string) *CartInfoVM {
	ci := cs.fetchCartItems(token)
	return ci
}
func (cs *CartService) ModifyQuantityOfGoodsInCate(token string, gid string, quantity uint) *CartInfoVM {

	gd := cs.goodsRepo.First(gid)
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
	ret := cs.fetchCartItems(token)
	return ret
}
func (cs *CartService) fetchCartItems(token string) *CartInfoVM {
	found := cs.cr.FindUserCartItemsBy(token)
	if len(found) == 0 {
		return nil
	}
	ci := cs.fromUCIlistToVM(token, found)
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
		ID: 0,
		Token: token,
		SkuId: ci.Gid,
		Pic: ci.RetrivePicStr(),
		Status: ci.Status,
		Name: ci.Name,
		SkuStrs: strings.Join(ci.Sku, ","),
		Price: ci.Price,
		Quantity: ci.Quantity,
		Selected: ci.Selected,
		OptionValueName: ci.OptionValueName,
	}
}

func (cs *CartService) CreateCartItemFor(token string, sku *goods.SKU, quantity uint) *CartInfoVM {
	uci := UserCartItem{
		ID: 0,
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
	// 关键：插入/更新后，直接重新从数据库加载购物车项，避免 items/goods 重复
	return cs.findCartInfoFor(token)
}
