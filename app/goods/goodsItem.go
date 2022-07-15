package goods

import (
	"fmt"

	"github.com/example/project/app/configs"
)

type GoodsItem struct {
	Gid             string      `json:"id"` //商品序号Id
	Name            string      `json:"name"`
	CategoryId      uint        `json:"catalogueId"`
	RecommendStatus string      `json:"recommendStatus"`
	PicUrl          string      `json:"pic"`
	MinPrice        string      `json:"minPrice"`
	OriginalPrice   string      `json:"originalPrice"`
	GoodsDetail     GoodsDetail `json:"goodsDetail"`
}

func (gi *GoodsItem) blongsTo(cateId uint) bool {
	return (cateId == gi.CategoryId)
}

func (gi *GoodsItem) sameAs(gId string) bool {
	return (gId == gi.Gid)
}

type GoodsDetail struct {
	Gid           string          `json:"id"`            // 商品Id
	Name          string          `json:"name"`          // 商品名
	Pics          []CarouselPicVM `json:"pics"`          // 商品组图
	ItemId        uint            `json:"goodsId"`       // 商品Id
	Stock         uint            `json:"stores"`        // 库存 0: 无货，该商品已售罄
	Unit          string          `json:"unit"`          // 数量单位
	Logistics     string          `json:"logistics"`     // 是否免运费
	Content       string          `json:"content"`       // 商品介绍
	Status        uint            `json:"status"`        //
	StatusStr     string          `json:"statusStr"`     //
	PicUrl        string          `json:"pic"`           // 商品大图
	MinPrice      string          `json:"minPrice"`      // 最低价格
	OriginalPrice string          `json:"originalPrice"` // 商品原价
	AfterSale     string          `json:"afterSale"`     // 售后服务支持 1:支持退款且退货，0:支持退款, 2:支持退货
}

type CarouselPicVM struct {
	Id  string `json:"id"`
	Pic string `json:"pic"`
}

type State uint

const (
	Pending State = iota
	Stopped
	Saling
)

func (s State) StateStr() string {
	switch s {
	case Pending:
		return "停售"
	case Stopped:
		return "下架"
	case Saling:
		return "在售"
	default:
		return "Unknown"
	}
}
func (gd *GoodsDetail) PicUrlSubstr() string {
	return gd.PicUrl[len(configs.Cfg.GoodsPicPrefix()):]
}
func (gd *GoodsDetail) setMultiPics(picNum int) {

	gd.Pics = make([]CarouselPicVM, 0)
	for i := 1; i <= picNum; i++ {
		id := gd.Gid + "-0" + fmt.Sprintf("%d", i)
		pic := CarouselPicVM{id, configs.Cfg.GoodsPicPrefix() + id + ".jpeg"}
		gd.Pics = append(gd.Pics, pic)
	}
}
