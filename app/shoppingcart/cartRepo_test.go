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
	cartRepo = nil
	cartRepo = GetCartsInstance()
}

// This will run after test finishes
// and receives the suite and test names as input
func (st *CartRepositoryTestSuite) AfterTest(suiteName, testName string) {}

// This will run before before the tests in the suite are run
func (st *CartRepositoryTestSuite) SetupSuite() {}

// This will run before each test in the suite
func (st *CartRepositoryTestSuite) SetupTest() {}

func (st *CartRepositoryTestSuite) Test_add_one_goods_into_an_empty_Cart() {

	expected := "IamTestToken"
	gid := "g7225946"
	c := cartRepo.AddOrderIntoCart("IamTestToken", 1, gid, 10)
	st.Equal(expected, c.getToken(), "should get token from a cart, expected=%v actual=%v\n", expected, c.getToken())
	expVlm := uint(10)
	st.Equal(expVlm, c.getVolumeById(1), "should put volume into a goodsid, expected=%v actual=%v\n", expVlm, c.getVolumeById(1))
}
func (st *CartRepositoryTestSuite) Test_add_more_volume_into_a_Cart_with_same_goodsid() {
	gid := "g7225946"
	cartRepo.AddOrderIntoCart("sameGoodsId", 1, gid, 10)

	c := cartRepo.AddOrderIntoCart("sameGoodsId", 1, gid, 13)

	st.Equal(1, len(c.Items), "should be 1, but it was %v ", len(c.Items))
	st.Equal(uint(13), c.getVolumeById(1), "should be 13, but it was %v", c.getVolumeById(1))
}
func (st *CartRepositoryTestSuite) Test_add_another_goods_into_a_Cart_with_one_goods() {
	gid1 := "g7225946"
	gid2 := "g1872110"
	cartRepo.AddOrderIntoCart("firstGoodsId", 1, gid1, 10)
	c := cartRepo.AddOrderIntoCart("secondGoodsId", 2, gid2, 20)

	length := len(cartRepo.cartInfos)
	st.Equal(2, length, "should have 2 tokens , but it was %v ", length)
	st.Equal(uint(20), c.getVolumeById(2), "should be 20 for secondGoodsId, but it was %v", c.getVolumeById(2))
}
