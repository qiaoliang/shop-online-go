package goods

import (
	"errors"
	"sort"
)

type SkuRepoMem struct {
	items map[string]SKU
}

func (s SkuRepoMem) First(skuid string) *SKU {
	sku, err := s.items[skuid]
	if err {
		return nil
	}
	return &sku
}
func (s SkuRepoMem) FindWithCarouselPics(skuid string) *SKU {
	return s.First(skuid)
}
func (s SkuRepoMem) Create(sku SKU) error {

	return nil
}
func (s SkuRepoMem) Update(skuid string, sku *SKU) error {
	found := s.First(skuid)
	if found == nil {
		return errors.New("can not find the target sku by skuid: " + skuid)
	}
	s.items[skuid] = *sku
	return nil
}

func (s SkuRepoMem) FindAll() []SKU {
	if len(s.items) != 0 {
		return orderBySkuID(s.items)
	}
	// add the first catalouge
	item1 := createSKU(0, 0, "g7225946", "持续交付1.0", 110, "册", "1", "DevOps 的第一本书", ONSAIL, "66.0", "99.0", BOTH, "1")
	item2 := createSKU(1, 0, "g7225947", "持续交付2.0", 200, "册", "1", "另一本DevOps的经典书。", ONSAIL, "99.0", "129.0", BOTH, "1")
	item3 := createSKU(2, 0, "g7225948", "DevOps实战指南", 10, "册", "1", "DevOps 黄皮书。", ONSAIL, "55.0", "85.0", BOTH, "1")
	item4 := createSKU(3, 0, "g7225949", "谷歌软件工程", 20, "册", "1", "解密硅谷头部互联网企业 如何打造软件工程文化。", ONSAIL, "77.0", "107.0", BOTH, "1")
	s.items[item1.SkuId] = item1
	s.items[item2.SkuId] = item2
	s.items[item3.SkuId] = item3
	s.items[item4.SkuId] = item4

	// add the second catalouge

	item5 := createSKU(11, 1, "g1872110", "驾驭大数据", 20, "册", "0", "另一本DevOps的经典书。", ONSAIL, "99.0", "129.0", BOTH, "1")
	item6 := createSKU(12, 1, "g1872111", "数据分析变革", 10, "册", "0", "一本DevOps的经典书。", ONSAIL, "66.0", "99.0", BOTH, "1")
	item7 := createSKU(13, 1, "g1872112", "大数据测试技术与实践", 2, "册", "0", "DevOps 黄皮书。", ONSAIL, "55.0", "85.0", BOTH, "1")
	item8 := createSKU(14, 1, "g1872113", "图解Spark 大数据快速分析实战", 5, "册", "0", "解密硅谷头部互联网企业 如何打造软件工程文化。", ONSAIL, "77.0", "107.0", BOTH, "1")

	s.items[item5.SkuId] = item5
	s.items[item6.SkuId] = item6
	s.items[item7.SkuId] = item7
	s.items[item8.SkuId] = item8
	return orderBySkuID(s.items)
}
func orderBySkuID(m map[string]SKU) []SKU {
	v := make([]SKU, 0, len(m))
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		v = append(v, m[key])
	}
	return v
}

func (s SkuRepoMem) Delete(sku SKU) error {
	delete(s.items, sku.SkuId)
	return nil
}

func createSKU(
	id uint, cateId uint,
	gid string, gName string, gStock uint,
	unit string, logistics string,
	desc string, status SalingStatus,
	minPrice string, origPrice string,
	afterSale AfterSaleType, recommandStatus string) SKU {

	detail := SKU{
		SkuId:           gid,
		Name:            gName,
		CategoryId:      cateId,
		RecommendStatus: recommandStatus,
		PicStr:          gid + ".jpeg",
		Unit:            unit,
		Stock:           gStock,
		MinPrice:        minPrice,
		OriginalPrice:   origPrice,
		Logistics:       logistics,
		Content:         desc,
		Status:          status,
		Aftersale:       afterSale,
		SkuCarouPictures: []SkuCarouPicture{
			{
				Id:     0,
				SkuId:  gid,
				PicStr: "-01.jpeg",
			},
			{
				Id:     0,
				SkuId:  gid,
				PicStr: "-02.jpeg",
			},
		},
	}
	return detail
}
