package ad

import (
	"github.com/example/project/app/configs"
)

var bs *BannerService

type BannerService struct {
	defaultlist Banners
}

func GetBannerService() *BannerService {
	if bs == nil {
		bs = newBannerService(configs.Cfg.Persistence)
	}
	return bs
}
func newBannerService(persistance bool) *BannerService {
	default_list := defaultList()
	return &BannerService{default_list}
}
func (s *BannerService) FetchBanners(btype string, token string) Banners {
	if btype == "indexBanner" {
		return s.defaultlist
	}
	//TODO: 根据 Banner Type  和 token ，返回不同的 Banner 广告
	return nil
}

func defaultList() Banners {
	ret := make(Banners, 0)
	ban1 := &BannerVM{
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

	ban2 := &BannerVM{
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
	ret = append(ret, *ban1)
	ret = append(ret, *ban2)
	return ret
}
