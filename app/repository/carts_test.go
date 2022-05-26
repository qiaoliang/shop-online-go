package repository

import (
	_ "fmt"
	"reflect"
	"testing"
)

func Test_CartRepo_is_empty(t *testing.T) {
	cr := GetCartsInstance()
	if cr.size() != 0 {
		t.Fatalf("carts should be empty at the beginning , expected=%v actual=%v\n", 0, cr.size())
	}
}

func Test_add_new_token_in_cartrepo(t *testing.T) {
	cr := GetCartsInstance()
	if cr.size() != 0 {
		t.Fatalf("carts should be empty at the beginning , expected=%v actual=%v\n", 0, cr.size())
	}
	cr.AddOrderIntoCart("IamToken", 1, 1)
	if cr.size() != 1 {
		t.Fatalf("should add one token , expected=%v actual=%v\n", 1, cr.size())
	}
	ct := cr.getCartBy("IamToken")
	if ct == nil {
		t.Fatalf("should get Cart by a given token , expected=%v actual=%v\n", "not nil", ct)
	}
	if ct.getToken() != "IamToken" {
		t.Fatalf("should get correct Token after adding , expected=%v actual=%v\n", "IamToken", ct.getToken())
	}

	expGoods := map[uint]uint{1: 1}
	ok := reflect.DeepEqual(expGoods, ct.getGoods())
	if !ok {
		t.Fatalf("should get correct goods after adding a new one, expected=%v actual=%v\n", expGoods, ct.getGoods())
	}
}
