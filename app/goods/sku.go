package goods

import "strings"

type SKU struct {
	SkuId            string            `json:"gid" gorm:"column:sku_id;primary_key;"`          //商品Id
	Name             string            `json:"name"`                                           // 商品名
	CategoryId       uint              `json:"catalogueId" gorm:"column:category_id"`          // 商品类别
	RecommendStatus  string            `json:"recommendStatus" gorm:"column:recommend_status"` // 商品推荐状态
	PicStr           string            `json:"pic" gorm:"column:pic_str"`                      //商品题图
	Unit             string            `json:"unit"`                                           // 数量单位
	Stock            uint              `json:"stores"`                                         // 库存 0: 无货，该商品已售罄
	MinPrice         string            `json:"minPrice" gorm:"column:min_price"`               //最低价格
	OriginalPrice    string            `json:"originalPrice" gorm:"column:original_price"`     //原价格
	Logistics        string            `json:"logistics"`                                      // 是否包邮 1:包邮,   0:不包邮, 空:无需配送
	Content          string            `json:"content"`                                        // 商品介绍
	Status           SalingStatus      `json:"status"`                                         // SKU状态 1: 商品已下架
	Aftersale        AfterSaleType     `json:"afterSale" gorm:"column:aftersale"`                     // 售后服务支持 1:支持退款且退货，0:支持退款, 2:支持退货
	SkuCarouPictures []SkuCarouPicture `gorm:"foreignKey:SkuId;references:SkuId"`
}

type SkuCarouPicture struct {
	Id     int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	SkuId  string `json:"gid" gorm:"column:sku_id"`
	PicStr string `json:"pic" gorm:"column:pic_str"`
}

type Tabler interface {
	TableName() string
}

func (s SkuCarouPicture) picId() string {
	vid := s.PicStr[0:strings.Index(s.PicStr, ".jpeg")]
	return s.SkuId + vid
}
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

type AfterSaleType uint

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
