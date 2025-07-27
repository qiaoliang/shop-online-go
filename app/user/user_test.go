package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Inflate(t *testing.T) {
	// 测试inflate方法
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
		UserLevel:   nil, // 初始化为nil
	}

	// 调用inflate方法
	user.inflate()

	// 验证UserLevel被正确初始化
	assert.NotNil(t, user.UserLevel, "UserLevel 应该被初始化")
	assert.Equal(t, GREENTYPE, user.UserLevel.Level, "UserLevel.Level 应该等于 UserLevelId")
	assert.Equal(t, "Green", user.UserLevel.Name, "UserLevel.Name 应该等于 'Green'")
}

func TestUser_InflateWithExistingUserLevel(t *testing.T) {
	// 测试inflate方法，当UserLevel已经存在时
	existingUserLevel := &UserLevel{
		Level: REDTYPE,
		Name:  "Red",
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
		UserLevelId: GREENTYPE, // 注意这里设置的是GREENTYPE
		UserLevel:   existingUserLevel,
	}

	// 调用inflate方法
	user.inflate()

	// 验证UserLevel被更新为正确的值
	assert.NotNil(t, user.UserLevel, "UserLevel 应该存在")
	assert.Equal(t, GREENTYPE, user.UserLevel.Level, "UserLevel.Level 应该等于 UserLevelId")
	assert.Equal(t, "Green", user.UserLevel.Name, "UserLevel.Name 应该等于 'Green'")
}

func TestUserType_String(t *testing.T) {
	testCases := []struct {
		name     string
		userType UserType
		expected string
	}{
		{"BLANKTYPE", BLANKTYPE, "NewBee"},
		{"GREENTYPE", GREENTYPE, "Green"},
		{"REDTYPE", REDTYPE, "Red"},
		{"Unknown", UserType(999), "N/A"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.userType.String()
			assert.Equal(t, tc.expected, result, "UserType.String() 应该返回正确的字符串")
		})
	}
}

func TestUser_Fields(t *testing.T) {
	// 测试User结构体的字段
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
		UserLevel:   &UserLevel{GREENTYPE, "Green"},
	}

	// 验证所有字段
	assert.Equal(t, "test-id", user.Id, "Id 字段应该正确")
	assert.Equal(t, "test-pwd", user.Pwd, "Pwd 字段应该正确")
	assert.Equal(t, "13800138000", user.Mobile, "Mobile 字段应该正确")
	assert.Equal(t, "测试用户", user.Nickname, "Nickname 字段应该正确")
	assert.Equal(t, "test-avatar.jpeg", user.AvatarUrl, "AvatarUrl 字段应该正确")
	assert.Equal(t, "测试省", user.Province, "Province 字段应该正确")
	assert.Equal(t, "测试市", user.City, "City 字段应该正确")
	assert.Equal(t, uint(1), user.AutoLogin, "AutoLogin 字段应该正确")
	assert.Equal(t, "测试信息", user.UserInfo, "UserInfo 字段应该正确")
	assert.Equal(t, GREENTYPE, user.UserLevelId, "UserLevelId 字段应该正确")
	assert.NotNil(t, user.UserLevel, "UserLevel 字段应该不为nil")
	assert.Equal(t, GREENTYPE, user.UserLevel.Level, "UserLevel.Level 字段应该正确")
	assert.Equal(t, "Green", user.UserLevel.Name, "UserLevel.Name 字段应该正确")
}

func TestUser_EmptyFields(t *testing.T) {
	// 测试User结构体的空字段
	user := &User{}

	// 验证空字段
	assert.Equal(t, "", user.Id, "空Id 字段应该为空字符串")
	assert.Equal(t, "", user.Pwd, "空Pwd 字段应该为空字符串")
	assert.Equal(t, "", user.Mobile, "空Mobile 字段应该为空字符串")
	assert.Equal(t, "", user.Nickname, "空Nickname 字段应该为空字符串")
	assert.Equal(t, "", user.AvatarUrl, "空AvatarUrl 字段应该为空字符串")
	assert.Equal(t, "", user.Province, "空Province 字段应该为空字符串")
	assert.Equal(t, "", user.City, "空City 字段应该为空字符串")
	assert.Equal(t, uint(0), user.AutoLogin, "空AutoLogin 字段应该为0")
	assert.Equal(t, "", user.UserInfo, "空UserInfo 字段应该为空字符串")
	assert.Equal(t, UserType(0), user.UserLevelId, "空UserLevelId 字段应该为0")
	assert.Nil(t, user.UserLevel, "空UserLevel 字段应该为nil")
}

