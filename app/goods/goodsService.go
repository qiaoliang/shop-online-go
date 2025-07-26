package goods

import (
	"strconv"
	"sync"

	"bookstore/app/configs"
	"strings"
)

var lockGS sync.Mutex
var goodsService *GoodsService = newGoodsService()

func GetGoodsService() *GoodsService {
	lockGS.Lock()
	defer lockGS.Unlock()
	if goodsService == nil {
		goodsService = newGoodsService()
	}
	return goodsService
}

func NewGoodsService(repo *SkuRepoDB, cateRepo *CategoryRepoDB) *GoodsService {
	return &GoodsService{repo: repo, cateRepo: cateRepo}
}

func newGoodsService() *GoodsService {
	return &GoodsService{repo: &SkuRepoDB{}, cateRepo: &CategoryRepoDB{}}
}

// 删除原有的 NewGoodsService（含 &SkuRepoMem{}）实现和 gs.items 相关代码，只保留数据库 repo 版本。

type GoodsItems []GoodsItem

type GoodsService struct {
	repo     *SkuRepoDB
	cateRepo *CategoryRepoDB
}

// 只保留数据库 repo 相关方法
func (gs *GoodsService) GetItemDetail(gid string) *GoodsDetail {
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
	return items
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
	// id = skuId + "-01"，如 g7225946-01
	picId := v.SkuId + strings.TrimSuffix(v.PicStr, ".jpeg")
	picUrl := configs.Cfg.GoodsPicPrefix() + v.SkuId + v.PicStr
	return CarouselPicVM{
		picId,
		picUrl,
	}
}
