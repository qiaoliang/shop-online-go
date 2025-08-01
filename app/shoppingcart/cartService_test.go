package cart

import (
	"strings"
	"testing"

	"bookstore/app/configs"
	"bookstore/app/goods"
	"bookstore/app/testutils"
	"bookstore/app/utils"

	"github.com/stretchr/testify/suite"
)

type CartServiceTestSuite struct {
	testutils.SupperSuite
	cs *CartService
}

func TestSkuRepoDBTestSuite(t *testing.T) {
	suite.Run(t, new(CartServiceTestSuite))
}

func (s *CartServiceTestSuite) BeforeTest(suiteName, testName string) {}

func (s *CartServiceTestSuite) AfterTest(suiteName, testName string) {}

func (s *CartServiceTestSuite) SetupSuite() {
	s.SupperSuite.SetupSuite()
	db := configs.Cfg.DBConnection()
	goodsRepo := goods.NewSkuRepoDB(db)
	repo := NewCartRepoDB(db)
	s.cs = NewCartService(goodsRepo, repo)
}
func (s *CartServiceTestSuite) TeardownSuite() {
	s.SupperSuite.TeardownSuite()
	s.cs = nil
	cartService = nil
}
func (s *CartServiceTestSuite) SetupTest() {}

func (s *CartServiceTestSuite) Test_GetPersistance() {
	db := configs.Cfg.DBConnection()
	goodsRepo := goods.NewSkuRepoDB(db)
	repo := NewCartRepoDB(db)
	cs := NewCartService(goodsRepo, repo)
	s.NotNil(cs.cr)
}

const (
	ANY_NUMBER      = 8888
	ANY_TOKEN       = "ANY"
	UNEXISTED_SKU   = "unexisted_SKU"
	EXISTED_SKU_ONE = "g7225946"
	EXISTED_SKU_TWO = "g7225947"
)

func (s *CartServiceTestSuite) Test_CreateCartInfoFor() {

	token := " Test_CreateCartInfoFor" + utils.RandomImpl{}.GenStr()
	exp_skuID := EXISTED_SKU_ONE
	quantity := uint(10)
	expGd, expIf := s.generateExp(exp_skuID, quantity, token)

	ci := s.cs.CreateCartItemFor(token, expGd, quantity)

	s.NotNil(ci)
	s.EqualValues(expIf, ci)
}

func (s *CartServiceTestSuite) Test_ModifyQuantityOfGoodsInCate() {

	token := "ModifyQuantityOfGoodsInCate" + utils.RandomImpl{}.GenStr()
	exp_skuID := EXISTED_SKU_ONE
	orgQuan := uint(10)
	expGd, expIf := s.generateExp(exp_skuID, orgQuan, token)
	s.cs.CreateCartItemFor(token, expGd, orgQuan)

	updatedQuan := uint(30)
	expIf.Items[0].Quantity = updatedQuan
	expIf.Pairs[0].Volume = updatedQuan

	ci := s.cs.ModifyQuantityOfGoodsInCate(token, exp_skuID, updatedQuan)

	s.NotNil(ci)
	s.EqualValues(expIf, ci)

}

func (s *CartServiceTestSuite) Test_Can_not_Same_unexisted_SKU_In_Cart() {
	s.Nil(s.cs.PutItemsInCart(ANY_TOKEN, UNEXISTED_SKU, ANY_NUMBER))
}
func (s *CartServiceTestSuite) Test_Put_a_new_cart_Item_Into_Cart() {
	token := "Test_Put_a_new_cart_Item_Into_Cart" + utils.RandomImpl{}.GenStr()
	exp_skuID := EXISTED_SKU_ONE
	number := uint(10)
	_, expIf := s.generateExp(exp_skuID, number, token)
	ci := s.cs.PutItemsInCart(token, exp_skuID, number)
	s.NotNil(ci)
	s.EqualValues(expIf, ci)
}
func (s *CartServiceTestSuite) Test_Should_add_more_volume_for_same_cart_Item_in_Cart() {

	token := "Should_add_more_volume_for_same_cart_Item" + utils.RandomImpl{}.GenStr()
	exp_skuID := EXISTED_SKU_ONE
	number := uint(10)
	_, expIf := s.generateExp(exp_skuID, number, token)
	s.cs.PutItemsInCart(token, exp_skuID, number)

	more := uint(10)
	expIf.Items[0].Quantity = number + more
	expIf.Pairs[0].Volume = number + more
	ci := s.cs.PutItemsInCart(token, exp_skuID, more)
	s.NotNil(ci)
	s.EqualValues(expIf, ci)
}
func (s *CartServiceTestSuite) Test_add_multiple_sku_for_same_cart() {

	token := "Test_add_multiple_sku_for_same_cart" + utils.RandomImpl{}.GenStr()
	number1 := uint(10)

	s.cs.PutItemsInCart(token, EXISTED_SKU_ONE, number1)

	number2 := uint(10)
	_, expItem, expIp := s.generateExpItem(EXISTED_SKU_TWO, number2, token)

	ci := s.cs.PutItemsInCart(token, EXISTED_SKU_TWO, number2)
	s.NotNil(ci)
	s.EqualValues(expItem, ci.findItemByGid(EXISTED_SKU_TWO))
	s.EqualValues(expIp, ci.findPairByGid(EXISTED_SKU_TWO))

}
func (s *CartServiceTestSuite) generateExp(sku_id string, quantity uint, token string) (*goods.SKU, *CartInfoVM) {
	gd := s.cs.goodsRepo.First(sku_id)
	expItem := NewCartItemVMBuilder(gd).quantity(quantity).picStr(configs.Cfg.GoodsPicPrefix() + gd.PicStr).build()
	expIp := NewItemPairVMBuilder().gid(gd.SkuId).volume(quantity).build()
	expIf := NewCartInfoVMBuilder().token(token).addItem(expItem).addIpair(expIp).build()
	return gd, expIf
}

