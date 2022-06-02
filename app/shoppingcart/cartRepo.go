package cart

import (
	"bookstore/app/configs"
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
	if _, ok := cs.cartInfos[token]; !ok {
		goodsDetail := goods.GetGoodsRepo().GetItemDetail(gid)
		cs.cartInfos[token] = cs.CreateCartInfoFor(token, goodsDetail, quantity)
		return cs.cartInfos[token]
	}
	fmt.Println("is here.")
	cs.cartInfos[token].AddMore(gid, quantity)
	return cs.cartInfos[token]
}
func (cs *CartRepo) UpdateQuantityOfGoodsInCate(token string, gid string, quantity uint) *CartInfo {
	cs.cartInfos[token].Update(gid, quantity)
	return cs.cartInfos[token]
}

func (cs *CartRepo) GetCartByToken(token string) *CartInfo {
	if _, ok := cs.cartInfos[token]; !ok {
		return nil
	}
	return cs.cartInfos[token]
}
func (cs *CartRepo) CreateCartInfoFor(token string, prod goods.GoodsDetail, quantity uint) *CartInfo {
	sku := []string{"sku1", "sku3"}
	selected := "1"
	optionValue := "optionValueName"

	item := CartItem{prod.Gid,
		configs.Cfg.GoodsPicPrefix() + prod.Gid + ".jpeg",
		0,
		prod.Name,
		sku,
		prod.MinPrice,
		quantity,
		selected,
		optionValue,
	}
	items := make([]CartItem, 0)
	items = append(items, item)
	ip := make([]ItemPair, 0)
	ip = append(ip, ItemPair{prod.Gid, quantity})
	cartInfo := CartInfo{token, "someThing_no_use", quantity, items, ip}
	cs.cartInfos[token] = &cartInfo
	return &cartInfo

}
