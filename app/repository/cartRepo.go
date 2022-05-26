package repository

import "fmt"

type Cart struct {
	Token string     `json:"token"`
	Pairs []ItemPair `json:"goods"`
}
type ItemPair struct {
	GoodsId uint `json:"goodsId"`
	Volume  uint `json:"number"`
}

func NewCart(token string) *Cart {
	return &Cart{token, make([]ItemPair, 0)}
}
func (c *Cart) getVolumefor(goodsid uint) uint {
	if len(c.Pairs) == 0 {
		return uint(0)
	}
	for _, item := range c.Pairs {
		if item.GoodsId == goodsid {
			return item.Volume
		}
	}
	return 0
}

func (c *Cart) addItem(id uint, count uint) {
	for i, item := range c.Pairs {
		if c.Pairs[i].GoodsId == id {
			c.Pairs[i].Volume = uint(item.Volume + count)
			return
		}
	}
	c.Pairs = append(c.Pairs, ItemPair{id, count})

}

func (c *Cart) getToken() string {
	return c.Token
}
func (c *Cart) getVolumeById(id uint) uint {
	if len(c.Pairs) == 0 {
		return uint(0)
	}

	for _, item := range c.Pairs {
		v := item
		fmt.Println(v.GoodsId)
		if v.GoodsId == id {
			return item.Volume
		}
	}
	return uint(0)
}

type CartRepo struct {
	carts map[string]*Cart
}

var cartRepo *CartRepo

func GetCartsInstance() *CartRepo {
	if cartRepo == nil {
		cartRepo = &CartRepo{make(map[string]*Cart, 0)}
	}
	return cartRepo
}

func (cr *CartRepo) size() int {
	return len(cr.carts)
}

func (cs *CartRepo) AddOrderIntoCart(token string, goodsId uint, volume uint) *Cart {
	if _, ok := cs.carts[token]; !ok {
		cs.carts[token] = NewCart(token)
	}
	cs.carts[token].addItem(goodsId, volume)
	ct := cs.carts[token]
	return ct
}
func (cs *CartRepo) GetCartBy(token string) *Cart {
	if _, ok := cs.carts[token]; !ok {
		return nil
	}
	return cs.carts[token]
}