func (s *CartServiceTestSuite) generateExpItem(sku_id string, quantity uint, token string) (*goods.SKU, *CartItemVM, *ItemPairVM) {
	gd := s.cs.goodsRepo.First(sku_id)
	expItem := NewCartItemVMBuilder(gd).quantity(quantity).picStr(configs.Cfg.GoodsPicPrefix() + gd.PicStr).build()
	expIp := NewItemPairVMBuilder().gid(gd.SkuId).volume(quantity).build()
	return gd, expItem, expIp
}

type CartInfoVMBuilder struct {
	civm *CartInfoVM
}

func NewCartInfoVMBuilder() *CartInfoVMBuilder {
	return &CartInfoVMBuilder{
		civm: &CartInfoVM{
			Token:  "",
			RedDot: 0,
			Items:  []CartItemVM{},
			Pairs:  []ItemPairVM{},
		},
	}
}
func (s *CartInfoVMBuilder) token(token string) *CartInfoVMBuilder {
	s.civm.Token = token
	return s
}
func (s *CartInfoVMBuilder) addItem(it *CartItemVM) *CartInfoVMBuilder {
	s.civm.Items = append(s.civm.Items, *it)
	return s
}
func (s *CartInfoVMBuilder) addIpair(it *ItemPairVM) *CartInfoVMBuilder {
	s.civm.Pairs = append(s.civm.Pairs, *it)
	return s
}

func (s *CartInfoVMBuilder) build() *CartInfoVM {
	s.civm.RedDot = uint(len(s.civm.Items))
	return s.civm
}

type CartItemVMBuilder struct {
	ci *CartItemVM
}

func NewCartItemVMBuilder(gd *goods.SKU) *CartItemVMBuilder {
	return &CartItemVMBuilder{
		&CartItemVM{
			Gid:             gd.SkuId,
			Pic:             gd.PicStr,
			Status:          0,
			Name:            gd.Name,
			Sku:             []string{"sku1", "sku3"},
			Price:           gd.MinPrice,
			Quantity:        0,
			Selected:        "1",
			OptionValueName: "OptionValueName",
		},
	}
}
func (s *CartItemVMBuilder) gid(gid string) *CartItemVMBuilder {
	s.ci.Gid = gid
	return s
}
func (s *CartItemVMBuilder) picStr(picStr string) *CartItemVMBuilder {
	s.ci.Pic = picStr
	return s
}
func (s *CartItemVMBuilder) selected(sel string) *CartItemVMBuilder {
	s.ci.Selected = sel
	return s
}

func (s *CartItemVMBuilder) sku(skus string) *CartItemVMBuilder {
	s.ci.Sku = strings.Split(skus, ",")
	return s
}
func (s *CartItemVMBuilder) quantity(number uint) *CartItemVMBuilder {
	s.ci.Quantity = number
	return s
}
func (s *CartItemVMBuilder) build() *CartItemVM {
	return s.ci
}

type ItemPairVMBuilder struct {
	ip *ItemPairVM
}

func NewItemPairVMBuilder() *ItemPairVMBuilder {
	return &ItemPairVMBuilder{
		&ItemPairVM{
			GoodsId: "undefined",
			Volume:  0,
		},
	}
}
func (s *ItemPairVMBuilder) volume(volume uint) *ItemPairVMBuilder {
	s.ip.Volume = volume
	return s
}
func (s *ItemPairVMBuilder) gid(gid string) *ItemPairVMBuilder {
	s.ip.GoodsId = gid
	return s
}
func (s *ItemPairVMBuilder) build() *ItemPairVM {
	return s.ip
}
