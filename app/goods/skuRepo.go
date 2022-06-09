package goods

import (
	"bookstore/app/configs"
	"sync"
)

var lockSku = &sync.Mutex{}
var skuRepo SkuRepoIf

func GetSkuRepoDB() SkuRepoIf {
	lockSku.Lock()
	defer lockSku.Unlock()
	if skuRepo == nil {
		if configs.Cfg.Persistence {
			skuRepo = &SkuRepoDB{}
		}
	}

	return skuRepo
}

type SkuRepoIf interface {
	Find(skuid string) *SKU
}
type SkuRepoDB struct {
}

func (s SkuRepoDB) Find(skuid string) *SKU {
	return nil
}
