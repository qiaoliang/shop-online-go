package goods

type SKU struct {
	Id              string `gorm:"primary_key;"`    //商品Id
	Name            string `json:"name"`            // 商品名
	CategoryId      uint   `json:"catalogueId"`     // 商品类别
	RecommendStatus string `json:"recommendStatus"` // 商品推荐状态
	PicStr          string `json:"pic"`             //商品题图
	Unit            string `json:"unit"`            // 数量单位
	Stock           uint   `json:"stores"`          // 库存 0: 无货，该商品已售罄
	MinPrice        string `json:"minPrice"`        //最低价格
	OriginalPrice   string `json:"originalPrice"`   //原价格
	Logistics       string `json:"logistics"`       // 是否包邮 1:包邮,   0:不包邮, 空:无需配送
	Content         string `json:"content"`         // 商品介绍
	Status          uint   `json:"status"`          // ？
	StatusStr       string `json:"statusStr"`       // SKU状态 1: 商品已下架
	AfterSale       string `json:"afterSale"`       // 售后服务支持 1:支持退款且退货，0:支持退款, 2:支持退货
}

type CarouselPicture struct {
	Id     string //Pic Id
	SkuId  string //Sku Id
	PicStr string //Pic name
}
