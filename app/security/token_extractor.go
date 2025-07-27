package security

import (
	"bookstore/app/user"
)

// UserTokenExtractor 实现TokenExtractor接口
type UserTokenExtractor struct {
	userRepo user.UserRepo
}

// NewUserTokenExtractor 创建新的用户token提取器
func NewUserTokenExtractor(userRepo user.UserRepo) *UserTokenExtractor {
	return &UserTokenExtractor{
		userRepo: userRepo,
	}
}

// ExtractUserFromToken 从token中提取用户信息
func (ute *UserTokenExtractor) ExtractUserFromToken(token string) *AuthContext {
	if token == "" {
		return nil
	}

	// 当前实现：token就是手机号
	// TODO: 未来可以升级为JWT或其他token机制
	mobileNumber := token

	// 通过手机号查找用户
	user := ute.userRepo.RetriveUserByMobile(mobileNumber)
	if user == nil {
		return nil
	}

	return &AuthContext{
		UserID:  user.Id,
		Mobile:  user.Mobile,
		IsValid: true,
	}
}

// SimpleTokenExtractor 简单的token提取器，用于测试或临时使用
type SimpleTokenExtractor struct {
	validTokens map[string]string // token -> mobile 映射
}

// NewSimpleTokenExtractor 创建简单的token提取器
func NewSimpleTokenExtractor() *SimpleTokenExtractor {
	return &SimpleTokenExtractor{
		validTokens: make(map[string]string),
	}
}

// AddValidToken 添加有效的token
func (ste *SimpleTokenExtractor) AddValidToken(token, mobile string) {
	ste.validTokens[token] = mobile
}

// ExtractUserFromToken 从token中提取用户信息（简单实现）
func (ste *SimpleTokenExtractor) ExtractUserFromToken(token string) *AuthContext {
	if token == "" {
		return nil
	}

	mobile, exists := ste.validTokens[token]
	if !exists {
		return nil
	}

	return &AuthContext{
		UserID:  mobile, // 简单实现中，用户ID就是手机号
		Mobile:  mobile,
		IsValid: true,
	}
}