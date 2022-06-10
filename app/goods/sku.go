package goods

type SKU struct {
	SkuId            string            `json:"gid" gorm:"column:Sku_Id;primary_key;"`          //商品Id
	Name             string            `json:"name"`                                           // 商品名
	CategoryId       uint              `json:"catalogueId" gorm:"column:Category_Id"`          // 商品类别
	RecommendStatus  string            `json:"recommendStatus" gorm:"column:Recommend_Status"` // 商品推荐状态
	PicStr           string            `json:"pic" gorm:"column:Pic_Str"`                      //商品题图
	Unit             string            `json:"unit"`                                           // 数量单位
	Stock            uint              `json:"stores"`                                         // 库存 0: 无货，该商品已售罄
	MinPrice         string            `json:"minPrice" gorm:"column:Min_Price"`               //最低价格
	OriginalPrice    string            `json:"originalPrice" gorm:"column:Original_Price"`     //原价格
	Logistics        string            `json:"logistics"`                                      // 是否包邮 1:包邮,   0:不包邮, 空:无需配送
	Content          string            `json:"content"`                                        // 商品介绍
	Status           SalingStatus      `json:"status"`                                         // ？
	StatusStr        string            `json:"statusStr"  gorm:"-"`                            // SKU状态 1: 商品已下架
	Aftersale        AfterSaleType     `json:"afterSale" gorm:"Aftersale"`                     // 售后服务支持 1:支持退款且退货，0:支持退款, 2:支持退货
	SkuCarouPictures []SkuCarouPicture `gorm:"foreignKey:SkuId"`
}

type SkuCarouPicture struct {
	Id     int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	SkuId  string `json:"gid" gorm:"column:Sku_Id"`
	PicStr string `json:"pic" gorm:"column:Pic_Str"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (SkuCarouPicture) TableName() string {
	return "SkuCarouselPics"
}

type SalingStatus int32

const (
	ONSAIL SalingStatus = iota
	OFFSAIL
)

func (c SalingStatus) String() string {
	switch c {
	case ONSAIL:
		return "在售"
	case OFFSAIL:
		return "已下架"
	}
	return "N/A"
}

type AfterSaleType int32

const (
	REFUND AfterSaleType = iota
	BOTH
	RETURNED
)

func (c AfterSaleType) String() string {
	switch c {
	case REFUND:
		return "支持退款"
	case BOTH:
		return "支持退款且退货"
	case RETURNED:
		return "支持退货"
	}
	return "N/A"
}
