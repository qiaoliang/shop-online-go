package repository


var carts = make(map[string] uint)

func PutIntoCart(token string,goodsId uint, count uint)  {
	if _,ok:= carts[token];!ok{
		carts[token] = goodsId
		return;
	}
}