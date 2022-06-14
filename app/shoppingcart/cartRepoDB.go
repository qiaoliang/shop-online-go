package cart

import (
	"bookstore/app/configs"
	"bookstore/app/goods"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type CartRepoIf interface {
	PutItemsInCart(token string, gid string, quantity uint) *CartInfoVM
	ModifyQuantityOfGoodsInCate(token string, gid string, quantity uint) *CartInfoVM
	GetCartByToken(token string) *CartInfoVM
	CreateCartInfoFor(token string, prod *goods.GoodsDetail, quantity uint) *CartInfoVM

	SaveUserCartItem(uci UserCartItem) error
	DeleteUserCartItem(uci UserCartItem) error
	DeleteUserCartItemsBy(token string) error
	UpdateUserCartItem(uci UserCartItem) error
	GetUserCartItem(uci UserCartItem) *UserCartItem
	FindUserCartItemsBy(token string) []UserCartItem
}

type CartRepoDB struct {
	cartInfos map[string]*CartInfoVM
	db        *gorm.DB
}

func init() {
	GetCartsRepo()
	gR := goods.GetGoodsRepo()
	gR.LoadGoods()
}
func GetCartsRepoIf() CartRepoIf {
	if cartRepo == nil {
		cartRepo = newCartsRepo(configs.Cfg.Persistence)
	}
	return cartRepo
}
func newCartsRepo(persistance bool) CartRepoIf {
	if persistance {
		return &CartRepoDB{make(map[string]*CartInfoVM, 0), configs.Cfg.DBConnection()}
	} else {
		return &CartRepo{make(map[string]*CartInfoVM, 0)}
	}
}

func (cs *CartRepoDB) SaveUserCartItem(uci UserCartItem) error {
	ret := cs.db.Create(&uci)
	return ret.Error
}
func (cs *CartRepoDB) DeleteUserCartItem(uci UserCartItem) error {
	log.Printf("uci token: %v\n", uci.Token)
	ret := cs.db.Where(map[string]interface{}{"Token": uci.Token, "sku_Id": uci.SkuId}).Delete(uci)
	return ret.Error
}
func (cs *CartRepoDB) DeleteUserCartItemsBy(token string) error {
	ret := cs.db.Where(map[string]interface{}{"Token": token}).Delete(&UserCartItem{})
	return ret.Error
}

func (cs *CartRepoDB) UpdateUserCartItem(uci UserCartItem) error {
	ret := cs.db.Where(map[string]interface{}{"Token": uci.Token, "sku_Id": uci.SkuId}).Select("*").Updates(&uci)
	return ret.Error
}
func (cs *CartRepoDB) GetUserCartItem(uci UserCartItem) *UserCartItem {
	found := UserCartItem{Token: uci.Token, SkuId: uci.SkuId}
	log.Printf("uci token: %v\n", uci.Token)
	cs.db.Where(&found).First(&found)
	return &found
}
func (cs *CartRepoDB) FindUserCartItemsBy(token string) []UserCartItem {
	found := []UserCartItem{}
	cs.db.Where(map[string]interface{}{"Token": token}).Find(&found)
	return found
}
func (cs *CartRepoDB) PutItemsInCart(token string, gid string, quantity uint) *CartInfoVM {
	goodsDetail := goods.GetGoodsRepo().GetItemDetail(gid)
	if goodsDetail == nil {
		log.Fatalf("sku %v is not found.\n", gid)
		return nil
	}
	if _, ok := cs.cartInfos[token]; !ok {
		return cs.CreateCartInfoFor(token, goodsDetail, quantity)
	}
	return nil
}
func (cs *CartRepoDB) ModifyQuantityOfGoodsInCate(token string, gid string, quantity uint) *CartInfoVM {

	goodsDetail := goods.GetGoodsRepo().GetItemDetail(gid)
	if goodsDetail == nil {
		fmt.Printf("没有找到 Gid 是 %v 的goodsDetail\n", gid)
	}
	if _, ok := cs.cartInfos[token]; !ok {
		fmt.Printf("没有找到 token: %v \n", token)
	}
	cs.cartInfos[token].Modify(goodsDetail, quantity)
	cs.cartInfos[token].caculateRedDot()
	return cs.cartInfos[token]
}

func (cs *CartRepoDB) GetCartByToken(token string) *CartInfoVM {
	if _, ok := cs.cartInfos[token]; !ok {
		return nil
	}
	cs.cartInfos[token].caculateRedDot()
	return cs.cartInfos[token]
}
func (cs *CartRepoDB) CreateCartInfoFor(token string, prod *goods.GoodsDetail, quantity uint) *CartInfoVM {
	items := make([]CartItemVM, 0)
	ips := make([]ItemPairVM, 0)
	ci := &CartInfoVM{token, quantity, items, ips}
	ci.AddMore(prod, quantity)
	ci.caculateRedDot()
	cs.cartInfos[token] = ci
	return ci

}
