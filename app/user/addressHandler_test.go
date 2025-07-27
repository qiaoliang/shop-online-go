package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bookstore/app/common/models"
	"bookstore/app/configs"
	"bookstore/app/testutils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AddressHandlerSuite struct {
	testutils.SupperSuite
	router *gin.Engine
	repo   AddressRepository
	service AddressService
	handler *AddressHandler
}

func TestAddressHandlerSuite(t *testing.T) {
	suite.Run(t, new(AddressHandlerSuite))
}

func (suite *AddressHandlerSuite) SetupSuite() {
	// 初始化配置
	configs.GetConfigInstance("config-test.yaml")

	// 创建数据库连接
	db := configs.Cfg.DBConnection()

	// 创建仓库和服务
	suite.repo = NewAddressRepositoryDB(db)

	// 创建用户仓库
	userRepo := NewUserRepoDB(db)

	suite.service = NewAddressService(suite.repo, userRepo, db)
	suite.handler = NewAddressHandler(suite.service)
	suite.router = setupTestRouter(suite.handler)
}

func (suite *AddressHandlerSuite) SetupTest() {
	// 清理测试数据
	db := configs.Cfg.DBConnection()
	db.Exec("DELETE FROM addresses")
}

func setupTestRouter(handler *AddressHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// 添加测试专用的认证中间件（仅用于GET接口）
	router.Use(testAuthMiddleware())

	v1 := router.Group("/v1")
	v1.POST("/user/shipping-address/add", handler.AddAddress) // 不需要认证中间件
	v1.GET("/user/shipping-address/list", handler.GetAddressList)
	v1.GET("/user/shipping-address/default", handler.GetDefaultAddress)

	return router
}

// 测试专用的认证中间件
func testAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 为测试设置固定的用户ID
		c.Set("userID", "test_user_id_1")
		c.Set("mobile", "test_mobile_1")
		c.Next()
	}
}

func (suite *AddressHandlerSuite) TestSuccessfulAddition() {
	reqBody := AddShippingAddressRequest{
		Token:         "test_user_id_1",
		LinkMan:       "Test User",
		Mobile:        "12345678901",
		ProvinceStr:   "Province",
		CityStr:       "City",
		AreaStr:       "Area",
		DetailAddress: "Detail",
		IsDefault:     1,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(http.MethodPost, "/v1/user/shipping-address/add", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	var resp models.JsonResult
	json.Unmarshal(w.Body.Bytes(), &resp)
	suite.Equal("Address added successfully", resp.Msg)

	// Verify data in repo
	addresses, err := suite.repo.ListByUserID("test_user_id_1")
	suite.NoError(err)
	suite.Len(addresses, 1)
	suite.Equal(reqBody.LinkMan, addresses[0].LinkMan)
	suite.Equal(reqBody.IsDefault, addresses[0].IsDefault)
}

func (suite *AddressHandlerSuite) TestInvalidRequestParameters() {
	reqBody := AddShippingAddressRequest{
		Token:   "test_user_id_1",
		LinkMan: "", // Missing required field
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(http.MethodPost, "/v1/user/shipping-address/add", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
	var resp models.JsonResult
	json.Unmarshal(w.Body.Bytes(), &resp)
	suite.Equal("400", resp.Code)
}

func (suite *AddressHandlerSuite) TestUserNotAuthenticated() {
	// 创建一个没有用户认证的router
	router := gin.Default()
	router.POST("/v1/user/shipping-address/add", suite.handler.AddAddress)

	reqBody := AddShippingAddressRequest{
		Token:         "", // Empty token
		LinkMan:       "Test User",
		Mobile:        "12345678901",
		ProvinceStr:   "Province",
		CityStr:       "City",
		AreaStr:       "Area",
		DetailAddress: "Detail",
		IsDefault:     1,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(http.MethodPost, "/v1/user/shipping-address/add", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
	var resp models.JsonResult
	json.Unmarshal(w.Body.Bytes(), &resp)
	suite.Equal("400", resp.Code)
}

// 新增测试：获取地址列表
func (suite *AddressHandlerSuite) TestGetAddressList() {
	// 先添加一些测试地址
	address1 := &Address{
		UserId:        "test_user_id_1",
		LinkMan:       "User 1",
		Mobile:        "12345678901",
		ProvinceStr:   "Province1",
		CityStr:       "City1",
		AreaStr:       "Area1",
		DetailAddress: "Detail1",
		IsDefault:     1,
	}
	suite.repo.Create(address1)

	address2 := &Address{
		UserId:        "test_user_id_1",
		LinkMan:       "User 2",
		Mobile:        "12345678902",
		ProvinceStr:   "Province2",
		CityStr:       "City2",
		AreaStr:       "Area2",
		DetailAddress: "Detail2",
		IsDefault:     0,
	}
	suite.repo.Create(address2)

	// 测试获取地址列表
	req, _ := http.NewRequest(http.MethodGet, "/v1/user/shipping-address/list", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	suite.Equal("200", response["code"])
	suite.Equal("Address list retrieved successfully", response["msg"])

	// 验证返回的数据
	data := response["data"].([]interface{})
	suite.Len(data, 2)
}

// 新增测试：获取默认地址
func (suite *AddressHandlerSuite) TestGetDefaultAddress() {
	// 先添加一个默认地址
	address := &Address{
		UserId:        "test_user_id_1",
		LinkMan:       "Default User",
		Mobile:        "12345678901",
		ProvinceStr:   "Province",
		CityStr:       "City",
		AreaStr:       "Area",
		DetailAddress: "Detail",
		IsDefault:     1,
	}
	suite.repo.Create(address)

	// 测试获取默认地址
	req, _ := http.NewRequest(http.MethodGet, "/v1/user/shipping-address/default", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	suite.Equal("200", response["code"])
	suite.Equal("Default address retrieved successfully", response["msg"])

	// 验证返回的数据
	data := response["data"].(map[string]interface{})
	suite.Equal("Default User", data["linkMan"])
	suite.Equal(float64(1), data["isDefault"])
}

// 新增测试：没有默认地址的情况
func (suite *AddressHandlerSuite) TestGetDefaultAddressWhenNone() {
	// 测试获取默认地址（没有默认地址）
	req, _ := http.NewRequest(http.MethodGet, "/v1/user/shipping-address/default", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	suite.Equal("200", response["code"])
	suite.Equal("No default address found", response["msg"])
	suite.Nil(response["data"])
}

// 新增测试：空地址列表
func (suite *AddressHandlerSuite) TestGetAddressListEmpty() {
	// 测试获取地址列表（空列表）
	req, _ := http.NewRequest(http.MethodGet, "/v1/user/shipping-address/list", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	suite.Equal("200", response["code"])
	suite.Equal("Address list retrieved successfully", response["msg"])

	// 验证返回的数据
	data := response["data"].([]interface{})
	suite.Len(data, 0)
}