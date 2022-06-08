package user

type User struct {
	Id          string     `json:"id" gorm:"primary_key"`
	Password    string     `json:"pwd"`
	Mobile      string     `json:"mobile"`
	Nickname    string     `json:"nick"`
	AvatarUrl   string     `json:"avatarUrl"`
	Province    string     `json:"province"`
	City        string     `json:"city"`
	AutoLogin   uint       `json:"autoLogin"`
	UserInfo    string     `json:"userInfo"`
	UserLevelId int32      `json:"-"`
	UserLevel   *UserLevel `json:"userLevel" gorm:"-"`
}
type UserLevel struct {
	Level UserType `json:"id"`
	Name  string   `json:"name"`
}

type UserType int32

const (
	BLANKTYPE UserType = iota
	GREENTYPE
	REDTYPE
)

func (c UserType) String() string {
	switch c {
	case BLANKTYPE:
		return "NewBee"
	case GREENTYPE:
		return "Green"
	case REDTYPE:
		return "Red"
	}
	return "N/A"
}
