package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bookstore/app/configs"
	"bookstore/app/testutils"
	"bookstore/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AddressHandlerTestSuite struct {
	testutils.SupperSuite
	router         *gin.Engine
	addressHandler *AddressHandler
	addressService AddressService
	addressRepo    AddressRepository
	userRepo       UserRepo
	testUserID     string
}

func TestAddressHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(AddressHandlerTestSuite))
}

func (s *AddressHandlerTestSuite) SetupSuite() {
	s.SupperSuite.SetupSuite()

	// 设置数据库连接
	db := configs.Cfg.DBConnection()
	s.addressRepo = NewAddressRepositoryDB(db)
	s.userRepo = NewUserRepoDB(db)
	s.addressService = NewAddressService(s.addressRepo, s.userRepo, db)
	s.addressHandler = NewAddressHandler(s.addressService)

	// 创建测试用户
	s.createTestUser()

	// 设置测试路由
	s.router = s.setupTestRouter()
}

func (s *AddressHandlerTestSuite) TeardownSuite() {
	s.SupperSuite.TeardownSuite()
	s.addressHandler = nil
	s.addressService = nil
	s.addressRepo = nil
	s.userRepo = nil
}

func (s *AddressHandlerTestSuite) SetupTest() {
	// 清理测试数据
	s.cleanupTestData()
}

func (s *AddressHandlerTestSuite) TearDownTest() {
	// 清理测试数据
	s.cleanupTestData()
}

// 创建测试用户
func (s *AddressHandlerTestSuite) createTestUser() {
	mobile := "13900001234" + utils.RandomImpl{}.GenStr()
	user, err := s.userRepo.CreateUser(mobile, "password123", "测试用户", "0", genUId)
	s.Nil(err, "创建测试用户失败")
	s.NotNil(user, "用户创建失败")
	s.testUserID = user.Id // 使用真实的用户ID，因为服务层期望用户ID
}

// 清理测试数据
func (s *AddressHandlerTestSuite) cleanupTestData() {
	// 删除测试用户的所有地址
	if s.testUserID != "" {
		addresses, _ := s.addressRepo.ListByUserID(s.testUserID)
		for _, addr := range addresses {
			s.addressRepo.Delete(addr.Id)
		}
	}
}

// 设置测试路由
func (s *AddressHandlerTestSuite) setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// 添加测试专用的认证中间件
	router.Use(s.testAuthMiddleware())

	v1 := router.Group("/v1")
	v1.POST("/user/shipping-address/add", s.addressHandler.AddAddress)
	v1.GET("/user/shipping-address/list", s.addressHandler.GetAddressList)
	v1.GET("/user/shipping-address/default", s.addressHandler.GetDefaultAddress)

	return router
}

// 测试专用的认证中间件
func (s *AddressHandlerTestSuite) testAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Authorization header中获取token
		authHeader := c.GetHeader("Authorization")
		var token string
		if authHeader != "" && len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token = authHeader[7:]
		}

		// 从query参数或form参数中获取token
		if token == "" {
			token = c.Query("token")
		}
		if token == "" {
			token = c.PostForm("token")
		}

		if token != "" {
			// 将token作为userID设置到上下文中
			// 在测试中，我们直接使用用户ID作为token
			c.Set("userID", token)
			c.Set("mobile", token)
		}

		c.Next()
	}
}

// 创建测试地址请求
func (s *AddressHandlerTestSuite) createTestAddressRequest() AddShippingAddressRequest {
	return AddShippingAddressRequest{
		LinkMan:     "张三",
		Mobile:      "13800138000",
		Address:     "测试街道123号",
		IsDefault:   false,
		ProvinceId:  "440000",
		CityId:      "440300",
		DistrictId:  "440303",
	}
}

// 测试添加地址 - 成功场景
func (s *AddressHandlerTestSuite) TestAddAddress_Success() {
	req := s.createTestAddressRequest()
	reqJSON, _ := json.Marshal(req)

	// 创建请求
	request := httptest.NewRequest("POST", "/v1/user/shipping-address/add", bytes.NewBuffer(reqJSON))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+s.testUserID)

	// 执行请求
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应
	s.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("200", response["code"])
	s.Equal("Address added successfully", response["msg"])

	// 验证地址是否已创建
	addresses, err := s.addressRepo.ListByUserID(s.testUserID)
	s.Nil(err, "获取地址列表不应该失败")
	s.Len(addresses, 1, "应该有一个地址")
	s.Equal(req.LinkMan, addresses[0].LinkMan, "联系人姓名应该匹配")
	s.Equal(req.Mobile, addresses[0].Mobile, "手机号应该匹配")
}

