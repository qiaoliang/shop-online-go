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

	// 确保 UserLevel 不为 nil
	if user.UserLevel == nil {
		user.inflate() // 调用 inflate 方法初始化 UserLevel
	}

	return &UserVM{
		user.Mobile,
		*user,
		*user.UserLevel,
	}
}
