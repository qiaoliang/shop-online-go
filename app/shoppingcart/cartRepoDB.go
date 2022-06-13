package cart

import (
	"bookstore/app/configs"
	"bookstore/app/goods"
	"fmt"

	"gorm.io/gorm"
)

type CartRepoIf interface {
	PutItemsInCart(token string, gid string, quantity uint) *CartInfo
	ModifyQuantityOfGoodsInCate(token string, gid string, quantity uint) *CartInfo
	GetCartByToken(token string) *CartInfo
	CreateCartInfoFor(token string, prod *goods.GoodsDetail, quantity uint) *CartInfo
	SaveUserCartItem(uci UserCartItem) error
}

type CartRepoDB struct {
	cartInfos map[string]*CartInfo
	db        *gorm.DB
}

func init() {
	GetCartsRepo()
	gR := goods.GetGoodsRepo()
	gR.LoadGoods()
}
func GetCartsRepoIf() CartRepoIf {
	if cartRepo == nil {
		cartRepo = NewCartsRepo(configs.Cfg.Persistence)
	}
	return cartRepo
}
func NewCartsRepo(persistance bool) CartRepoIf {
	if cartRepo == nil {
		if persistance {
			cartRepo = &CartRepoDB{make(map[string]*CartInfo, 0), configs.Cfg.DBConnection()}
		} else {
			cartRepo = &CartRepo{make(map[string]*CartInfo, 0)}
		}
	}
	return cartRepo
}

func (cs *CartRepoDB) SaveUserCartItem(uci UserCartItem) error {
	ret := cs.db.Create(&uci)
	return ret.Error
}

func (cs *CartRepoDB) PutItemsInCart(token string, gid string, quantity uint) *CartInfo {
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
func (cs *CartRepoDB) ModifyQuantityOfGoodsInCate(token string, gid string, quantity uint) *CartInfo {

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

func (cs *CartRepoDB) GetCartByToken(token string) *CartInfo {
	if _, ok := cs.cartInfos[token]; !ok {
		return nil
	}
	cs.cartInfos[token].caculateRedDot()
	return cs.cartInfos[token]
}
func (cs *CartRepoDB) CreateCartInfoFor(token string, prod *goods.GoodsDetail, quantity uint) *CartInfo {
	items := make([]CartItemVM, 0)
	ips := make([]ItemPairVM, 0)
	ci := &CartInfo{token, quantity, items, ips}
	ci.AddMore(prod, quantity)
	ci.caculateRedDot()
	cs.cartInfos[token] = ci
	return ci

}
