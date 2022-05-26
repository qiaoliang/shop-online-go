package repository

import (
  	_"fmt"
	"testing"
)

func Test_Cart_is_empty(t *testing.T){


	if 0 != len(carts) {
		t.Fatalf("初始时，所有购物车应该是空的 ，期望值=%v 实际结果=%v\n",0,len(carts))
	}

	//成功 输出日志
	t.Logf("初始时，所有购物车是空的")
}