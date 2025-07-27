package security

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthContext 定义认证上下文中的用户信息
type AuthContext struct {
	UserID   string `json:"user_id"`
	Mobile   string `json:"mobile"`
	IsValid  bool   `json:"is_valid"`
}

// TokenExtractor 定义token提取接口，避免直接依赖UserService
type TokenExtractor interface {
	ExtractUserFromToken(token string) *AuthContext
}

// AuthMiddleware 认证中间件
type AuthMiddleware struct {
	tokenExtractor TokenExtractor
}

// NewAuthMiddleware 创建新的认证中间件
func NewAuthMiddleware(extractor TokenExtractor) *AuthMiddleware {
	return &AuthMiddleware{
		tokenExtractor: extractor,
	}
}

// Authenticate 认证中间件函数
func (am *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中提取token
		token := extractTokenFromRequest(c)

		if token == "" {
			// 对于不需要认证的接口，继续处理
			c.Next()
			return
		}

		// 使用token提取器获取用户信息
		authContext := am.tokenExtractor.ExtractUserFromToken(token)

		if authContext == nil || !authContext.IsValid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无效的认证token",
			})
			c.Abort()
			return
		}

		// 将用户信息注入到gin上下文中
		c.Set("userID", authContext.UserID)
		c.Set("mobile", authContext.Mobile)
		c.Set("authContext", authContext)

		c.Next()
	}
}

// RequireAuth 要求认证的中间件
func (am *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中提取token
		token := extractTokenFromRequest(c)

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "缺少认证token",
			})
			c.Abort()
			return
		}

		// 使用token提取器获取用户信息
		authContext := am.tokenExtractor.ExtractUserFromToken(token)

		if authContext == nil || !authContext.IsValid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无效的认证token",
			})
			c.Abort()
			return
		}

		// 将用户信息注入到gin上下文中
		c.Set("userID", authContext.UserID)
		c.Set("mobile", authContext.Mobile)
		c.Set("authContext", authContext)

		c.Next()
	}
}

// extractTokenFromRequest 从请求中提取token
func extractTokenFromRequest(c *gin.Context) string {
	// 1. 从Authorization header中提取Bearer token
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1]
		}
	}

	// 2. 从query参数中提取token
	if token := c.Query("token"); token != "" {
		return token
	}

	// 3. 从form参数中提取token
	if token := c.PostForm("token"); token != "" {
		return token
	}

	// 4. 从JSON body中提取token（如果Content-Type是application/json）
	if c.ContentType() == "application/json" {
		var body map[string]interface{}
		if err := c.ShouldBindJSON(&body); err == nil {
			if token, ok := body["token"].(string); ok && token != "" {
				return token
			}
		}
	}

	return ""
}