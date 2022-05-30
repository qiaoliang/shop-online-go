package user

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
	"errors"
	"fmt"
	"sync"
)

var lockUR = &sync.Mutex{}

type MemoryUserRepo struct {
	userlist map[string]*User
}

var userRepo *MemoryUserRepo

func GetUserRepoInstance() *MemoryUserRepo {
	lockUR.Lock()
	defer lockUR.Unlock()
	if userRepo == nil {
		userRepo = &MemoryUserRepo{}
		userRepo.userlist = make(map[string]*User, 10)
		userRepo.CreateAdmin("13900007997", "1234")
	}
	return userRepo
}

func (r *MemoryUserRepo) TotalUsers() int {
	return len(r.userlist)
}

func (r *MemoryUserRepo) findUser(mobile string, pwd string) *User {
	found := r.retriveUserByMobile(mobile)
	if found == nil || found.Password != pwd {
		return nil
	}
	return found
}
func (r *MemoryUserRepo) retriveUserByMobile(mobile string) *User {
	return userRepo.userlist[mobile]
}

func (r *MemoryUserRepo) CreateUser(mobile string, pwd string, nickname string) (user *User, err error) {
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
		UserLevel: &LEVELGREEN,
	}
	return r.userlist[mobile], nil
}
func (r *MemoryUserRepo) CreateAdmin(mobile string, pwd string) {
	r.CreateUser(mobile, pwd, "超级塞亚人")
}
