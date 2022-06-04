package addresses

type Address struct {
	Id            string `json:"id"`
	LinkMan       string `json:"linkMan"`
	Mobile        string `json:"mobile"`
	IsDefault     bool   `json:"isDefault"`
	ProvinceStr   string `json:"provinceStr"`
	CityStr       string `json:"cityStr"`
	AreaStr       string `json:"areaStr"`
	DetailAddress string `json:"address"`
}
