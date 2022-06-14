package cart

import (
	"bookstore/app/goods"
	"fmt"
)

type CartRepo struct {
	cartInfos map[string]*CartInfoVM
}

var cartRepo CartRepoIf

func init() {
	GetCartsRepo()
	gR := goods.GetGoodsRepo()
	gR.LoadGoods()
}
func GetCartsRepo() CartRepoIf {
	if cartRepo == nil {
		cartRepo = &CartRepo{make(map[string]*CartInfoVM, 0)}
	}
	return cartRepo
}
func (cs *CartRepo) FindUserCartItemsBy(token string) []UserCartItem {
	//TODO: save to memory db, not implemented
	return nil
}

func (cs *CartRepo) DeleteUserCartItemsBy(token string) error {
	//TODO: save to memory db, not implemented
	return nil
}
func (cs *CartRepo) GetUserCartItem(uci UserCartItem) *UserCartItem {
	//TODO: save to memory db, not implemented
	return nil
}
func (cs *CartRepo) UpdateUserCartItem(uci UserCartItem) error {
	//TODO: save to memory db, not implemented
	return nil
}
func (cs *CartRepo) DeleteUserCartItem(uci UserCartItem) error {
	//TODO: save to memory db, not implemented
	return nil
}
func (cs *CartRepo) SaveUserCartItem(uci UserCartItem) error {
	//TODO: save to memory db, not implemented
	return nil
}
func (cs *CartRepo) PutItemsInCart(token string, gid string, quantity uint) *CartInfoVM {
	goodsDetail := goods.GetGoodsRepo().GetItemDetail(gid)
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
func (cs *CartRepo) ModifyQuantityOfGoodsInCate(token string, gid string, quantity uint) *CartInfoVM {

	goodsDetail := goods.GetGoodsRepo().GetItemDetail(gid)
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

func (cs *CartRepo) GetCartByToken(token string) *CartInfoVM {
	if _, ok := cs.cartInfos[token]; !ok {
		return nil
	}
	cs.cartInfos[token].caculateRedDot()
	return cs.cartInfos[token]
}
func (cs *CartRepo) CreateCartInfoFor(token string, prod *goods.GoodsDetail, quantity uint) *CartInfoVM {
	items := make([]CartItemVM, 0)
	ips := make([]ItemPairVM, 0)
	ci := &CartInfoVM{token, quantity, items, ips}
	ci.AddMore(prod, quantity)
	ci.caculateRedDot()
	cs.cartInfos[token] = ci
	return ci

}
