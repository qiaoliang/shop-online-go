package user

import (
	"testing"
)

func TestUpdateUserByToken(t *testing.T) {
	// 创建一个测试用的 UserService
	us := newUserService(false)

	// 注册一个测试用户
	mobile := "13900001234"
	pwd := "password123"
	nickname := "测试用户"
	autoLogin := "1"

	user, err := us.RegisterNewUser(mobile, pwd, nickname, autoLogin)
	if err != nil {
		t.Fatalf("注册用户失败: %v", err)
	}

	// 确认用户已创建
	if user == nil {
		t.Fatal("用户创建失败")
	}

	// 使用手机号作为token（根据当前实现）
	token := mobile

	// 准备更新数据
	newNickname := "更新后的昵称"
	newAvatarUrl := "http://example.com/new-avatar.jpg"
	newProvince := "广东省"
	newCity := "深圳市"

	updateData := User{
		Nickname:  newNickname,
		AvatarUrl: newAvatarUrl,
		Province:  newProvince,
		City:      newCity,
	}

	// 执行更新
	updatedUser := us.UpdateUserByToken(token, updateData)

	// 验证更新结果
	if updatedUser == nil {
		t.Fatal("更新用户信息失败，返回了nil")
	}

	if updatedUser.Nickname != newNickname {
		t.Errorf("昵称更新失败，期望: %s, 实际: %s", newNickname, updatedUser.Nickname)
	}

	if updatedUser.AvatarUrl != newAvatarUrl {
		t.Errorf("头像URL更新失败，期望: %s, 实际: %s", newAvatarUrl, updatedUser.AvatarUrl)
	}

	if updatedUser.Province != newProvince {
		t.Errorf("省份更新失败，期望: %s, 实际: %s", newProvince, updatedUser.Province)
	}

	if updatedUser.City != newCity {
		t.Errorf("城市更新失败，期望: %s, 实际: %s", newCity, updatedUser.City)
	}

	// 再次获取用户，确认更新已持久化
	fetchedUser := us.FindUserByToken(token)
	if fetchedUser == nil {
		t.Fatal("无法通过token获取更新后的用户")
	}

	if fetchedUser.Nickname != newNickname {
		t.Errorf("持久化后昵称不匹配，期望: %s, 实际: %s", newNickname, fetchedUser.Nickname)
	}
}