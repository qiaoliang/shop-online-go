package user

import (
	"testing"

	"bookstore/app/configs"

	"github.com/stretchr/testify/assert"
)

func TestUserToVM_NilUser(t *testing.T) {
	// 测试当参数为 nil 时的情况
	result := userToVM(nil)

	// 验证返回值也为 nil
	assert.Nil(t, result, "当输入为 nil 时，userToVM 应该返回 nil")
}

func TestUserToVM_ValidUser(t *testing.T) {
	// 初始化配置
	configs.GetConfigInstance("../../config-test.yaml")

	// 创建一个测试用户
	userLevel := &UserLevel{
		Level: GREENTYPE,
		Name:  "Green",
	}

	user := &User{
		Id:          "test-id",
		Pwd:         "test-pwd",
		Mobile:      "13800138000",
		Nickname:    "测试用户",
		AvatarUrl:   "test-avatar.jpeg",
		Province:    "测试省",
		City:        "测试市",
		AutoLogin:   1,
		UserInfo:    "测试信息",
		UserLevelId: GREENTYPE,
		UserLevel:   userLevel,
	}

	// 调用被测试函数
	result := userToVM(user)

	// 验证结果
	assert.NotNil(t, result, "当输入有效用户时，userToVM 不应返回 nil")
	assert.Equal(t, user.Mobile, result.Token, "Token 应该等于用户的手机号")
	assert.Equal(t, *user, result.UserInfo, "UserInfo 应该等于输入的用户信息")
	assert.Equal(t, *userLevel, result.UserLevel, "UserLevel 应该等于用户的等级信息")

	// 验证头像URL是否已正确添加前缀
	expectedAvatarUrl := configs.Cfg.AvatarPicPrefix() + "test-avatar.jpeg"
	assert.Equal(t, expectedAvatarUrl, result.UserInfo.AvatarUrl, "头像URL应该已添加正确的前缀")
}