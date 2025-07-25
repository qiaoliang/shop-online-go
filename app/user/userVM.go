package user

import "bookstore/app/configs"

type UserVM struct {
	Token     string    `json:"token"`
	UserInfo  User      `json:"base"`
	UserLevel UserLevel `json:"userLevel"`
}

func userToVM(user *User) *UserVM {
	if user == nil {
		return nil
	}
	user.AvatarUrl = configs.Cfg.AvatarPicPrefix() + user.AvatarUrl
	return &UserVM{
		user.Mobile,
		*user,
		*user.UserLevel,
	}
}
