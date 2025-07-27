package addresses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"bookstore/app/common/models"
	"bookstore/app/configs"
	"bookstore/app/testutils"
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
	suite.SupperSuite.SetupSuite()
	gin.SetMode(gin.TestMode)

	db := configs.Cfg.DBConnection()
	suite.repo = NewAddressRepositoryDB(db)
	suite.service = NewAddressService(suite.repo, db)
	suite.handler = NewAddressHandler(suite.service)

	suite.router = gin.Default()
	suite.router.Use(func(c *gin.Context) {
		c.Set("userID", "test_user_id_1")
		c.Next()
	})
	suite.router.POST("/v1/user/shipping-address/add", suite.handler.AddAddress)
}

func (suite *AddressHandlerSuite) SetupTest() {
	// 清理测试数据
	addresses, _ := suite.repo.ListByUserID("test_user_id_1")
	for _, addr := range addresses {
		suite.repo.Delete(addr.Id)
	}
}

func (suite *AddressHandlerSuite) TestSuccessfulAddition() {
	reqBody := AddShippingAddressRequest{
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

	suite.Equal(http.StatusUnauthorized, w.Code)
	var resp models.JsonResult
	json.Unmarshal(w.Body.Bytes(), &resp)
	suite.Equal("401", resp.Code)
}

func (suite *AddressHandlerSuite) TestIsDefaultLogicNewDefault() {
	reqBody := AddShippingAddressRequest{
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

	// Verify that the address was created as default
	addresses, err := suite.repo.ListByUserID("test_user_id_1")
	suite.NoError(err)
	suite.Len(addresses, 1)
	suite.Equal(1, addresses[0].IsDefault)
}

func (suite *AddressHandlerSuite) TestIsDefaultLogicAddNonDefault() {
	// First add a default address
	reqBody1 := AddShippingAddressRequest{
		LinkMan:       "Default User",
		Mobile:        "12345678901",
		ProvinceStr:   "Province",
		CityStr:       "City",
		AreaStr:       "Area",
		DetailAddress: "Detail",
		IsDefault:     1,
	}
	jsonBody1, _ := json.Marshal(reqBody1)

	req1, _ := http.NewRequest(http.MethodPost, "/v1/user/shipping-address/add", bytes.NewBuffer(jsonBody1))
	req1.Header.Set("Content-Type", "application/json")

	w1 := httptest.NewRecorder()
	suite.router.ServeHTTP(w1, req1)
	suite.Equal(http.StatusOK, w1.Code)

	// Then add a non-default address
	reqBody2 := AddShippingAddressRequest{
		LinkMan:       "Non-Default User",
		Mobile:        "12345678902",
		ProvinceStr:   "Province",
		CityStr:       "City",
		AreaStr:       "Area",
		DetailAddress: "Detail2",
		IsDefault:     0,
	}
	jsonBody2, _ := json.Marshal(reqBody2)

	req2, _ := http.NewRequest(http.MethodPost, "/v1/user/shipping-address/add", bytes.NewBuffer(jsonBody2))
	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()
	suite.router.ServeHTTP(w2, req2)
	suite.Equal(http.StatusOK, w2.Code)

	// Verify that we have two addresses, one default and one non-default
	addresses, err := suite.repo.ListByUserID("test_user_id_1")
	suite.NoError(err)
	suite.Len(addresses, 2)

	defaultCount := 0
	for _, addr := range addresses {
		if addr.IsDefault == 1 {
			defaultCount++
		}
	}
	suite.Equal(1, defaultCount)
}