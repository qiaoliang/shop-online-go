package user

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
	"errors"
	"fmt"
)

type UserRepoIf interface {
	TotalUsers() int
	findUser(mobile string, pwd string) *User
	retriveUserByMobile(mobile string) *User
	CreateUser(mobile string, pwd string, nickname string) (user *User, err error)
	CreateAdmin(mobile string, pwd string)
}

type UserRepoDB struct {
	userlist map[string]*User
	db       *configs.DBConn
}

func GetUserRepoDB(db *configs.DBConn) UserRepoIf {
	if userRepo == nil {
		userRepo = &UserRepoDB{make(map[string]*User, 10), db}
		userRepo.CreateAdmin("13900007997", "1234")
	}
	return userRepo
}

func (r *UserRepoDB) TotalUsers() int {
	return len(r.userlist)
}

func (r *UserRepoDB) findUser(mobile string, pwd string) *User {
	found := r.retriveUserByMobile(mobile)
	if found == nil || found.Password != pwd {
		return nil
	}
	return found
}
func (r *UserRepoDB) retriveUserByMobile(mobile string) *User {
	return r.userlist[mobile]
}

func (r *UserRepoDB) CreateUser(mobile string, pwd string, nickname string) (user *User, err error) {
	if r.findUser(mobile, pwd) != nil {
		return nil, errors.New("hello,error")
	}
	userId := fmt.Sprintf("userId%v", utils.RandomStr(10))
	avatarUrl := configs.Cfg.AvatarPicPrefix() + utils.GenerateAavatarStr()
	r.userlist[mobile] = &User{
		Id:        userId,
		Password:  pwd,
		Mobile:    mobile,
		Nickname:  nickname,
		AvatarUrl: avatarUrl,
		Province:  "未知",
		City:      "未知",
		AutoLogin: 0,
		UserInfo:  "FakeUserInfo",
		UserLevel: &UserLevel{GREENTYPE, GREENTYPE.String()},
	}
	return r.userlist[mobile], nil
}
func (r *UserRepoDB) CreateAdmin(mobile string, pwd string) {
	r.CreateUser(mobile, pwd, "超级塞亚人")
}
