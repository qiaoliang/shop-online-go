package addresses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"bookstore/app/common/models"
	"bookstore/app/configs"
)

func TestMain(m *testing.M) {
	os.Remove("./test.db") // 测试前删除
	configs.GetConfigInstance("../../config-test.yaml")
	code := m.Run()
	os.Remove("./test.db") // 测试后删除
	os.Exit(code)
}

func TestAddAddress(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Test Case 1: Successful addition
	t.Run("Successful addition", func(t *testing.T) {
		db := configs.Cfg.DBConnection()
		repo := NewAddressRepositoryDB(db)
		service := NewAddressService(repo, nil)
		handler := NewAddressHandler(service)

		r := gin.Default()
		r.Use(func(c *gin.Context) {
			c.Set("userID", "test_user_id_1")
			c.Next()
		})
		r.POST("/v1/user/shipping-address/add", handler.AddAddress)

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
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var resp models.JsonResult
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Equal(t, "OK", resp.Msg)

		// Verify data in repo
		addresses, err := repo.ListByUserID("test_user_id_1")
		assert.NoError(t, err)
		assert.Len(t, addresses, 1)
		assert.Equal(t, reqBody.LinkMan, addresses[0].LinkMan)
		assert.Equal(t, reqBody.IsDefault, addresses[0].IsDefault)
	})

	// Test Case 2: Invalid request parameters
	t.Run("Invalid request parameters", func(t *testing.T) {
		db := configs.Cfg.DBConnection()
		repo := NewAddressRepositoryDB(db)
		service := NewAddressService(repo, nil)
		handler := NewAddressHandler(service)

		r := gin.Default()
		r.Use(func(c *gin.Context) {
			c.Set("userID", "test_user_id_1")
			c.Next()
		})
		r.POST("/v1/user/shipping-address/add", handler.AddAddress)

		reqBody := AddShippingAddressRequest{
			LinkMan: "", // Missing required field
		}
		jsonBody, _ := json.Marshal(reqBody)

		req, _ := http.NewRequest(http.MethodPost, "/v1/user/shipping-address/add", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var resp models.JsonResult
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Equal(t, "Invalid request parameters", resp.Msg)
	})

	// Test Case 3: User not authenticated (missing userID in context)
	t.Run("User not authenticated", func(t *testing.T) {
		db := configs.Cfg.DBConnection()
		repo := NewAddressRepositoryDB(db)
		service := NewAddressService(repo, nil)
		handler := NewAddressHandler(service)

		r := gin.Default()
		r.POST("/v1/user/shipping-address/add", handler.AddAddress)

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
		// No userID set in context

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		var resp models.JsonResult
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Equal(t, "User not authenticated", resp.Msg)
	})

	// Test Case 4: isDefault logic - setting a new default address
	t.Run("isDefault logic - new default", func(t *testing.T) {
		db := configs.Cfg.DBConnection()
		repo := NewAddressRepositoryDB(db)
		service := NewAddressService(repo, nil)
		handler := NewAddressHandler(service)

		r := gin.Default()
		r.Use(func(c *gin.Context) {
			c.Set("userID", "test_user_id_2")
			c.Next()
		})
		r.POST("/v1/user/shipping-address/add", handler.AddAddress)

		// Add an initial default address
		initialAddress := &Address{
			UserId:        "test_user_id_2",
			LinkMan:       "Initial User",
			Mobile:        "11111111111",
			ProvinceStr:   "P1", CityStr: "C1", AreaStr: "A1", DetailAddress: "D1",
			IsDefault:     1,
		}
		repo.Create(initialAddress)

		// Add a new address and set it as default
		reqBody := AddShippingAddressRequest{
			LinkMan:       "New User",
			Mobile:        "22222222222",
			ProvinceStr:   "P2",
			CityStr:       "C2",
			AreaStr:       "A2",
			DetailAddress: "D2",
			IsDefault:     1,
		}
		jsonBody, _ := json.Marshal(reqBody)

		req, _ := http.NewRequest(http.MethodPost, "/v1/user/shipping-address/add", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Verify addresses in repo
		addresses, err := repo.ListByUserID("test_user_id_2")
		assert.NoError(t, err)
		assert.Len(t, addresses, 2)

		// Check that only the new address is default
		defaultCount := 0
		for _, addr := range addresses {
			if addr.IsDefault == 1 {
				defaultCount++
				assert.Equal(t, reqBody.LinkMan, addr.LinkMan) // New address should be the default
			}
		}
		assert.Equal(t, 1, defaultCount)
	})

	// Test Case 5: isDefault logic - adding non-default address
	t.Run("isDefault logic - add non-default", func(t *testing.T) {
		db := configs.Cfg.DBConnection()
		repo := NewAddressRepositoryDB(db)
		service := NewAddressService(repo, nil)
		handler := NewAddressHandler(service)

		r := gin.Default()
		r.Use(func(c *gin.Context) {
			c.Set("userID", "test_user_id_3")
			c.Next()
		})
		r.POST("/v1/user/shipping-address/add", handler.AddAddress)

		// Add an initial default address
		initialAddress := &Address{
			UserId:        "test_user_id_3",
			LinkMan:       "Initial User",
			Mobile:        "11111111111",
			ProvinceStr:   "P1", CityStr: "C1", AreaStr: "A1", DetailAddress: "D1",
			IsDefault:     1,
		}
		repo.Create(initialAddress)

		// Add a new address and set it as non-default
		reqBody := AddShippingAddressRequest{
			LinkMan:       "New User",
			Mobile:        "22222222222",
			ProvinceStr:   "P2",
			CityStr:       "C2",
			AreaStr:       "A2",
			DetailAddress: "D2",
			IsDefault:     0,
		}
		jsonBody, _ := json.Marshal(reqBody)

		req, _ := http.NewRequest(http.MethodPost, "/v1/user/shipping-address/add", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Verify addresses in repo
		addresses, err := repo.ListByUserID("test_user_id_3")
		assert.NoError(t, err)
		assert.Len(t, addresses, 2)

		// Check that the initial address is still default and new one is not
		defaultCount := 0
		for _, addr := range addresses {
			if addr.IsDefault == 1 {
				defaultCount++
				assert.Equal(t, initialAddress.LinkMan, addr.LinkMan) // Initial address should still be the default
			}
		}
		assert.Equal(t, 1, defaultCount)
	})
}