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

func TestUserToVM_UserWithNilUserLevel(t *testing.T) {
	// 初始化配置
	configs.GetConfigInstance("../../config-test.yaml")

	// 创建一个没有UserLevel的用户
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
		UserLevel:   nil, // 故意设置为nil
	}

	// 调用被测试函数
	result := userToVM(user)

	// 验证结果
	assert.NotNil(t, result, "当输入有效用户时，userToVM 不应返回 nil")
	assert.Equal(t, user.Mobile, result.Token, "Token 应该等于用户的手机号")
	assert.Equal(t, *user, result.UserInfo, "UserInfo 应该等于输入的用户信息")

	// 验证UserLevel被正确初始化
	assert.NotNil(t, result.UserLevel, "UserLevel 应该被自动初始化")
	assert.Equal(t, GREENTYPE, result.UserLevel.Level, "UserLevel.Level 应该等于 UserLevelId")
	assert.Equal(t, "Green", result.UserLevel.Name, "UserLevel.Name 应该等于 'Green'")

	// 验证头像URL是否已正确添加前缀
	expectedAvatarUrl := configs.Cfg.AvatarPicPrefix() + "test-avatar.jpeg"
	assert.Equal(t, expectedAvatarUrl, result.UserInfo.AvatarUrl, "头像URL应该已添加正确的前缀")
}

func TestUserToVM_UserWithDifferentUserLevels(t *testing.T) {
	// 初始化配置
	configs.GetConfigInstance("../../config-test.yaml")

	testCases := []struct {
		name        string
		userLevelId UserType
		expectedName string
	}{
		{"BLANKTYPE", BLANKTYPE, "NewBee"},
		{"GREENTYPE", GREENTYPE, "Green"},
		{"REDTYPE", REDTYPE, "Red"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
				UserLevelId: tc.userLevelId,
				UserLevel:   nil,
			}

			result := userToVM(user)

			assert.NotNil(t, result, "userToVM 不应返回 nil")
			assert.Equal(t, tc.userLevelId, result.UserLevel.Level, "UserLevel.Level 应该正确")
			assert.Equal(t, tc.expectedName, result.UserLevel.Name, "UserLevel.Name 应该正确")
		})
	}
}

func TestUserToVM_UserWithEmptyAvatarUrl(t *testing.T) {
	// 初始化配置
	configs.GetConfigInstance("../../config-test.yaml")

	user := &User{
		Id:          "test-id",
		Pwd:         "test-pwd",
		Mobile:      "13800138000",
		Nickname:    "测试用户",
		AvatarUrl:   "", // 空头像URL
		Province:    "测试省",
		City:        "测试市",
		AutoLogin:   1,
		UserInfo:    "测试信息",
		UserLevelId: GREENTYPE,
		UserLevel:   &UserLevel{GREENTYPE, "Green"},
	}

	result := userToVM(user)

	assert.NotNil(t, result, "userToVM 不应返回 nil")
	// 验证空头像URL也被添加了前缀
	expectedAvatarUrl := configs.Cfg.AvatarPicPrefix() + ""
	assert.Equal(t, expectedAvatarUrl, result.UserInfo.AvatarUrl, "空头像URL也应该添加前缀")
}

func TestUserToVM_UserWithSpecialCharacters(t *testing.T) {
	// 初始化配置
	configs.GetConfigInstance("../../config-test.yaml")

	user := &User{
		Id:          "test-id",
		Pwd:         "test-pwd",
		Mobile:      "13800138000",
		Nickname:    "测试用户@#$%",
		AvatarUrl:   "test-avatar.jpeg",
		Province:    "测试省@#$%",
		City:        "测试市@#$%",
		AutoLogin:   1,
		UserInfo:    "测试信息@#$%",
		UserLevelId: GREENTYPE,
		UserLevel:   &UserLevel{GREENTYPE, "Green"},
	}

	result := userToVM(user)

	assert.NotNil(t, result, "userToVM 不应返回 nil")
	assert.Equal(t, user.Nickname, result.UserInfo.Nickname, "包含特殊字符的昵称应该正确保留")
	assert.Equal(t, user.Province, result.UserInfo.Province, "包含特殊字符的省份应该正确保留")
	assert.Equal(t, user.City, result.UserInfo.City, "包含特殊字符的城市应该正确保留")
	assert.Equal(t, user.UserInfo, result.UserInfo.UserInfo, "包含特殊字符的用户信息应该正确保留")
}

func TestUserToVM_UserWithLongValues(t *testing.T) {
	// 初始化配置
	configs.GetConfigInstance("../../config-test.yaml")

	longString := "这是一个非常长的字符串，用来测试userToVM函数是否能正确处理长文本内容。这个字符串包含了中文字符、英文字符、数字和特殊符号，确保转换过程中不会出现截断或编码问题。"

	user := &User{
		Id:          "test-id",
		Pwd:         "test-pwd",
		Mobile:      "13800138000",
		Nickname:    longString,
		AvatarUrl:   "test-avatar.jpeg",
		Province:    longString,
		City:        longString,
		AutoLogin:   1,
		UserInfo:    longString,
		UserLevelId: GREENTYPE,
		UserLevel:   &UserLevel{GREENTYPE, "Green"},
	}

	result := userToVM(user)

	assert.NotNil(t, result, "userToVM 不应返回 nil")
	assert.Equal(t, longString, result.UserInfo.Nickname, "长昵称应该正确保留")
	assert.Equal(t, longString, result.UserInfo.Province, "长省份应该正确保留")
	assert.Equal(t, longString, result.UserInfo.City, "长城市应该正确保留")
	assert.Equal(t, longString, result.UserInfo.UserInfo, "长用户信息应该正确保留")
}

