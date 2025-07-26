package banner

type BannerService struct {
	repo *BannerRepoDB
}

func NewBannerService(repo *BannerRepoDB) *BannerService {
	return &BannerService{repo: repo}
}

func (s *BannerService) FetchBanners(btype string, token string) []BannerVM {
	banners := s.repo.FindAllBanners()
	result := make([]BannerVM, len(banners))

	for i, b := range banners {
		result[i] = BannerVM{
			BusinessId: b.BusinessId,
			DateAdd:    b.DateAdd,
			Id:         b.Id,
			LinkUrl:    b.LinkUrl,
			Paixu:      b.Paixu,
			PicUrl:     b.PicUrl,
			Remark:     b.Remark,
			Status:     b.Status,
			StatusStr:  b.StatusStr,
			Title:      b.Title,
			Catelogue:  b.Type,
			UserId:     b.UserId,
		}
	}

	return result
}
