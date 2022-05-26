package repository

type Cart struct {
	token string
	goods map[uint]uint
}

func NewCart(token string) *Cart {
	return &Cart{token, make(map[uint]uint, 0)}
}
func (c Cart) getVolumefor(goodsid uint) uint {
	return c.goods[goodsid]
}

func (c Cart) addItem(goodsid uint, volume uint) {
	if c.goods == nil {
		c.goods = make(map[uint]uint)
	}
	c.goods[goodsid] += volume
}

func (c Cart) getToken() string {
	return c.token
}
func (c Cart) getVolumeById(id uint) uint {
	return c.goods[id]
}

func (c Cart) getGoods() map[uint]uint {
	return c.goods
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

func (cr CartRepo) size() int {
	return len(cr.carts)
}

func (cs CartRepo) AddOrderIntoCart(token string, goodsId uint, volume uint) {
	if _, ok := cs.carts[token]; !ok {
		cs.carts[token] = NewCart(token)
	}
	cs.carts[token].addItem(goodsId, volume)
}
func (cs CartRepo) getCartBy(token string) *Cart {
	if _, ok := cs.carts[token]; !ok {
		return nil
	}
	return cs.carts[token]
}
