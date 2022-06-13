package cart

type UserCartItemBuilder struct {
	item UserCartItem
}

func NewUCIBuilder() UserCartItemBuilder {
	return UserCartItemBuilder{UserCartItem{
		0,
		"token-13900007997", //Token
		"7225946",           //Gid
		"7225946.jpeg",      //Pic
		0,                   //Status
		"持续交付1.0",           //Name
		"sku1,sku2",         //SkuStr
		"66.0",              //Price
		110,                 //Quantity
		"1",                 //Selected
		"OptionValueName",   //OptionValueName
	}}
}
func (b UserCartItemBuilder) id(id uint) UserCartItemBuilder {
	b.item.Id = id
	return b
}

func (b UserCartItemBuilder) token(token string) UserCartItemBuilder {
	b.item.Token = token
	return b
}

func (b UserCartItemBuilder) gid(gid string) UserCartItemBuilder {
	b.item.Gid = gid
	return b
}

func (b UserCartItemBuilder) name(name string) UserCartItemBuilder {
	b.item.Name = name
	return b
}
func (b UserCartItemBuilder) optionValueName(value string) UserCartItemBuilder {
	b.item.OptionValueName = value
	return b
}
func (b UserCartItemBuilder) status(status uint) UserCartItemBuilder {
	b.item.Status = status
	return b
}
func (b UserCartItemBuilder) sku(skuStrs string) UserCartItemBuilder {
	b.item.SkuStrs = skuStrs
	return b
}
func (b UserCartItemBuilder) price(price string) UserCartItemBuilder {
	b.item.Price = price
	return b
}
func (b UserCartItemBuilder) quantity(quantity uint) UserCartItemBuilder {
	b.item.Quantity = quantity
	return b
}
func (b UserCartItemBuilder) selected(Selected string) UserCartItemBuilder {
	b.item.Selected = Selected
	return b
}
func (b UserCartItemBuilder) pic(pic string) UserCartItemBuilder {
	b.item.Pic = pic
	return b
}

func (b UserCartItemBuilder) build() UserCartItem {
	return b.item
}
