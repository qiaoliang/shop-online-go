package models

type GoodsItem struct {
	Id          	uint        `json:"id"`
 	Name			string		`json:"name"`
 	CatalogueId		uint		`json:"catalogueId"`
 	RecommendStatus string      `json:"recommendStatus"`
	PicUrl			string		`json:"pic"`
	MinPrice    	string    	`json:"minPrice"`
	RiginalPrice 	string 		`json:"originalPrice"`
}