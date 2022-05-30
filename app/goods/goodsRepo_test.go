package goods

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GoodsRepositoryTestSuite struct {
	suite.Suite
}

func TestGoodsRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(GoodsRepositoryTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (st *GoodsRepositoryTestSuite) BeforeTest(suiteName, testName string) {

}

func (st *GoodsRepositoryTestSuite) AfterTest(suiteName, testName string) {}

func (st *GoodsRepositoryTestSuite) SetupSuite() {
}

func (st *GoodsRepositoryTestSuite) SetupTest() {
	goodsRepo = nil
	GetGoodsRepo()
	configs.GetConfigInstance(utils.GetConfigFileForTest())
}

func (st *GoodsRepositoryTestSuite) Test_should_initial_empty_goodRepo() {
	st.Equal(0, len(goodsRepo.GetGoodsList()))
}

func (st *GoodsRepositoryTestSuite) Test_goods_load_from_() {
	goodsRepo.loadGoods()
	st.Equal(8, len(goodsRepo.GetGoodsList()))
}
func (st *GoodsRepositoryTestSuite) Test_Create_goods_for_Shop() {
	item := goodsRepo.createGoods(0, 0, "g7227946", "持续交付1.0", 10, "册", "0", "一本DevOps的经典书。", uint(Saling), "66.0", "99.0", "1", "1")
	st.Equal("http://localhost:9090/pic/goods/g7227946.jpeg", item.PicUrl)
	pics := item.GoodsDetail.Pics
	st.Equal(2, len(pics))
	st.Equal(pics[0].Id, "g7227946-01")
	st.Contains(pics[0].Pic, "g7227946-01")
	st.Equal(pics[1].Id, "g7227946-02")
	st.Contains(pics[1].Pic, "g7227946-02")

}

func (st *GoodsRepositoryTestSuite) Test_Get_a_goods_detail() {
	goodsRepo.loadGoods()
	g := goodsRepo.getItemDetail(0, "token")
	st.Equal(uint(0), g.Id)
}
