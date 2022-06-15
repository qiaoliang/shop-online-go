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
	UserLevelId UserType   `json:"-"  gorm:"column:User_Level_Id"`
	UserLevel   *UserLevel `json:"userLevel" gorm:"-"`
}

type UserLevel struct {
	Level UserType `json:"id"`
	Name  string   `json:"name"`
}

func (s *User) inflate() {
	s.UserLevel = &UserLevel{s.UserLevelId, s.UserLevelId.String()}
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
