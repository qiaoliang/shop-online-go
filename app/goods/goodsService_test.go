package goods

import (
	"bookstore/app/configs"
	"bookstore/app/testutils"
	"fmt"
	"testing"

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
func (s *GoodsServiceTestSuite) Test_SKU_to_Item() {
	exp := prepare_GoodsItem_cd20_with_pics()

	sku := prepareSku_Cd10_With_Pics()
	ret := s.serv.skuToGoodsItem(sku)
	s.EqualValues(exp, ret)
}

func prepare_GoodsItem_cd20_with_pics() GoodsItem {
	gid := "g7225946"
	gName := "持续交付1.0"
	minPrice := "66.0"
	origPrice := "99.0"
	gd := GoodsDetail{
		gid,            //"gId"
		gName,          //name
		nil,            //"Pics"
		0,              //"ItemId":
		110,            //"Stock":
		"册",            //Unit
		"1",            //"Logistics":
		"DevOps 的第一本书", //"Content":
		uint(SalingStatus(ONSAIL)),
		SalingStatus(ONSAIL).String(),
		"g7225946.jpeg",             //picURL
		minPrice,                    //MinPrice
		origPrice,                   //OriginalPrice
		string(AfterSaleType(BOTH)), //AfterSale
	}
	gd.Pics = make([]Picture, 0)
	for i := 1; i <= 2; i++ {
		id := gd.Gid + "-0" + fmt.Sprintf("%d", i)
		pic := Picture{id, configs.Cfg.GoodsPicPrefix() + id + ".jpeg"}
		gd.Pics = append(gd.Pics, pic)
	}
	item := GoodsItem{
		gid,   //id
		gName, //name
		0,     //catalogueId
		"1",   //recommandStatus
		configs.Cfg.GoodsPicPrefix() + gid + ".jpeg", //picURL
		minPrice,  //MinPrice
		origPrice, //originalPrice
		gd,
	}
	return item
}