// 测试添加地址 - 无效的JSON请求
func (s *AddressHandlerTestSuite) TestAddAddress_InvalidJSON() {
	// 创建无效的JSON请求
	invalidJSON := `{"linkMan": "张三", "mobile": "13800138000", "address": "测试地址", "provinceId": "440000", "cityId": "440300", "districtId": "440303"`

	request := httptest.NewRequest("POST", "/v1/user/shipping-address/add", bytes.NewBufferString(invalidJSON))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+s.testUserID)

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应
	s.Equal(http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("400", response["code"])
	s.Equal("Invalid request parameters", response["msg"])
}

// 测试添加地址 - 缺少必填字段
func (s *AddressHandlerTestSuite) TestAddAddress_MissingRequiredFields() {
	// 创建缺少必填字段的请求
	req := AddShippingAddressRequest{
		LinkMan:     "", // 缺少联系人
		Mobile:      "13800138000",
		Address:     "测试地址",
		ProvinceId:  "440000",
		CityId:      "440300",
		DistrictId:  "440303",
	}
	reqJSON, _ := json.Marshal(req)

	request := httptest.NewRequest("POST", "/v1/user/shipping-address/add", bytes.NewBuffer(reqJSON))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+s.testUserID)

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应 - 由于Gin的binding验证，缺少必填字段会导致绑定失败
	s.Equal(http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("400", response["code"])
	s.Equal("Invalid request parameters", response["msg"])
}

// 测试添加地址 - 用户未认证
func (s *AddressHandlerTestSuite) TestAddAddress_Unauthorized() {
	req := s.createTestAddressRequest()
	reqJSON, _ := json.Marshal(req)

	// 创建没有认证token的请求
	request := httptest.NewRequest("POST", "/v1/user/shipping-address/add", bytes.NewBuffer(reqJSON))
	request.Header.Set("Content-Type", "application/json")
	// 不设置Authorization header

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应
	s.Equal(http.StatusUnauthorized, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("401", response["code"])
	s.Equal("User not authenticated", response["msg"])
}

// 测试添加地址 - 默认地址
func (s *AddressHandlerTestSuite) TestAddAddress_DefaultAddress() {
	req := s.createTestAddressRequest()
	req.IsDefault = true
	reqJSON, _ := json.Marshal(req)

	request := httptest.NewRequest("POST", "/v1/user/shipping-address/add", bytes.NewBuffer(reqJSON))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+s.testUserID)

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应
	s.Equal(http.StatusOK, w.Code)

	// 验证地址是否已创建为默认地址
	addresses, err := s.addressRepo.ListByUserID(s.testUserID)
	s.Nil(err, "获取地址列表不应该失败")
	s.Len(addresses, 1, "应该有一个地址")
	s.Equal(1, addresses[0].IsDefault, "地址应该是默认地址")
}

// 测试获取地址列表 - 成功场景
func (s *AddressHandlerTestSuite) TestGetAddressList_Success() {
	// 先添加一些测试地址
	req1 := s.createTestAddressRequest()
	req1.LinkMan = "张三"
	s.addressService.AddAddress(s.testUserID, req1)

	req2 := s.createTestAddressRequest()
	req2.LinkMan = "李四"
	s.addressService.AddAddress(s.testUserID, req2)

	// 创建请求
	request := httptest.NewRequest("GET", "/v1/user/shipping-address/list", nil)
	request.Header.Set("Authorization", "Bearer "+s.testUserID)

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应
	s.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("200", response["code"])
	s.Equal("Address list retrieved successfully", response["msg"])

	// 验证数据
	data, ok := response["data"].([]interface{})
	s.True(ok, "data字段应该是数组")
	s.Len(data, 2, "应该有两个地址")
}

