package models

type GoodsItem struct {
	Id          	uint        `json:"id"`
 	Name			string		`json:"name"`
 	CatalogueId		uint		`json:"catalogueId"`
 	RecommendStatus string      `json:"recommendStatus"`
	PicUrl			string		`json:"pic"`
	MinPrice    	string    	`json:"minPrice"`
	RiginalPrice 	string 		`json:"originalPrice"`
	GoodsDetail		GoodsDetail	`json:"goodsDetail"`
}
type GoodsDetail struct{
	Id          	uint    	    		`json:"id"`
	Pics	        []Picture				`json:"pics"`
	ItemId	        uint					`json:"goodsId"`
	StockInfo		GoodsStockInfo			`json:"basicInfo"`
	Logistics		string 					`json:"logistics"`
	Content			string					`json:"content"`
}

type GoodsStockInfo struct{
	Id  			uint  				`json:"id"`
	ItemId	        uint				`json:"goodsId"`
	Status 			uint 				`json:"status"`
	StatusStr		string				`json:"statusStr"`
}
type Picture struct{
	Id  			string  				`json:"id"`
	Pic	        	string				`json:"pic"`
}