func TestUserLevel_Fields(t *testing.T) {
	// 测试UserLevel结构体的字段
	userLevel := &UserLevel{
		Level: REDTYPE,
		Name:  "Red",
	}

	// 验证字段
	assert.Equal(t, REDTYPE, userLevel.Level, "Level 字段应该正确")
	assert.Equal(t, "Red", userLevel.Name, "Name 字段应该正确")
}

func TestUserLevel_EmptyFields(t *testing.T) {
	// 测试UserLevel结构体的空字段
	userLevel := &UserLevel{}

	// 验证空字段
	assert.Equal(t, UserType(0), userLevel.Level, "空Level 字段应该为0")
	assert.Equal(t, "", userLevel.Name, "空Name 字段应该为空字符串")
}

func TestUser_WithSpecialCharacters(t *testing.T) {
	// 测试包含特殊字符的User
	user := &User{
		Id:          "test-id@#$%",
		Pwd:         "test-pwd@#$%",
		Mobile:      "13800138000",
		Nickname:    "测试用户@#$%",
		AvatarUrl:   "test-avatar@#$%.jpeg",
		Province:    "测试省@#$%",
		City:        "测试市@#$%",
		AutoLogin:   1,
		UserInfo:    "测试信息@#$%",
		UserLevelId: GREENTYPE,
		UserLevel:   &UserLevel{GREENTYPE, "Green@#$%"},
	}

	// 验证包含特殊字符的字段
	assert.Equal(t, "test-id@#$%", user.Id, "包含特殊字符的Id 字段应该正确")
	assert.Equal(t, "test-pwd@#$%", user.Pwd, "包含特殊字符的Pwd 字段应该正确")
	assert.Equal(t, "测试用户@#$%", user.Nickname, "包含特殊字符的Nickname 字段应该正确")
	assert.Equal(t, "test-avatar@#$%.jpeg", user.AvatarUrl, "包含特殊字符的AvatarUrl 字段应该正确")
	assert.Equal(t, "测试省@#$%", user.Province, "包含特殊字符的Province 字段应该正确")
	assert.Equal(t, "测试市@#$%", user.City, "包含特殊字符的City 字段应该正确")
	assert.Equal(t, "测试信息@#$%", user.UserInfo, "包含特殊字符的UserInfo 字段应该正确")
	assert.Equal(t, "Green@#$%", user.UserLevel.Name, "包含特殊字符的UserLevel.Name 字段应该正确")
}

func TestUser_WithUnicodeCharacters(t *testing.T) {
	// 测试包含Unicode字符的User
	user := &User{
		Id:          "test-id🚀🎉💯",
		Pwd:         "test-pwd🚀🎉💯",
		Mobile:      "13800138000",
		Nickname:    "测试用户🚀🎉💯",
		AvatarUrl:   "test-avatar🚀🎉💯.jpeg",
		Province:    "测试省🚀🎉💯",
		City:        "测试市🚀🎉💯",
		AutoLogin:   1,
		UserInfo:    "测试信息🚀🎉💯",
		UserLevelId: GREENTYPE,
		UserLevel:   &UserLevel{GREENTYPE, "Green🚀🎉💯"},
	}

	// 验证包含Unicode字符的字段
	assert.Equal(t, "test-id🚀🎉💯", user.Id, "包含Unicode字符的Id 字段应该正确")
	assert.Equal(t, "test-pwd🚀🎉💯", user.Pwd, "包含Unicode字符的Pwd 字段应该正确")
	assert.Equal(t, "测试用户🚀🎉💯", user.Nickname, "包含Unicode字符的Nickname 字段应该正确")
	assert.Equal(t, "test-avatar🚀🎉💯.jpeg", user.AvatarUrl, "包含Unicode字符的AvatarUrl 字段应该正确")
	assert.Equal(t, "测试省🚀🎉💯", user.Province, "包含Unicode字符的Province 字段应该正确")
	assert.Equal(t, "测试市🚀🎉💯", user.City, "包含Unicode字符的City 字段应该正确")
	assert.Equal(t, "测试信息🚀🎉💯", user.UserInfo, "包含Unicode字符的UserInfo 字段应该正确")
	assert.Equal(t, "Green🚀🎉💯", user.UserLevel.Name, "包含Unicode字符的UserLevel.Name 字段应该正确")
}

