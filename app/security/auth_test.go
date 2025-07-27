package security

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockTokenExtractor 用于测试的mock token提取器
type MockTokenExtractor struct {
	validTokens map[string]*AuthContext
}

func NewMockTokenExtractor() *MockTokenExtractor {
	return &MockTokenExtractor{
		validTokens: make(map[string]*AuthContext),
	}
}

func (m *MockTokenExtractor) AddValidToken(token string, userID, mobile string) {
	m.validTokens[token] = &AuthContext{
		UserID:  userID,
		Mobile:  mobile,
		IsValid: true,
	}
}

func (m *MockTokenExtractor) ExtractUserFromToken(token string) *AuthContext {
	return m.validTokens[token]
}

func TestAuthMiddleware_Authenticate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		token          string
		expectedStatus int
		expectedUserID string
		expectedMobile string
	}{
		{
			name:           "有效token",
			token:          "valid_token",
			expectedStatus: http.StatusOK,
			expectedUserID: "user123",
			expectedMobile: "13900007997",
		},
		{
			name:           "无效token",
			token:          "invalid_token",
			expectedStatus: http.StatusUnauthorized,
			expectedUserID: "",
			expectedMobile: "",
		},
		{
			name:           "无token",
			token:          "",
			expectedStatus: http.StatusOK,
			expectedUserID: "",
			expectedMobile: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建mock token提取器
			mockExtractor := NewMockTokenExtractor()
			mockExtractor.AddValidToken("valid_token", "user123", "13900007997")

			// 创建认证中间件
			authMiddleware := NewAuthMiddleware(mockExtractor)

			// 创建gin路由
			router := gin.New()
			router.Use(authMiddleware.Authenticate())
			router.GET("/test", func(c *gin.Context) {
				userID := c.GetString("userID")
				mobile := c.GetString("mobile")
				c.JSON(http.StatusOK, gin.H{
					"userID": userID,
					"mobile": mobile,
				})
			})

			// 创建测试请求
			req, _ := http.NewRequest("GET", "/test", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}

			// 执行请求
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// 验证结果
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				// 验证响应中的用户信息
				response := w.Body.String()
				assert.Contains(t, response, tt.expectedUserID)
				assert.Contains(t, response, tt.expectedMobile)
			}
		})
	}
}

func TestExtractTokenFromRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		setupRequest   func(*http.Request)
		expectedToken  string
	}{
		{
			name: "从Authorization header提取Bearer token",
			setupRequest: func(req *http.Request) {
				req.Header.Set("Authorization", "Bearer test_token_123")
			},
			expectedToken: "test_token_123",
		},
		{
			name: "从query参数提取token",
			setupRequest: func(req *http.Request) {
				req.URL.RawQuery = "token=query_token_456"
			},
			expectedToken: "query_token_456",
		},
		{
			name: "从form参数提取token",
			setupRequest: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				req.Body = http.NoBody
			},
			expectedToken: "",
		},
		{
			name: "无token",
			setupRequest: func(req *http.Request) {
				// 不设置任何token
			},
			expectedToken: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/test", nil)
			tt.setupRequest(req)

			// 创建gin上下文
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request = req

			// 提取token
			token := extractTokenFromRequest(c)

			// 验证结果
			assert.Equal(t, tt.expectedToken, token)
		})
	}
}