package user

import (
	"bookstore/app/utils"
	"errors"

	"gorm.io/gorm"
)

type UserRepoIf interface {
	TotalUsers() int
	findUser(mobile string, pwd string) *User
	retriveUserByMobile(mobile string) *User
	CreateUser(mobile string, pwd string, nickname string, genUserId UserIdGen) (user *User, err error)
	CreateAdmin(mobile string, pwd string)
	DeleteByMobile(mobile string)
}

type UserRepoDB struct {
	userlist map[string]*User
	db       *gorm.DB
}

func GetUserRepoDB(db *gorm.DB) UserRepoIf {
	if userRepo == nil {
		userRepo = &UserRepoDB{make(map[string]*User, 10), db}
	}
	return userRepo
}

func (r *UserRepoDB) TotalUsers() int {
	var users []User
	r.db.Find(&users)
	return len(users)
}
func (r *UserRepoDB) DeleteByMobile(mobile string) {
	var user User
	r.db.Where("mobile = ?", mobile).Delete(&user)
}
func (r *UserRepoDB) findUser(mobile string, pwd string) *User {
	found := r.retriveUserByMobile(mobile)
	if found == nil || found.Password != pwd {
		return nil
	}
	return found
}
func (r *UserRepoDB) retriveUserByMobile(mobile string) *User {
	var user User
	result := r.db.Where("mobile = ?", mobile).First(&user)
	if result.Error != nil {
		return nil
	}
	return &user
}

func (r *UserRepoDB) CreateUser(mobile string, pwd string, nickname string, genUserId UserIdGen) (user *User, err error) {
	if r.findUser(mobile, pwd) != nil {
		return nil, errors.New("hello,error")
	}
	userId := genUserId()
	avatarUrl := utils.NewRandom().GenAavatarStr()
	newUser := &User{
		Id:          userId,
		Password:    pwd,
		Mobile:      mobile,
		Nickname:    nickname,
		AvatarUrl:   avatarUrl,
		Province:    "未知",
		City:        "未知",
		AutoLogin:   0,
		UserInfo:    "FakeUserInfo",
		UserLevelId: 1,
		UserLevel:   &UserLevel{GREENTYPE, GREENTYPE.String()},
	}
	r.db.Create(newUser)
	if r.db.Error != nil {
		return nil, r.db.Error
	}
	r.userlist[mobile] = newUser
	return r.userlist[mobile], nil
}
func (r *UserRepoDB) CreateAdmin(mobile string, pwd string) {
	r.CreateUser(mobile, pwd, "超级塞亚人", genUId)
}
