package user

import "bookstore/app/utils"

type UserRepo interface {
	CreateUser(mobile string) *User
}

type MemoryUserRepo struct {
	repo map[string]*User
}

func (*MemoryUserRepo) CreateUser(mobile string, pwd string, nickname string) *User {

	userId := "userId" + utils.GenerateStr(10)
	avatarUrl := "" + utils.GenerateAavatarStr()
	usr := &User{
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

	return usr
}
