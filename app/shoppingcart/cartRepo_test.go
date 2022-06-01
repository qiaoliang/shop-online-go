package cart

import (
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CartRepositoryTestSuite struct {
	suite.Suite
}

func (st *CartRepositoryTestSuite) TestExample() {
	st.Equal(true, true)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CartRepositoryTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (st *CartRepositoryTestSuite) BeforeTest(suiteName, testName string) {

}

// This will run after test finishes
// and receives the suite and test names as input
func (st *CartRepositoryTestSuite) AfterTest(suiteName, testName string) {
	cartRepo = nil
}

// This will run before before the tests in the suite are run
func (st *CartRepositoryTestSuite) SetupSuite() {}

// This will run before each test in the suite
func (st *CartRepositoryTestSuite) SetupTest() {
	cartRepo = nil
	cartRepo = GetCartsInstance()
}

func (st *CartRepositoryTestSuite) Test_add_one_goods_into_an_empty_Cart() {

	expected := "IamTestToken"
	gid := "g7225946"
	c := cartRepo.AddOrderIntoCart("IamTestToken", gid, 10)
	st.Equal(expected, c.getToken(), "should get token from a cart, expected=%v actual=%v\n", expected, c.getToken())
	expVlm := uint(10)
	st.Equal(expVlm, c.getVolumeById(gid), "should put volume into a goodsid, expected=%v actual=%v\n", expVlm, c.getVolumeById(gid))
}
func (st *CartRepositoryTestSuite) Test_add_more_volume_into_a_Cart_with_same_goodsid() {
	gid := "g7225946"
	cartRepo.AddOrderIntoCart("sameGoodsId", gid, 10)

	c := cartRepo.AddOrderIntoCart("sameGoodsId", gid, 13)

	st.Equal(1, len(c.Items), "should be 1, but it was %v ", len(c.Items))
	st.Equal(uint(13), c.getVolumeById(gid), "should be 13, but it was %v", c.getVolumeById(gid))
}
func (st *CartRepositoryTestSuite) Test_add_another_goods_into_a_Cart_with_one_goods() {
	gid1 := "g7225946"
	gid2 := "g1872110"
	cartRepo.AddOrderIntoCart("firstGoodsId", gid1, 10)
	c := cartRepo.AddOrderIntoCart("secondGoodsId", gid2, 20)

	length := len(cartRepo.cartInfos)
	st.Equal(2, length, "should have 2 tokens , but it was %v ", length)
	st.Equal(uint(20), c.getVolumeById(gid2), "should be 20 for secondGoodsId, but it was %v", c.getVolumeById(gid2))
}
