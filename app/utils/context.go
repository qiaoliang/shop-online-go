package utils

import (
	"github.com/gin-gonic/gin"
)

// GetUserIDFromContext 从gin上下文中获取用户ID
func GetUserIDFromContext(c *gin.Context) string {
	if userID, exists := c.Get("userID"); exists {
		if id, ok := userID.(string); ok {
			return id
		}
	}
	return ""
}

// GetMobileFromContext 从gin上下文中获取手机号
func GetMobileFromContext(c *gin.Context) string {
	if mobile, exists := c.Get("mobile"); exists {
		if m, ok := mobile.(string); ok {
			return m
		}
	}
	return ""
}

// GetAuthContextFromContext 从gin上下文中获取完整的认证上下文
func GetAuthContextFromContext(c *gin.Context) interface{} {
	if authContext, exists := c.Get("authContext"); exists {
		return authContext
	}
	return nil
}