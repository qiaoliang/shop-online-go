package goods

import (
	"log"
	"strconv"
	"testing"

	"github.com/example/project/app/configs"
	"github.com/example/project/app/testutils"

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
	s.serv = newGoodsService(true)
}
func (s *GoodsServiceTestSuite) TeardownSuite() {
	s.SupperSuite.TeardownSuite()
	s.serv = nil
}

func (s *GoodsServiceTestSuite) SetupTest() {}

func (s *GoodsServiceTestSuite) Test_get_GoodsDetail() {
	exp := prepare_GoodsItem_cd10_with_pics()
	ret := s.serv.GetItemDetail(exp.Gid)
	s.NotNil(ret)
	s.EqualValues(exp.GoodsDetail, *ret)
}
func (s *GoodsServiceTestSuite) Test_get_GoodsDetail_from_cache() {
	cachedItem := prepare_GoodsItem_cd10_with_pics()
	cachedItem.Gid = "IamCached"
	cachedItem.GoodsDetail.Gid = "IamCached"
	s.serv.items = append(s.serv.items, cachedItem)

	ret := s.serv.GetItemDetail("IamCached")
	s.NotNil(ret)
	s.EqualValues(cachedItem.GoodsDetail, *ret)

}

func (s *GoodsServiceTestSuite) Test_get_Goods_for_a_category() {
	s.serv.LoadGoods()
	log.Printf("%v\n", s.serv.items)
	result := s.serv.GetCategory(uint(0))
	s.Equal(4, len(result))

	exp := prepare_Devops_category()
	s.EqualValues(exp, result)
}
func (s *GoodsServiceTestSuite) Test_SKU_to_Item() {
	exp := prepare_GoodsItem_cd10_with_pics()

	sku := prepareSku_Cd10_With_Pics()
	ret := s.serv.skuToGoodsItem(sku)
	s.EqualValues(&exp, ret)
}

func (s *GoodsServiceTestSuite) Test_Load_GoodsItems() {
	ret := s.serv.LoadGoods()
	s.Equal(8, len(ret))
	exp := prepare_GoodsItem_cd10_with_pics()
	var r GoodsItem
	found := false
	for _, v := range ret {
		if v.Gid == exp.Gid {
			found = true
			r = v
			break
		}
	}
	s.True(found)
	s.EqualValues(exp, r)
}

func prepare_Devops_category() GoodsItems {
	items := make(GoodsItems, 0)
	cd10 := assemble_Item("g7225946", "????????????1.0", 0, 110, "DevOps ???????????????", "66.0", "99.0")
	cd20 := assemble_Item("g7225947", "????????????2.0", 0, 200, "?????????DevOps???????????????", "99.0", "129.0")
	devops := assemble_Item("g7225948", "DevOps????????????", 0, 10, "DevOps ????????????", "55.0", "85.0")
	ggSE := assemble_Item("g7225949", "??????????????????", 0, 20, "????????????????????????????????? ?????????????????????????????????", "77.0", "107.0")
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
		"???",
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
