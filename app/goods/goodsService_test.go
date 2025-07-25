package goods

import (
	"strconv"
	"testing"

	"bookstore/app/configs"
	"bookstore/app/testutils"

	"github.com/stretchr/testify/suite"
)

type GoodsServiceTestSuite struct {
	testutils.SupperSuite
	serv *GoodsService
}

func TestGoodsService(t *testing.T) {
	suite.Run(t, new(GoodsServiceTestSuite))
}

func (s *GoodsServiceTestSuite) BeforeTest(suiteName, testName string) {}

func (s *GoodsServiceTestSuite) AfterTest(suiteName, testName string) {}

func (s *GoodsServiceTestSuite) SetupSuite() {
	s.SupperSuite.SetupSuite()
	db := configs.Cfg.DBConnection()
	skuRepo := NewSkuRepoDB(db)
	cateRepo := &CategoryRepo{}
	s.serv = NewGoodsService(skuRepo, cateRepo)
}
func (s *GoodsServiceTestSuite) TeardownSuite() {
	s.SupperSuite.TeardownSuite()
	s.serv = nil
}

func (s *GoodsServiceTestSuite) SetupTest() {}

func (s *GoodsServiceTestSuite) Test_get_GoodsDetail() {
	result := s.serv.GetItemDetail("g7225946")
	exp := GoodsDetail{Gid: "g7225946", Name: "持续交付1.0", Pics: []CarouselPicVM{CarouselPicVM{Id: "g7225946-01", Pic: "http://localhost:9090/pic/goods/g7225946-01.jpeg"}, CarouselPicVM{Id: "g7225946-02", Pic: "http://localhost:9090/pic/goods/g7225946-02.jpeg"}}, ItemId: 0, Stock: 10, Unit: "本", Logistics: "1", Content: "这是第一本 DevOps 的书", Status: 0, StatusStr: "在售", PicUrl: "http://localhost:9090/pic/goods/g7225946.jpeg", MinPrice: "66.0", OriginalPrice: "99.0", AfterSale: "1"}
	s.Equal(exp, *result)
}

func (s *GoodsServiceTestSuite) Test_SKU_to_Item() {
	sku := s.serv.repo.FindWithCarouselPics("g7225946")
	item := s.serv.skuToGoodsItem(*sku)
	exp := &GoodsItem{Gid: "g7225946", Name: "持续交付1.0", CategoryId: 0, RecommendStatus: "1", PicUrl: "http://localhost:9090/pic/goods/g7225946.jpeg", MinPrice: "66.0", OriginalPrice: "99.0", GoodsDetail: GoodsDetail{Gid: "g7225946", Name: "持续交付1.0", Pics: []CarouselPicVM{CarouselPicVM{Id: "g7225946-01", Pic: "http://localhost:9090/pic/goods/g7225946-01.jpeg"}, CarouselPicVM{Id: "g7225946-02", Pic: "http://localhost:9090/pic/goods/g7225946-02.jpeg"}}, ItemId: 0, Stock: 10, Unit: "本", Logistics: "1", Content: "这是第一本 DevOps 的书", Status: 0, StatusStr: "在售", PicUrl: "http://localhost:9090/pic/goods/g7225946.jpeg", MinPrice: "66.0", OriginalPrice: "99.0", AfterSale: "1"}}
	s.Equal(exp, item)
}

func prepare_Devops_category() GoodsItems {
	items := make(GoodsItems, 0)
	cd10 := assemble_Item("g7225946", "持续交付1.0", 0, 110, "DevOps 的第一本书", "66.0", "99.0")
	cd20 := assemble_Item("g7225947", "持续交付2.0", 0, 200, "另一本DevOps的经典书。", "99.0", "129.0")
	devops := assemble_Item("g7225948", "DevOps实战指南", 0, 10, "DevOps 黄皮书。", "55.0", "85.0")
	ggSE := assemble_Item("g7225949", "谷歌软件工程", 0, 20, "解密硅谷头部互联网企业 如何打造软件工程文化。", "77.0", "107.0")
	items = append(items, cd10)
	items = append(items, cd20)
	items = append(items, devops)
	items = append(items, ggSE)
	return items
}
func prepare_GoodsItem_cd10_with_pics() GoodsItem {
	items := prepare_Devops_category()
	return items[0]
}

func assemble_Item(gid string, gName string, CatalogueId int, stock int, content string, minPrice string, origPrice string) GoodsItem {
	picExt := ".jpeg"
	gd := aDetail(gid, gName, stock, content, picExt, minPrice, origPrice)
	gd.Pics = append_pics_to_gd(gd.Gid)
	item := aItem(gid, gName, CatalogueId, picExt, minPrice, origPrice, gd)
	return item
}
func aDetail(gid string, gName string, stock int, content string, picExt string, minPrice string, origPrice string) GoodsDetail {
	gd := GoodsDetail{
		gid,
		gName,
		nil,
		0,
		uint(stock),
		"本",
		"1",
		content,
		uint(SalingStatus(ONSAIL)),
		SalingStatus(ONSAIL).String(),
		configs.Cfg.GoodsPicPrefix() + gid + picExt,
		minPrice,
		origPrice,
		strconv.Itoa(int(AfterSaleType(BOTH))),
	}
	return gd
}

func aItem(gid string, gName string, CatalogueId int, picExt string, minPrice string, origPrice string, gd GoodsDetail) GoodsItem {
	item := GoodsItem{
		gid,
		gName,
		uint(CatalogueId),
		"1",
		configs.Cfg.GoodsPicPrefix() + gid + picExt,
		minPrice,
		origPrice,
		gd,
	}
	return item
}

func append_pics_to_gd(gid string) []CarouselPicVM {
	pics := make([]CarouselPicVM, 0)
	pic1 := CarouselPicVM{gid + "-01", configs.Cfg.GoodsPicPrefix() + gid + "-01.jpeg"}
	pic2 := CarouselPicVM{gid + "-02", configs.Cfg.GoodsPicPrefix() + gid + "-02.jpeg"}
	pics = append(pics, pic1, pic2)
	return pics
}
