package goods

import (
	"bookstore/app/configs"
	"sync"
)

var lockGS sync.Mutex
var goodsService *GoodsService

func GetGoodsService() *GoodsService {
	lockGS.Lock()
	defer lockGS.Unlock()
	if goodsService == nil {
		goodsService = newGoodsService(configs.Cfg.Persistence)
	}
	return goodsService
}

func newGoodsService(usingDB bool) *GoodsService {
	if usingDB {
		db := configs.Cfg.DBConnection()
		repo := getSkuRepoDB(db)
		return &GoodsService{make([]GoodsItem, 0), &repo}
	}
	return &GoodsService{make([]GoodsItem, 0), nil}
}

type GoodsItems []GoodsItem

type GoodsService struct {
	items GoodsItems
	repo  *SkuRepoIf
}

func (gr *GoodsService) GetItemDetail(id string) *GoodsDetail {

	return nil
}

func (gs *GoodsService) LoadGoods() GoodsItems {
	return gs.items

}
func (gs *GoodsService) skuToGoodsItem(sku SKU) *GoodsItem {
	return nil
}
func (gr *GoodsService) GetGoodsList() GoodsItems {
	return gr.items
}
