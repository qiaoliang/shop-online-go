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

func Test_add_new_token_into_empty_cartrepo(t *testing.T) {
	cr := GetCartsInstance()
	if cr.size() != 0 {
		t.Fatalf("carts should be empty at the beginning , expected=%v actual=%v\n", 0, cr.size())
	}
	cr.AddOrderIntoCart("IamToken", 1, 1)
	if cr.size() != 1 {
		t.Fatalf("should add one token , expected=%v actual=%v\n", 1, cr.size())
	}
	cart := cr.GetCartBy("IamToken")

	if cart == nil {
		t.Fatalf("should get Cart by a given token , expected=%v actual=%v\n", "not nil", cart)
	}
	if cart.getToken() != "IamToken" {
		t.Fatalf("should get correct Token after adding , expected=%v actual=%v\n", "IamToken", cart.getToken())
	}
	if len(cart.Pairs) == 0 {
		t.Fatalf("should get a ItemPair at least after adding, bu	t it is 0\n")
	}

	exp := ItemPair{1, 1}
	if !reflect.DeepEqual(exp, cart.Pairs[0]) {
		t.Fatalf(" expected=%v actual=%v\n", exp, cart.Pairs[0])
	}
}

func Test_add_another_goods_for_same_token(t *testing.T) {
	cr := GetCartsInstance()
	ct := &Cart{"one", make([]ItemPair, 0)}
	ct.Pairs = append(ct.Pairs, ItemPair{uint(1), uint(3)})
	cr.carts["one"] = ct

	cr.AddOrderIntoCart("one", uint(2), uint(4))
	if cr.size() != 1 {
		t.Fatalf(" expected= 1 actual=%v\n", cr.size())
	}

	ct = cr.GetCartBy("one")
	result := ct.getVolumeById(uint(2))
	if result != 4 {
		t.Fatalf(" expected= 2 actual=%v\n", result)
	}

}
