package cart

import (
	_ "fmt"
	"testing"
)

func Test_add_one_goods_into_an_empty_Cart(t *testing.T) {
	cartRepo = nil
	cartRepo = GetCartsInstance()
	expected := "IamTestToken"
	c := cartRepo.AddOrderIntoCart("IamTestToken", 1, 10)
	if expected != c.getToken() {
		t.Fatalf("should get token from a cart, expected=%v actual=%v\n", expected, c.getToken())
	}
	expVlm := uint(10)
	if expVlm != c.getVolumeById(1) {
		t.Fatalf("should put volume into a goodsid, expected=%v actual=%v\n", expVlm, c.getVolumeById(1))
	}
}
func Test_add_more_volume_into_a_Cart_with_same_goodsid(t *testing.T) {
	cartRepo = nil
	cartRepo = GetCartsInstance()
	cartRepo.AddOrderIntoCart("sameGoodsId", 1, 10)

	c := cartRepo.AddOrderIntoCart("sameGoodsId", 1, 13)

	lens := len(c.Items)
	if lens != 1 {
		t.Fatalf("should be 1, but it was %v ", lens)
	}
	if c.getVolumeById(1) != 13 {
		t.Fatalf("should be 13, but it was %v", c.getVolumeById(1))
	}
}
func Test_add_another_goods_into_a_Cart_with_one_goods(t *testing.T) {
	cartRepo = nil
	cartRepo = GetCartsInstance()
	cartRepo.AddOrderIntoCart("firstGoodsId", 1, 10)
	c := cartRepo.AddOrderIntoCart("secondGoodsId", 2, 20)
	length := len(cartRepo.cartInfos)
	if length != 2 {
		t.Fatalf("should have 2 tokens , but it was %v ", length)
	}
	if c.getVolumeById(2) != 20 {
		t.Fatalf("should be 20 for secondGoodsId, but it was %v", c.getVolumeById(2))
	}
}
