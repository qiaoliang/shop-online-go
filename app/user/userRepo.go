package user

import (
	"bookstore/app/utils"
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
	found := userRepo.userlist[mobile]
	if found == nil || found.Password != pwd {
		return nil
	}
	return found
}
func (r *MemoryUserRepo) CreateUser(mobile string, pwd string, nickname string) *User {

	userId := "userId" + utils.GenerateStr(10)
	avatarUrl := "" + utils.GenerateAavatarStr()
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
		UserLevel: 0,
	}
	return r.userlist[mobile]
}
func (r *MemoryUserRepo) CreateAdmin(mobile string, pwd string) {
	r.CreateUser(mobile, pwd, "超级塞亚人")
}
