package user

type User struct {
	Id        string     `json:"id"`
	Password  string     `json:"pwd"`
	Mobile    string     `json:"mobile"`
	Nickname  string     `json:"nick"`
	AvatarUrl string     `json:"avatarUrl"`
	Province  string     `json:"province"`
	City      string     `json:"city"`
	AutoLogin uint       `json:"autoLogin"`
	UserInfo  string     `json:"userInfo"`
	UserLevel *UserLevel `json:"userLevel"`
}
type UserLevel struct {
	Level uint   `json:"id"`
	Name  string `json:"name"`
}

var LEVELGREEN = UserLevel{1, "Green"}
var LEVELRED = UserLevel{2, "Red"}
var LEVELBlANK = UserLevel{0, "Blank"}
