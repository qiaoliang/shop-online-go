package cart

import (
	"bookstore/app/goods"
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CartRepositoryTestSuite struct {
	suite.Suite
	gRepo *goods.GoodsRepo
}

func (st *CartRepositoryTestSuite) TestExample() {
	st.Equal(true, true)
}

func TestCartRepositoryTestSuite(t *testing.T) {
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
	st.gRepo = goods.GetGoodsRepo()
	st.gRepo.LoadGoods()
}

func (st *CartRepositoryTestSuite) Test_Create_Cart_with_goods_For_a_Token() {
	token := "13900007997"
	gid := "g7225946"
	quantity := uint(100)
	prod := st.gRepo.GetItemDetail(gid)
	st.Equal(gid, prod.Gid)
	cartRepo.CreateCartInfoFor(token, prod, quantity)
	cart := cartRepo.GetCartByToken(token)
	st.NotNil(cart)
	items := cart.Items
	st.Equal(1, len(items), "should have only one item.")
	st.Equal(gid, items[0].Gid)
	st.Equal(quantity, items[0].Quantity)
}

func (st *CartRepositoryTestSuite) Test_add_one_goods_into_an_empty_Cart() {

	expected := "IamTestToken"
	gid := "g7225946"
	c := cartRepo.AddOrderIntoCart("IamTestToken", gid, 10)
	st.Equal(expected, c.getToken(), "should get token from a cart, expected=%v actual=%v\n", expected, c.getToken())
	expVlm := uint(10)
	st.Equal(expVlm, c.getVolumeById(gid), "should put volume into a goodsid, expected=%v actual=%v\n", expVlm, c.getVolumeById(gid))
}
func (st *CartRepositoryTestSuite) Test_add_different_goods_into_the_Cart_for_same_token() {
	gid1 := "g7225946"
	gid2 := "g7225947"
	token := "13911057997"
	cartRepo.AddOrderIntoCart(token, gid1, 10)
	cartRepo.AddOrderIntoCart(token, gid2, 20)

	st.Equal(1, len(cartRepo.cartInfos), "should have 1 carts for same token, but it was %v ", len(cartRepo.cartInfos))

	c := cartRepo.GetCartByToken(token)
	st.Equal(2, len(c.Items), "should be 2 Items, but it was %v ", len(c.Items))
	st.Equal(2, len(c.Pairs), "should be 2 Pairs, but it was %v ", len(c.Items))
	its := c.Items

	it := its[0]
	st.Equal(gid1, it.Gid)
	st.Equal(uint(10), it.Quantity)
	st.Equal("66.0", it.Price)
	st.Contains(it.Pic, gid1)

	it = its[1]
	st.Equal(gid2, it.Gid)
	st.Equal(uint(20), it.Quantity)
	st.Equal("99.0", it.Price)
	st.Contains(it.Pic, gid2)

}

func (st *CartRepositoryTestSuite) Test_should_seperate_carts_for_different_tokens() {
	gid := "g7225946"
	t1 := "TokenOne"
	t2 := "TokenTwo"
	cartRepo.AddOrderIntoCart(t1, gid, 10)
	cartRepo.AddOrderIntoCart(t2, gid, 20)

	length := len(cartRepo.cartInfos)
	st.Equal(2, length, "should have 2 tokens , but it was %v ", length)

	ci := cartRepo.GetCartByToken(t1)
	st.Equal(t1, ci.getToken())
	its := ci.Items
	st.Equal(1, len(its))
	it := its[0]
	st.Equal(gid, it.Gid)
	st.Equal(uint(10), it.Quantity)
	st.Equal("66.0", it.Price)
	st.Contains(it.Pic, gid)

	ci = cartRepo.GetCartByToken(t2)
	st.Equal(t2, ci.getToken())
	its = ci.Items
	st.Equal(1, len(its))
	it = its[0]
	st.Equal(gid, it.Gid)
	st.Equal(uint(20), it.Quantity)
	st.Equal("66.0", it.Price)
	st.Contains(it.Pic, gid)

}
