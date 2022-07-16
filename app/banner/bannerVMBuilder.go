package ad

import "github.com/example/project/app/configs"

type BannerVMBuilder struct {
	b BannerVM
}

var pic_prefix = configs.Cfg.BannerPicPrefix()

func NewBannerBuilder() BannerVMBuilder {
	return BannerVMBuilder{
		b: BannerVM{
			BusinessId: 0,
			DateAdd:    "2022-05-05 11:26:09",
			Id:         222083,
			LinkUrl:    "https://gitee.com/sagittatius",
			Paixu:      0,
			PicUrl:     pic_prefix + "b0000.jpeg",
			Remark:     "跳转 gitee sagittatius",
			Status:     0,
			StatusStr:  "any",
			Title:      "any",
			Catelogue:  "any",
			UserId:     1605,
		},
	}
}

func (s BannerVMBuilder) bizId(bizId uint) BannerVMBuilder {
	s.b.BusinessId = bizId
	return s
}
func (s BannerVMBuilder) DateAdd(date string) BannerVMBuilder {
	s.b.DateAdd = date
	return s
}
func (s BannerVMBuilder) Id(Id uint) BannerVMBuilder {
	s.b.Id = Id
	return s
}
func (s BannerVMBuilder) LinkUrl(url string) BannerVMBuilder {
	s.b.LinkUrl = url
	return s
}
func (s BannerVMBuilder) Paixu(paixu uint) BannerVMBuilder {
	s.b.Paixu = paixu
	return s
}

func (s BannerVMBuilder) PicUrl(picUrl string) BannerVMBuilder {
	s.b.PicUrl = pic_prefix + picUrl
	return s
}

func (s BannerVMBuilder) Status(status uint) BannerVMBuilder {
	s.b.Status = status
	return s
}

func (s BannerVMBuilder) Title(t string) BannerVMBuilder {
	s.b.Title = t
	return s
}
func (s BannerVMBuilder) Catelogue(c string) BannerVMBuilder {
	s.b.Catelogue = c
	return s
}
func (s BannerVMBuilder) UserId(uid uint) BannerVMBuilder {
	s.b.UserId = uid
	return s
}
func (s BannerVMBuilder) build() BannerVM {
	return s.b
}
