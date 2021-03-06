package goods

import (
	"strconv"
	"sync"

	"github.com/example/project/app/configs"
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

func newGoodsService(persistance bool) *GoodsService {

	return &GoodsService{make([]GoodsItem, 0), NewSkuRepo(persistance), NewCategoryRepo(persistance)}

}

type GoodsItems []GoodsItem

type GoodsService struct {
	items    GoodsItems
	repo     SkuRepoIf
	cateRepo CategoryRepoIf
}

func (gs *GoodsService) GetCategory(cId uint) GoodsItems {
	result := make(GoodsItems, 0)
	for _, v := range gs.items {
		if v.blongsTo(cId) {
			result = append(result, v)
		}
	}
	return result
}

func (gs *GoodsService) GetItemDetail(gid string) *GoodsDetail {
	ret := gs.getFromCache(gid)
	if ret != nil {
		return ret
	}
	sku := gs.repo.FindWithCarouselPics(gid)
	item := gs.skuToGoodsItem(*sku)
	return &item.GoodsDetail
}

func (gs *GoodsService) LoadGoods() GoodsItems {
	skus := gs.repo.FindAll()
	items := make(GoodsItems, 0)
	for _, sku := range skus {
		i := gs.skuToGoodsItem(sku)
		items = append(items, *i)
	}
	gs.items = items
	return items
}

func (gs *GoodsService) getFromCache(gid string) *GoodsDetail {
	for _, v := range gs.items {
		if v.sameAs(gid) {
			return &v.GoodsDetail
		}
	}
	return nil
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
		sku.Status.String(),
		configs.Cfg.GoodsPicPrefix() + sku.SkuId + ".jpeg", //picURL
		sku.MinPrice,                     //MinPrice
		sku.OriginalPrice,                //OriginalPrice
		strconv.Itoa(int(sku.Aftersale)), //AfterSale
	}
	gd.Pics = make([]CarouselPicVM, 0)
	for _, v := range sku.SkuCarouPictures {
		pic := gs.picToPicVM(v)
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

func (*GoodsService) picToPicVM(v SkuCarouPicture) CarouselPicVM {
	pic := CarouselPicVM{
		v.picId(),
		configs.Cfg.GoodsPicPrefix() + v.SkuId + v.PicStr,
	}
	return pic
}
