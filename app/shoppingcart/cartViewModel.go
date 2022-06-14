package cart

import (
	"bookstore/app/configs"
	"bookstore/app/goods"
)

type CartInfoVM struct {
	Token  string       `json:"token"`
	RedDot uint         `json:"number"` //等于用户购物车中SKU的品类个数（京东购物车的逻辑）
	Items  []CartItemVM `json:"items"`
	Pairs  []ItemPairVM `json:"goods"`
}

type CartItemVM struct {
	Gid             string   `json:"key"`
	Pic             string   `json:"pic"`
	Status          uint     `json:"status"` // === 1 【失效】
	Name            string   `json:"name"`
	Sku             []string `json:"sku"`
	Price           string   `json:"price"`
	Quantity        uint     `json:"number"`
	Selected        string   `json:"selected"`
	OptionValueName string   `json:"optionValueName"`
}

func (civm *CartItemVM) RetrivePicStr() string {
	l := len(configs.Cfg.GoodsPicPrefix())
	return civm.Pic[l:]
}

func (civm *CartItemVM) AddMore(quantity uint) {
	civm.Quantity += quantity
}

type ItemPairVM struct {
	GoodsId string `json:"goodsId"`
	Volume  uint   `json:"number"`
}

func (ip *ItemPairVM) AddMore(quantity uint) {
	ip.Volume = ip.Volume + quantity
}

func (ci *CartInfoVM) caculateRedDot() {
	ci.RedDot = uint(len(ci.Items))
}

func (ci *CartInfoVM) getToken() string {
	return ci.Token
}
func (ci *CartInfoVM) FindBy(skuid string) (*CartItemVM, *ItemPairVM) {
	return ci.findItemByGid(skuid), ci.findPairByGid(skuid)

}

func (ci *CartInfoVM) findItemByGid(gid string) *CartItemVM {
	for i := range ci.Items {
		it := &ci.Items[i]
		if it.Gid == gid {
			return it
		}
	}
	return nil
}
func (ci *CartInfoVM) findPairByGid(gid string) *ItemPairVM {
	for i := range ci.Pairs {
		it := &ci.Pairs[i]
		if it.GoodsId == gid {
			return it
		}
	}
	return nil
}

func (ci *CartInfoVM) AddMore(prod *goods.GoodsDetail, quantity uint) {
	item := ci.findItemByGid(prod.Gid)
	if item != nil {
		updatedQuantity := item.Quantity + quantity
		item.Quantity = updatedQuantity
		ip := ci.findPairByGid(prod.Gid)
		ip.Volume = item.Quantity
		return
	}
	item = ci.createCartItem(prod, quantity)
	ip := ItemPairVM{prod.Gid, quantity}
	ci.Items = append(ci.Items, *item)
	ci.Pairs = append(ci.Pairs, ip)
}

func (ci *CartInfoVM) Modify(gid string, quantity uint) *CartItemVM {
	item := ci.findItemByGid(gid)
	if item != nil {
		item.Quantity = quantity
		ip := ci.findPairByGid(gid)
		ip.Volume = item.Quantity
		return item
	}
	return nil
}

func (c *CartInfoVM) getVolumeById(gid string) uint {
	if len(c.Pairs) == 0 {
		return uint(0)
	}

	for _, item := range c.Pairs {
		v := item
		if v.GoodsId == gid {
			return item.Volume
		}
	}
	return uint(0)
}

func (ci *CartInfoVM) createCartItem(prod *goods.GoodsDetail, quantity uint) *CartItemVM {
	sku := []string{"sku1", "sku3"}
	selected := "1"
	optionValue := "optionValueName"

	item := CartItemVM{prod.Gid,
		configs.Cfg.GoodsPicPrefix() + prod.Gid + ".jpeg",
		0,
		prod.Name,
		sku,
		prod.MinPrice,
		quantity,
		selected,
		optionValue,
	}
	return &item
}
