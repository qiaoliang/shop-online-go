package user

import (
	"bookstore/app/utils"
	"fmt"
	"sync"
)

var once sync.Once

type MemoryUserRepo struct {
	userlist map[string]User
}

var userRepo *MemoryUserRepo

func GetUserRepoInstance() *MemoryUserRepo {
	if userRepo == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				userRepo = &MemoryUserRepo{}
				userRepo.userlist = make(map[string]User)
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return userRepo
}

func (r *MemoryUserRepo) TotalUsers() int {
	return len(r.userlist)
}
func (r *MemoryUserRepo) CreateUser(mobile string, pwd string, nickname string) *User {

	userId := "userId" + utils.GenerateStr(10)
	avatarUrl := "" + utils.GenerateAavatarStr()
	usr := User{
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
	r.userlist[userId] = usr
	return &usr
}
