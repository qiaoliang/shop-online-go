package models

type GoodsItem struct {
	Id          	uint        `json:"id"`
 	Name			string		`json:"name"`
 	CatalogueId		uint		`json:"catalogueId"`
 	RecommendStatus string      `json:"recommendStatus"`
	PicUrl			string		`json:"pic"`
	MinPrice    	string    	`json:"minPrice"`
	RiginalPrice 	string 		`json:"originalPrice"`
	GoodsDetail		GoodsDetail
}
type GoodsDetail struct{
	Id          	uint    	    		`json:"id"`
	Pics	        []string				`json:"pics"`
	StockInfo		GoodsStockInfo			`json:"BasicInfo"`
	Logistics		string 					`json:"logistics"`
	Content			string					`json:"content"`
}

type GoodsStockInfo struct{
	Id  			uint  				`json:"id"`
	Status 			uint 				`json:"status"`
	StatusStr		string				`json:"statusStr"`
}

func InitGoodsList() [] GoodsItem {
	items := []GoodsItem{}
	items1 := &GoodsItem{
		0,        			//id
		"持续交付 1.0 ",		//name
		0,        			//catalogueId
		"1", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225946.jpeg",//picURL
		"66.0",				//MinPrice
		"99.0",				//originalPrice
		GoodsDetail{},
	}
	items2 := &GoodsItem{
		1,        			//id
		"持续交付 2.0",	//name
		0,        			//catalogueId
		"2", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225947.jpeg",//picURL
		"77.0",				//MinPrice
		"109.0",			//originalPrice
		GoodsDetail{},
	}
	items3 := &GoodsItem{
		2,        			//id
		"DevOps 实战指南",	//name
		0,        			//catalogueId
		"3", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225948.jpeg",//picURL
		"60.0",				//MinPrice
		"89.0",			//originalPrice
		GoodsDetail{},
	}
	items4 := &GoodsItem{
		3,        			//id
		"谷歌软件工程",	//name
		0,        			//catalogueId
		"4", 				//recommandStatus
		"http://localhost:9090/pic/goods/g7225949.jpeg",//picURL
		"80.0",				//MinPrice
		"129.0",			//originalPrice
		GoodsDetail{},
	}
	items = append(items, *items1)
	items = append(items, *items2)
	items = append(items, *items3)
	items = append(items, *items4)


	// add the second catalouge

	items5 := &GoodsItem{
		10,        			//id
		"驾驭大数据",		//name
		1,        			//catalogueId
		"1", 				//recommandStatus
		"http://localhost:9090/pic/goods/g1872110.jpeg",//picURL
		"50.0",				//MinPrice
		"75.0",				//originalPrice
		GoodsDetail{},
	}
	items6 := &GoodsItem{
		11,        			//id
		"数据分析变革",		//name
		1,        			//catalogueId
		"2", 				//recommandStatus
		"http://localhost:9090/pic/goods/g1872111.jpeg",//picURL
		"41.0",				//MinPrice
		"65.0",			//originalPrice
		GoodsDetail{},
	}
	items7 := &GoodsItem{
		12,        			//id
		"大数据测试技术与实践",	//name
		1,        			//catalogueId
		"3", 				//recommandStatus
		"http://localhost:9090/pic/goods/g1872112.jpeg",//picURL
		"60.0",				//MinPrice
		"89.0",			//originalPrice
		GoodsDetail{},
	}
	items8 := &GoodsItem{
		13,        			//id
		"图解Spark 大数据快速分析实战",	//name
		1,        			//catalogueId
		"4", 				//recommandStatus
		"http://localhost:9090/pic/goods/g1872113.jpeg",//picURL
		"80.0",				//MinPrice
		"129.0",			//originalPrice
		GoodsDetail{},
	}
	items = append(items, *items5)
	items = append(items, *items6)
	items = append(items, *items7)
	items = append(items, *items8)
	return items
}