// 测试获取地址列表 - 用户未认证
func (s *AddressHandlerTestSuite) TestGetAddressList_Unauthorized() {
	request := httptest.NewRequest("GET", "/v1/user/shipping-address/list", nil)
	// 不设置Authorization header

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应
	s.Equal(http.StatusUnauthorized, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("401", response["code"])
	s.Equal("User not authenticated", response["msg"])
}

// 测试获取地址列表 - 空列表
func (s *AddressHandlerTestSuite) TestGetAddressList_EmptyList() {
	request := httptest.NewRequest("GET", "/v1/user/shipping-address/list", nil)
	request.Header.Set("Authorization", "Bearer "+s.testUserID)

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应
	s.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("200", response["code"])
	s.Equal("Address list retrieved successfully", response["msg"])

	// 验证数据为空
	data, ok := response["data"].([]interface{})
	s.True(ok, "data字段应该是数组")
	s.Len(data, 0, "应该返回空列表")
}

// 测试获取默认地址 - 成功场景
func (s *AddressHandlerTestSuite) TestGetDefaultAddress_Success() {
	// 先添加一个默认地址
	req := s.createTestAddressRequest()
	req.IsDefault = true
	s.addressService.AddAddress(s.testUserID, req)

	request := httptest.NewRequest("GET", "/v1/user/shipping-address/default", nil)
	request.Header.Set("Authorization", "Bearer "+s.testUserID)

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应
	s.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("200", response["code"])
	s.Equal("Default address retrieved successfully", response["msg"])

	// 验证数据
	data, ok := response["data"].(map[string]interface{})
	s.True(ok, "data字段应该是对象")
	s.Equal("张三", data["linkMan"])
	s.Equal("13800138000", data["mobile"])
	s.Equal(float64(1), data["isDefault"]) // JSON中的数字会被解析为float64
}

// 测试获取默认地址 - 用户未认证
func (s *AddressHandlerTestSuite) TestGetDefaultAddress_Unauthorized() {
	request := httptest.NewRequest("GET", "/v1/user/shipping-address/default", nil)
	// 不设置Authorization header

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应
	s.Equal(http.StatusUnauthorized, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("401", response["code"])
	s.Equal("User not authenticated", response["msg"])
}

// 测试获取默认地址 - 没有默认地址
func (s *AddressHandlerTestSuite) TestGetDefaultAddress_NoDefaultAddress() {
	// 添加一个非默认地址
	req := s.createTestAddressRequest()
	req.IsDefault = false
	s.addressService.AddAddress(s.testUserID, req)

	request := httptest.NewRequest("GET", "/v1/user/shipping-address/default", nil)
	request.Header.Set("Authorization", "Bearer "+s.testUserID)

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应
	s.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("200", response["code"])
	s.Equal("No default address found", response["msg"])
	s.Nil(response["data"], "data字段应该为null")
}

// 测试获取默认地址 - 完全空列表
func (s *AddressHandlerTestSuite) TestGetDefaultAddress_EmptyList() {
	request := httptest.NewRequest("GET", "/v1/user/shipping-address/default", nil)
	request.Header.Set("Authorization", "Bearer "+s.testUserID)

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应
	s.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("200", response["code"])
	s.Equal("No default address found", response["msg"])
	s.Nil(response["data"], "data字段应该为null")
}

// 测试添加地址 - 服务层错误
func (s *AddressHandlerTestSuite) TestAddAddress_ServiceError() {
	// 使用不存在的用户ID来触发服务层错误
	req := s.createTestAddressRequest()
	reqJSON, _ := json.Marshal(req)

	request := httptest.NewRequest("POST", "/v1/user/shipping-address/add", bytes.NewBuffer(reqJSON))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer nonexistent_user_id")

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应 - 由于用户不存在，服务层会返回500
	s.Equal(http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("500", response["code"])
	s.Equal("Failed to add address", response["msg"])
}

// 测试获取地址列表 - 服务层错误
func (s *AddressHandlerTestSuite) TestGetAddressList_ServiceError() {
	// 使用不存在的用户ID来测试
	request := httptest.NewRequest("GET", "/v1/user/shipping-address/list", nil)
	request.Header.Set("Authorization", "Bearer nonexistent_user_id")

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应 - 由于用户不存在，会返回空列表
	s.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("200", response["code"])
	s.Equal("Address list retrieved successfully", response["msg"])

	// 验证返回的是空列表
	data, ok := response["data"].([]interface{})
	s.True(ok, "data字段应该是数组")
	s.Empty(data, "应该返回空列表")
}

// 测试获取默认地址 - 服务层错误
func (s *AddressHandlerTestSuite) TestGetDefaultAddress_ServiceError() {
	// 使用不存在的用户ID来测试
	request := httptest.NewRequest("GET", "/v1/user/shipping-address/default", nil)
	request.Header.Set("Authorization", "Bearer nonexistent_user_id")

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, request)

	// 验证响应 - 由于用户不存在，会返回没有默认地址的响应
	s.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.Nil(err, "响应应该是有效的JSON")
	s.Equal("200", response["code"])
	s.Equal("No default address found", response["msg"])

	// 验证返回的数据为空
	s.Nil(response["data"], "data字段应该为空")
}

// 测试请求参数绑定验证
func (s *AddressHandlerTestSuite) TestAddAddress_RequestValidation() {
	testCases := []struct {
		name        string
		request     AddShippingAddressRequest
		expectError bool
	}{
		{
			name: "有效请求",
			request: AddShippingAddressRequest{
				LinkMan:     "张三",
				Mobile:      "13800138000",
				Address:     "测试地址",
				ProvinceId:  "440000",
				CityId:      "440300",
				DistrictId:  "440303",
			},
			expectError: false,
		},
		{
			name: "缺少联系人",
			request: AddShippingAddressRequest{
				LinkMan:     "",
				Mobile:      "13800138000",
				Address:     "测试地址",
				ProvinceId:  "440000",
				CityId:      "440300",
				DistrictId:  "440303",
			},
			expectError: true,
		},
		{
			name: "缺少手机号",
			request: AddShippingAddressRequest{
				LinkMan:     "张三",
				Mobile:      "",
				Address:     "测试地址",
				ProvinceId:  "440000",
				CityId:      "440300",
				DistrictId:  "440303",
			},
			expectError: true,
		},
		{
			name: "缺少地址",
			request: AddShippingAddressRequest{
				LinkMan:     "张三",
				Mobile:      "13800138000",
				Address:     "",
				ProvinceId:  "440000",
				CityId:      "440300",
				DistrictId:  "440303",
			},
			expectError: true,
		},
		{
			name: "缺少省份ID",
			request: AddShippingAddressRequest{
				LinkMan:     "张三",
				Mobile:      "13800138000",
				Address:     "测试地址",
				ProvinceId:  "",
				CityId:      "440300",
				DistrictId:  "440303",
			},
			expectError: true,
		},
		{
			name: "缺少城市ID",
			request: AddShippingAddressRequest{
				LinkMan:     "张三",
				Mobile:      "13800138000",
				Address:     "测试地址",
				ProvinceId:  "440000",
				CityId:      "",
				DistrictId:  "440303",
			},
			expectError: true,
		},
		{
			name: "缺少区县ID",
			request: AddShippingAddressRequest{
				LinkMan:     "张三",
				Mobile:      "13800138000",
				Address:     "测试地址",
				ProvinceId:  "440000",
				CityId:      "440300",
				DistrictId:  "",
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			reqJSON, _ := json.Marshal(tc.request)

			request := httptest.NewRequest("POST", "/v1/user/shipping-address/add", bytes.NewBuffer(reqJSON))
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("Authorization", "Bearer "+s.testUserID)

			w := httptest.NewRecorder()
			s.router.ServeHTTP(w, request)

			if tc.expectError {
				s.Equal(http.StatusBadRequest, w.Code)
			} else {
				s.Equal(http.StatusOK, w.Code)
			}
		})
	}
}

// 测试Handler构造函数
func (s *AddressHandlerTestSuite) TestNewAddressHandler() {
	handler := NewAddressHandler(s.addressService)
	s.NotNil(handler, "Handler不应该为nil")
	s.Equal(s.addressService, handler.addressService, "addressService应该正确设置")
}

// 测试请求结构体标签
func (s *AddressHandlerTestSuite) TestAddShippingAddressRequestTags() {
	req := AddShippingAddressRequest{}

	// 验证结构体字段存在
	s.Empty(req.LinkMan, "LinkMan字段应该存在")
	s.Empty(req.Mobile, "Mobile字段应该存在")
	s.Empty(req.Address, "Address字段应该存在")
	s.False(req.IsDefault, "IsDefault字段应该存在且默认为false")
	s.Empty(req.ProvinceId, "ProvinceId字段应该存在")
	s.Empty(req.CityId, "CityId字段应该存在")
	s.Empty(req.DistrictId, "DistrictId字段应该存在")
}