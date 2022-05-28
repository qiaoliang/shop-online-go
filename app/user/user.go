package user

type User struct {
	Id        string `json:"userId"`
	Password  string `json:"pwd"`
	Mobile    string `json:"mobile"`
	Nickname  string `json:"nick"`
	AvatarUrl string `json:"avatarUrl"`
	Province  string `json:"province"`
	City      string `json:"city"`
	AutoLogin uint   `json:"autoLogin"`
	UserInfo  string `json:"userInfo"`
	UserLevel uint   `json:"userLevel"`
}
