package ad

import (
	"bookstore/app/configs"
)

var bs *BannerService

type BannerService struct {
	bl Banners
}

func GetBannerService() *BannerService {
	if bs == nil {
		bs = newBannerService(configs.Cfg.Persistence)
	}
	return bs
}
func newBannerService(persistance bool) *BannerService {
	return &BannerService{make(Banners, 0)}
}
func (s *BannerService) FetchBanners(btype string, token string) Banners {
	if btype == "indexBanner" {
		return s.initBannerData()
	}
	//TODO: 需要实现根据 Banner 的类型与用户token ，返回不同的 Banner 广告
	return s.bl
}

func (s *BannerService) initBannerData() Banners {
	ban1 := &Banner{
		0,
		"2022-05-05 11:26:09",
		222083,
		"https://gitee.com/sagittatius",
		0,
		configs.Cfg.BannerPicPrefix() + "b0001.jpeg",
		"跳转gitee sagittatius",
		0,
		"any",
		"any",
		"any",
		1605,
	}

	ban2 := &Banner{
		1,
		"2022-05-05 11:26:09",
		222084,
		"https://gitee.com/sagittatius",
		0,
		configs.Cfg.BannerPicPrefix() + "b0002.jpeg",
		"跳转gitee sagittatius",
		0,
		"any",
		"any",
		"any",
		1606,
	}
	s.bl = append(s.bl, *ban1)
	s.bl = append(s.bl, *ban2)
	return s.bl
}
