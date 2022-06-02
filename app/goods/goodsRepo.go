package goods

import (
	"bookstore/app/configs"
	"sync"
)

var lockGR = &sync.Mutex{}
var goodsRepo *GoodsRepo

func GetGoodsRepo() *GoodsRepo {
	lockGR.Lock()
	defer lockGR.Unlock()
	if goodsRepo == nil {
		goodsRepo = &GoodsRepo{}
		goodsRepo.items = make([]GoodsItem, 0)
	}

	return goodsRepo
}

type GoodsRepo struct {
	items []GoodsItem
}

func (gr *GoodsRepo) GetItemDetail(id string) GoodsDetail {
	goods := gr.GetGoodsList()
	for _, item := range goods {
		if item.sameAs(id) {
			return item.GoodsDetail
		}
	}
	return GoodsDetail{}
}

func (gr *GoodsRepo) LoadGoods() []GoodsItem {
	if len(gr.items) != 0 {
		return gr.GetGoodsList()
	}
	// add the first catalouge
	item1 := gr.createGoods(0, 0, "g7225946", "持续交付1.0", 10, "册", "0", "一本DevOps的经典书。", uint(Saling), "66.0", "99.0", "1", "1")
	item2 := gr.createGoods(1, 0, "g7225947", "持续交付2.0", 20, "册", "0", "另一本DevOps的经典书。", uint(Saling), "99.0", "129.0", "1", "1")
	item3 := gr.createGoods(2, 0, "g7225948", "DevOps实战指南", 2, "册", "0", "DevOps 黄皮书。", uint(Saling), "55.0", "85.0", "1", "1")
	item4 := gr.createGoods(3, 0, "g7225949", "谷歌软件工程", 5, "册", "0", "解密硅谷头部互联网企业 如何打造软件工程文化。", uint(Saling), "77.0", "107.0", "1", "1")
	gr.items = append(gr.items, item1)
	gr.items = append(gr.items, item2)
	gr.items = append(gr.items, item3)
	gr.items = append(gr.items, item4)

	// add the second catalouge

	item5 := gr.createGoods(11, 1, "g1872110", "驾驭大数据", 20, "册", "0", "另一本DevOps的经典书。", uint(Saling), "99.0", "129.0", "1", "1")
	item6 := gr.createGoods(12, 1, "g1872111", "数据分析变革", 10, "册", "0", "一本DevOps的经典书。", uint(Saling), "66.0", "99.0", "1", "1")
	item7 := gr.createGoods(13, 1, "g1872112", "大数据测试技术与实践", 2, "册", "0", "DevOps 黄皮书。", uint(Saling), "55.0", "85.0", "1", "1")
	item8 := gr.createGoods(14, 1, "g1872113", "图解Spark 大数据快速分析实战", 5, "册", "0", "解密硅谷头部互联网企业 如何打造软件工程文化。", uint(Saling), "77.0", "107.0", "1", "1")

	gr.items = append(gr.items, item5)
	gr.items = append(gr.items, item6)
	gr.items = append(gr.items, item7)
	gr.items = append(gr.items, item8)
	return gr.items

}
func (gr *GoodsRepo) GetGoodsList() []GoodsItem {
	return gr.items
}

func (gr *GoodsRepo) createGoods(
	id uint, cateId uint,
	gid string, gName string, gStock uint,
	unit string, logistics string,
	desc string, status uint,
	minPrice string, origPrice string,
	afterSale string, recommandStatus string) GoodsItem {

	var detail = GoodsDetail{
		gid,       //"gId"
		gName,     //name
		nil,       //"Pics"
		0,         //"ItemId":
		gStock,    //"Stock":
		unit,      //Unit
		logistics, //"Logistics":
		desc,      //"Content":
		uint(Saling),
		Saling.StateStr(),
		configs.Cfg.GoodsPicPrefix() + gid + ".jpeg", //picURL
		minPrice,  //MinPrice
		origPrice, //OriginalPrice
		afterSale, //AfterSale
	}
	detail.setMultiPics(2)
	items := &GoodsItem{
		gid,             //id
		gName,           //name
		cateId,          //catalogueId
		recommandStatus, //recommandStatus
		configs.Cfg.GoodsPicPrefix() + gid + ".jpeg", //picURL
		minPrice,  //MinPrice
		origPrice, //originalPrice
		detail,
	}
	return *items
}
