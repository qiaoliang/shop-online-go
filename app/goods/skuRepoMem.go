package goods

type SkuRepoMem struct {
}

func (s SkuRepoMem) First(skuid string) *SKU {
	//TODO: Not implementated yet
	return nil
}
func (s SkuRepoMem) FindWithCarouselPics(skuid string) *SKU {
	//TODO: Not implementated yet
	return nil
}
func (s SkuRepoMem) Create(sku SKU) error {
	//TODO: Not implementated yet
	return nil
}
func (s SkuRepoMem) Update(skuid string, sku SKU) error {
	//TODO: Not implementated yet
	return nil
}

func (s SkuRepoMem) FindAll() []SKU {
	//TODO: Not implementated yet
	return nil
}
