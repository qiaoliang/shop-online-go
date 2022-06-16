package cart

import (
	"errors"
)

type CartRepoMem struct {
	carts map[string]*UserCart
}

var cartRepo CartRepoIf

func (cs *CartRepoMem) FindUserCartItemsBy(token string) []UserCartItem {
	if _, ok := cs.carts[token]; !ok {
		return nil
	}
	return cs.carts[token].Items
}

func (cs *CartRepoMem) DeleteUserCartItemsBy(token string) error {
	delete(cs.carts, token)
	return nil
}
func (cs *CartRepoMem) GetUserCartItem(uci UserCartItem) *UserCartItem {
	cart, ok := cs.carts[uci.Token]
	if !ok {
		return nil
	}
	item, _ := cart.findUserCartItem(uci.SkuId)
	return item
}
func (cs *CartRepoMem) UpdateUserCartItem(uci *UserCartItem) error {
	ret := cs.GetUserCartItem(*uci)
	if ret == nil {
		return errors.New("can not find CartItem")
	}
	ret.Quantity = uci.Quantity
	return nil
}
func (cs *CartRepoMem) DeleteUserCartItem(uci UserCartItem) error {
	cart, ok := cs.carts[uci.Token]
	if !ok {
		return errors.New("can not find userCartItem")
	}
	_, index := cart.findUserCartItem(uci.SkuId)
	if index == -1 {
		return errors.New("can not find userCartItem")
	}
	cart.Items = append(cart.Items[:index], cart.Items[index+1:]...)
	return nil

}
func (cs *CartRepoMem) SaveUserCartItem(uci UserCartItem) error {
	cart, ok := cs.carts[uci.Token]
	if !ok {
		cart := UserCart{
			Token: uci.Token,
			Items: []UserCartItem{uci},
		}
		cs.carts[uci.Token] = &cart
		return nil
	}
	item, index := cart.findUserCartItem(uci.SkuId)
	if item != nil {
		cart.Items[index] = uci
		return nil
	}
	cart.Items = append(cart.Items, uci)
	return nil

}
