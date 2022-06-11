package goods

import (
	"bookstore/app/configs"
	"strings"
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
	gd := GoodsDetail{
		sku.SkuId,     //"gId"
		sku.Name,      //name
		nil,           //"Pics"
		0,             //"ItemId":
		sku.Stock,     //"Stock":
		sku.Unit,      //Unit
		sku.Logistics, //"Logistics":
		sku.Content,   //"Content":
		uint(sku.Status),
		sku.StatusStr,
		configs.Cfg.GoodsPicPrefix() + sku.SkuId + ".jpeg", //picURL
		sku.MinPrice,          //MinPrice
		sku.OriginalPrice,     //OriginalPrice
		string(sku.Aftersale), //AfterSale
	}
	gd.Pics = make([]Picture, 0)
	for _, v := range sku.SkuCarouPictures {
		vid := v.PicStr[0:strings.Index(v.PicStr, ".jpeg")]
		id := v.SkuId + vid
		pic := Picture{id, configs.Cfg.GoodsPicPrefix() + v.SkuId + v.PicStr}
		gd.Pics = append(gd.Pics, pic)
	}

	items := &GoodsItem{
		sku.SkuId,           //id
		sku.Name,            //name
		sku.CategoryId,      //catalogueId
		sku.RecommendStatus, //recommandStatus
		configs.Cfg.GoodsPicPrefix() + sku.PicStr, //picURL
		sku.MinPrice,      //MinPrice
		sku.OriginalPrice, //originalPrice
		gd,
	}

	return items
}
func (gr *GoodsService) GetGoodsList() GoodsItems {
	return gr.items
}
