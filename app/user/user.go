package user

type User struct {
	Id          string     `json:"id" gorm:"column:id;primaryKey"`
	Pwd         string     `json:"pwd" gorm:"column:pwd"`
	Mobile      string     `json:"mobile"`
	Nickname    string     `json:"nick"`
	AvatarUrl   string     `json:"avatarUrl" gorm:"column:avatar_url"`
	Province    string     `json:"province"`
	City        string     `json:"city"`
	AutoLogin   uint       `json:"autoLogin" gorm:"column:auto_login"`
	UserInfo    string     `json:"userInfo" gorm:"column:user_info"`
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