func TestUser_WithLongValues(t *testing.T) {
	// 测试包含长值的User
	longString := "这是一个非常长的字符串，用来测试User结构体是否能正确处理长文本内容。这个字符串包含了中文字符、英文字符、数字和特殊符号，确保结构体字段不会出现截断或编码问题。"

	user := &User{
		Id:          longString,
		Pwd:         longString,
		Mobile:      "13800138000",
		Nickname:    longString,
		AvatarUrl:   longString,
		Province:    longString,
		City:        longString,
		AutoLogin:   1,
		UserInfo:    longString,
		UserLevelId: GREENTYPE,
		UserLevel:   &UserLevel{GREENTYPE, longString},
	}

	// 验证长值字段
	assert.Equal(t, longString, user.Id, "长Id 字段应该正确")
	assert.Equal(t, longString, user.Pwd, "长Pwd 字段应该正确")
	assert.Equal(t, longString, user.Nickname, "长Nickname 字段应该正确")
	assert.Equal(t, longString, user.AvatarUrl, "长AvatarUrl 字段应该正确")
	assert.Equal(t, longString, user.Province, "长Province 字段应该正确")
	assert.Equal(t, longString, user.City, "长City 字段应该正确")
	assert.Equal(t, longString, user.UserInfo, "长UserInfo 字段应该正确")
	assert.Equal(t, longString, user.UserLevel.Name, "长UserLevel.Name 字段应该正确")
}

func TestUser_WithDifferentUserTypes(t *testing.T) {
	// 测试不同的UserType值
	testCases := []struct {
		name        string
		userType    UserType
		expectedStr string
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
				UserLevelId: tc.userType,
				UserLevel:   &UserLevel{tc.userType, tc.expectedStr},
			}

			// 验证UserType相关字段
			assert.Equal(t, tc.userType, user.UserLevelId, "UserLevelId 字段应该正确")
			assert.Equal(t, tc.userType, user.UserLevel.Level, "UserLevel.Level 字段应该正确")
			assert.Equal(t, tc.expectedStr, user.UserLevel.Name, "UserLevel.Name 字段应该正确")
			assert.Equal(t, tc.expectedStr, tc.userType.String(), "UserType.String() 应该返回正确的字符串")
		})
	}
}

func TestUser_WithDifferentAutoLoginValues(t *testing.T) {
	// 测试不同的AutoLogin值
	testCases := []struct {
		name     string
		autoLogin uint
	}{
		{"Zero", 0},
		{"One", 1},
		{"Large", 999999},
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
				AutoLogin:   tc.autoLogin,
				UserInfo:    "测试信息",
				UserLevelId: GREENTYPE,
				UserLevel:   &UserLevel{GREENTYPE, "Green"},
			}

			// 验证AutoLogin字段
			assert.Equal(t, tc.autoLogin, user.AutoLogin, "AutoLogin 字段应该正确")
		})
	}
}

func TestUser_InflateMultipleTimes(t *testing.T) {
	// 测试多次调用inflate方法
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
		UserLevel:   nil,
	}

	// 第一次调用inflate
	user.inflate()
	assert.NotNil(t, user.UserLevel, "第一次inflate后UserLevel应该不为nil")
	assert.Equal(t, GREENTYPE, user.UserLevel.Level, "第一次inflate后UserLevel.Level应该正确")
	assert.Equal(t, "Green", user.UserLevel.Name, "第一次inflate后UserLevel.Name应该正确")

	// 修改UserLevelId
	user.UserLevelId = REDTYPE

	// 第二次调用inflate
	user.inflate()
	assert.NotNil(t, user.UserLevel, "第二次inflate后UserLevel应该不为nil")
	assert.Equal(t, REDTYPE, user.UserLevel.Level, "第二次inflate后UserLevel.Level应该更新")
	assert.Equal(t, "Red", user.UserLevel.Name, "第二次inflate后UserLevel.Name应该更新")
}