package cart

import (
	"bookstore/app/configs"
	"log"

	"gorm.io/gorm"
)

type CartRepoIf interface {
	SaveUserCartItem(uci UserCartItem) error
	DeleteUserCartItem(uci UserCartItem) error
	DeleteUserCartItemsBy(token string) error
	UpdateUserCartItem(uci *UserCartItem) error
	GetUserCartItem(uci UserCartItem) *UserCartItem
	FindUserCartItemsBy(token string) []UserCartItem
}

type CartRepoDB struct {
	cartInfos map[string]*CartInfoVM
	db        *gorm.DB
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
		return &CartRepoMem{make(map[string]*UserCart, 0)}
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

func (cs *CartRepoDB) UpdateUserCartItem(uci *UserCartItem) error {
	ret := cs.db.Where(map[string]interface{}{"Token": uci.Token, "sku_Id": uci.SkuId}).Select("*").Updates(uci)
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
