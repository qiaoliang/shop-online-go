package banner

type BannerVM struct {
	BusinessId uint   `json:"businessId"`
	DateAdd    string `json:"dateAdd"`
	Id         uint   `json:"id"`
	LinkUrl    string `json:"linkUrl"`
	Paixu      uint   `json:"paixu"`
	PicUrl     string `json:"picUrl"`
	Remark     string `json:"remark"`
	Status     uint   `json:"status"`
	StatusStr  string `json:"statusStr"`
	Title      string `json:"title"`
	Catelogue  string `json:"type"`
	UserId     uint   `json:"userId"`
}