func TestUserToVM_UserWithZeroValues(t *testing.T) {
	// 初始化配置
	configs.GetConfigInstance("../../config-test.yaml")

	user := &User{
		Id:          "",
		Pwd:         "",
		Mobile:      "",
		Nickname:    "",
		AvatarUrl:   "",
		Province:    "",
		City:        "",
		AutoLogin:   0,
		UserInfo:    "",
		UserLevelId: BLANKTYPE,
		UserLevel:   &UserLevel{BLANKTYPE, "NewBee"},
	}

	result := userToVM(user)

	assert.NotNil(t, result, "userToVM 不应返回 nil")
	assert.Equal(t, "", result.Token, "空手机号应该正确保留")
	assert.Equal(t, "", result.UserInfo.Id, "空ID应该正确保留")
	assert.Equal(t, "", result.UserInfo.Pwd, "空密码应该正确保留")
	assert.Equal(t, "", result.UserInfo.Nickname, "空昵称应该正确保留")
	assert.Equal(t, "", result.UserInfo.Province, "空省份应该正确保留")
	assert.Equal(t, "", result.UserInfo.City, "空城市应该正确保留")
	assert.Equal(t, uint(0), result.UserInfo.AutoLogin, "零值AutoLogin应该正确保留")
	assert.Equal(t, "", result.UserInfo.UserInfo, "空用户信息应该正确保留")
	assert.Equal(t, BLANKTYPE, result.UserLevel.Level, "BLANKTYPE应该正确保留")
	assert.Equal(t, "NewBee", result.UserLevel.Name, "NewBee名称应该正确保留")
}

func TestUserToVM_UserWithUnicodeCharacters(t *testing.T) {
	// 初始化配置
	configs.GetConfigInstance("../../config-test.yaml")

	user := &User{
		Id:          "test-id",
		Pwd:         "test-pwd",
		Mobile:      "13800138000",
		Nickname:    "测试用户🚀🎉💯",
		AvatarUrl:   "test-avatar.jpeg",
		Province:    "测试省🚀🎉💯",
		City:        "测试市🚀🎉💯",
		AutoLogin:   1,
		UserInfo:    "测试信息🚀🎉💯",
		UserLevelId: GREENTYPE,
		UserLevel:   &UserLevel{GREENTYPE, "Green"},
	}

	result := userToVM(user)

	assert.NotNil(t, result, "userToVM 不应返回 nil")
	assert.Equal(t, user.Nickname, result.UserInfo.Nickname, "包含Unicode字符的昵称应该正确保留")
	assert.Equal(t, user.Province, result.UserInfo.Province, "包含Unicode字符的省份应该正确保留")
	assert.Equal(t, user.City, result.UserInfo.City, "包含Unicode字符的城市应该正确保留")
	assert.Equal(t, user.UserInfo, result.UserInfo.UserInfo, "包含Unicode字符的用户信息应该正确保留")
}

func TestUserToVM_UserWithDifferentAvatarFormats(t *testing.T) {
	// 初始化配置
	configs.GetConfigInstance("../../config-test.yaml")

	testCases := []struct {
		name           string
		avatarUrl      string
		expectedSuffix string
	}{
		{"JPEG", "avatar.jpeg", "avatar.jpeg"},
		{"JPG", "avatar.jpg", "avatar.jpg"},
		{"PNG", "avatar.png", "avatar.png"},
		{"GIF", "avatar.gif", "avatar.gif"},
		{"WEBP", "avatar.webp", "avatar.webp"},
		{"WithPath", "path/to/avatar.jpeg", "path/to/avatar.jpeg"},
		{"WithQuery", "avatar.jpeg?size=100", "avatar.jpeg?size=100"},
		{"WithFragment", "avatar.jpeg#section", "avatar.jpeg#section"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := &User{
				Id:          "test-id",
				Pwd:         "test-pwd",
				Mobile:      "13800138000",
				Nickname:    "测试用户",
				AvatarUrl:   tc.avatarUrl,
				Province:    "测试省",
				City:        "测试市",
				AutoLogin:   1,
				UserInfo:    "测试信息",
				UserLevelId: GREENTYPE,
				UserLevel:   &UserLevel{GREENTYPE, "Green"},
			}

			result := userToVM(user)

			assert.NotNil(t, result, "userToVM 不应返回 nil")
			expectedAvatarUrl := configs.Cfg.AvatarPicPrefix() + tc.expectedSuffix
			assert.Equal(t, expectedAvatarUrl, result.UserInfo.AvatarUrl, "头像URL应该正确添加前缀")
		})
	}
}