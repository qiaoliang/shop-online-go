package repository

import (
	"bookstore/app/models"
)

func InitGoodsList() [] models.GoodsItem {
	items := []models.GoodsItem{}

    store:=models.GoodsStockInfo{0,0,2,"状态信息码。1为已下架"};
    picA :=models.Picture{"g7227946-01","http://localhost:9090/pic/goods/g7227946-01.jpeg"}
    picB :=models.Picture{"g7227946-02","http://localhost:9090/pic/goods/g7227946-02.jpeg"}
    pics := make([]models.Picture,0)
    pics = append(pics,picA)
    pics = append(pics,picB)
    var detail = models.GoodsDetail{
			0,				//"Id"
			pics,			//"Pics"
			0,			//"ItemId":
			store,		//"StockInfo":
			"0",		//"Logistics":
			"contents", //"Content":
		}
	items1 := &models.GoodsItem{
		0,        			//id
		"持续交付 1.0",		//name
		0,        			//catalogueId
		"1", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225946.jpeg",//picURL
		"66.0",				//MinPrice
		"99.0",				//originalPrice
		detail,

	}
	items2 := &models.GoodsItem{
		1,        			//id
		"持续交付 2.0",	//name
		0,        			//catalogueId
		"2", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225947.jpeg",//picURL
		"77.0",				//MinPrice
		"109.0",			//originalPrice
		models.GoodsDetail{},
	}
	items3 := &models.GoodsItem{
		2,        			//id
		"DevOps 实战指南",	//name
		0,        			//catalogueId
		"3", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225948.jpeg",//picURL
		"60.0",				//MinPrice
		"89.0",			//originalPrice
		models.GoodsDetail{},
	}
	items4 := &models.GoodsItem{
		3,        			//id
		"谷歌软件工程",	//name
		0,        			//catalogueId
		"4", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225949.jpeg",//picURL
		"80.0",				//MinPrice
		"129.0",			//originalPrice
		models.GoodsDetail{},
	}
	items = append(items, *items1)
	items = append(items, *items2)
	items = append(items, *items3)
	items = append(items, *items4)


	// add the second catalouge

	items5 := &models.GoodsItem{
		10,        			//id
		"驾驭大数据",		//name
		1,        			//catalogueId
		"1", 				//recommandStatus
		"http://localhost:9090/pic/goods/g1872110.jpeg",//picURL
		"50.0",				//MinPrice
		"75.0",				//originalPrice
		models.GoodsDetail{},
	}
	items6 := &models.GoodsItem{
		11,        			//id
		"数据分析变革",		//name
		1,        			//catalogueId
		"2", 				//recommandStatus
		"http://localhost:9090/pic/goods/g1872111.jpeg",//picURL
		"41.0",				//MinPrice
		"65.0",			//originalPrice
		models.GoodsDetail{},
	}
	items7 := &models.GoodsItem{
		12,        			//id
		"大数据测试技术与实践",	//name
		1,        			//catalogueId
		"3", 				//recommandStatus
		"http://localhost:9090/pic/goods/g1872112.jpeg",//picURL
		"60.0",				//MinPrice
		"89.0",			//originalPrice
		models.GoodsDetail{},
	}
	items8 := &models.GoodsItem{
		13,        			//id
		"图解Spark 大数据快速分析实战",	//name
		1,        			//catalogueId
		"4", 				//recommandStatus
		"http://localhost:9090/pic/goods/g1872113.jpeg",//picURL
		"80.0",				//MinPrice
		"129.0",			//originalPrice
		models.GoodsDetail{},
	}
	items = append(items, *items5)
	items = append(items, *items6)
	items = append(items, *items7)
	items = append(items, *items8)
	return items
}