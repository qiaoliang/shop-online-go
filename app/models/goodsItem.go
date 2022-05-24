package models

type GoodsItem struct {
	Id          	uint        `json:"id"`
 	Name			string		`json:"name"`
 	RecommendStatus string      `json:"recommendStatus"`
	PicUrl			string		`json:"pic"`
	MinPrice    	string    	`json:"minPrice"`
	OriginalPrice 	string 		`json:"originalPrice"`
}