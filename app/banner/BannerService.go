package ad

type BannerService struct {
	repo *BannerRepoDB
}

func NewBannerService(repo *BannerRepoDB) *BannerService {
	return &BannerService{repo: repo}
}

func (s *BannerService) FetchBanners(btype string, token string) []Banner {
	return s.repo.FindAllBanners()
}
