package goods

import "bookstore/app/configs"

var goodsRepo GoodsRepo

type GoodsRepo struct {
	items []GoodsItem
}

func NewGoodsRepo() *GoodsRepo {
	goodsRepo := GoodsRepo{}
	goodsRepo.initRepo()
	return &goodsRepo
}
func (gr *GoodsRepo) initRepo() {
	goodsRepo.items = make([]GoodsItem, 0)
}
func (gr *GoodsRepo) creatData() []GoodsItem {

	picA := Picture{"g7227946-01", configs.Cfg.StaticPicURI() + "/goods/g7227946-01.jpeg"}
	picB := Picture{"g7227946-02", configs.Cfg.StaticPicURI() + "/goods/g7227946-02.jpeg"}
	pics := make([]Picture, 0)
	pics = append(pics, picA)
	pics = append(pics, picB)
	var detail = GoodsDetail{
		0,          //"Id"
		"持续交付 1.0", //name
		pics,       //"Pics"
		0,          //"ItemId":
		1,          //"Stock":
		"本",        //Unit
		"0",        //"Logistics":
		"contents", //"Content":
		2,
		"状态信息码。1为已下架",
		configs.Cfg.StaticPicURI() + "/goods/g7225946.jpeg", //picURL
		"66.0", //MinPrice
		"99.0", //OriginalPrice
		"1",    //AfterSale
	}
	items1 := &GoodsItem{
		0,          //id
		"持续交付 1.0", //name
		0,          //catalogueId
		"1",        //recommandStatus
		configs.Cfg.StaticPicURI() + "/goods/g7225946.jpeg", //picURL
		"66.0", //MinPrice
		"99.0", //originalPrice
		detail,
	}
	items2 := &GoodsItem{
		1,          //id
		"持续交付 2.0", //name
		0,          //catalogueId
		"2",        //recommandStatus
		configs.Cfg.StaticPicURI() + "/goods/g7225947.jpeg", //picURL
		"77.0",  //MinPrice
		"109.0", //originalPrice
		GoodsDetail{},
	}
	items3 := &GoodsItem{
		2,             //id
		"DevOps 实战指南", //name
		0,             //catalogueId
		"3",           //recommandStatus
		configs.Cfg.StaticPicURI() + "/goods/g7225948.jpeg", //picURL
		"60.0", //MinPrice
		"89.0", //originalPrice
		GoodsDetail{},
	}
	items4 := &GoodsItem{
		3,        //id
		"谷歌软件工程", //name
		0,        //catalogueId
		"4",      //recommandStatus
		configs.Cfg.StaticPicURI() + "/goods/g7225949.jpeg", //picURL
		"80.0",  //MinPrice
		"129.0", //originalPrice
		GoodsDetail{},
	}
	gr.items = append(gr.items, *items1)
	gr.items = append(gr.items, *items2)
	gr.items = append(gr.items, *items3)
	gr.items = append(gr.items, *items4)

	// add the second catalouge

	items5 := &GoodsItem{
		10,      //id
		"驾驭大数据", //name
		1,       //catalogueId
		"1",     //recommandStatus
		configs.Cfg.StaticPicURI() + "/goods/g1872110.jpeg", //picURL
		"50.0", //MinPrice
		"75.0", //originalPrice
		GoodsDetail{},
	}
	items6 := &GoodsItem{
		11,       //id
		"数据分析变革", //name
		1,        //catalogueId
		"2",      //recommandStatus
		configs.Cfg.StaticPicURI() + "/goods/g1872111.jpeg", //picURL
		"41.0", //MinPrice
		"65.0", //originalPrice
		GoodsDetail{},
	}
	items7 := &GoodsItem{
		12,           //id
		"大数据测试技术与实践", //name
		1,            //catalogueId
		"3",          //recommandStatus
		configs.Cfg.StaticPicURI() + "/goods/g1872112.jpeg", //picURL
		"60.0", //MinPrice
		"89.0", //originalPrice
		GoodsDetail{},
	}
	items8 := &GoodsItem{
		13, //id
		"图解Spark 大数据快速分析实战", //name
		1,   //catalogueId
		"4", //recommandStatus
		configs.Cfg.StaticPicURI() + "/goods/g1872113.jpeg", //picURL
		"80.0",  //MinPrice
		"129.0", //originalPrice
		GoodsDetail{},
	}
	gr.items = append(gr.items, *items5)
	gr.items = append(gr.items, *items6)
	gr.items = append(gr.items, *items7)
	gr.items = append(gr.items, *items8)
	return gr.GetGoodsList()

}
func (gr *GoodsRepo) GetGoodsList() []GoodsItem {
	return gr.items
}
func InitGoodsRepo() []GoodsItem {
	NewGoodsRepo()
	goodsRepo.initRepo()
	return goodsRepo.GetGoodsList()
}
