package repository

import (
	_ "fmt"
	"testing"
)

func Test_add_one_goods_into_an_empty_Cart(t *testing.T) {
	expected := "IamToken"
	c := NewCart("IamToken")
	if expected != c.getToken() {
		t.Fatalf("should get token from a cart, expected=%v actual=%v\n", expected, c.getToken())
	}
	expVlm := uint(10)
	c.addItem(1, 10)
	if expVlm != c.getVolumeById(1) {
		t.Fatalf("should put volume into a goodsid, expected=%v actual=%v\n", expVlm, c.getVolumefor(1))
	}
}
func Test_add_more_volume_into_a_Cart_with_same_goodsid(t *testing.T) {

	c := &Cart{"one", make([]ItemPair, 0)}
	c.Pairs = append(c.Pairs, ItemPair{uint(1), uint(3)})

	c.addItem(1, 10)

	lens := len(c.Pairs)
	if lens != 1 {
		t.Fatalf("should be 1, but it was %v ", lens)
	}
	if c.getVolumeById(1) != 13 {
		t.Fatalf("should be 13, but it was %v", c.getVolumefor(1))
	}
}
func Test_add_another_goods_into_a_Cart_with_one_goods(t *testing.T) {

	c := &Cart{"one", make([]ItemPair, 0)}
	c.Pairs = append(c.Pairs, ItemPair{uint(1), uint(3)})
	c.addItem(2, 20)

	length := len(c.Pairs)
	if length != 2 {
		t.Fatalf("should be 2, but it was %v ", length)
	}
	if c.getVolumeById(2) != 20 {
		t.Fatalf("should be 20, but it was %v", c.getVolumefor(2))
	}
}
