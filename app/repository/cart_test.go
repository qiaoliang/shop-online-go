package repository

import (
	_ "fmt"
	"reflect"
	"testing"
)

func TestCart(t *testing.T) {
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
	expVlm = uint(15)
	c.addItem(1, 5)
	if expVlm != c.getVolumeById(1) {
		t.Fatalf("should add more volume for a given goodsid, expected=%v actual=%v\n", expVlm, c.getVolumefor(1))
	}

	expGoods := map[uint]uint{1: 15}
	ok := reflect.DeepEqual(expGoods, c.getGoods())
	if !ok {
		t.Fatalf("should get goods for a goodsid, expected=%v actual=%v\n", expGoods, c.getGoods())
	}
}
