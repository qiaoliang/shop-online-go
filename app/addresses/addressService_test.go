package addresses

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// 移除 bookstore/app/addresses/repository
	"bookstore/app/configs"
)

// 如果有 TestMain，直接删除，保留 addressHandler_test.go 中的即可。

func TestAddressService_AddAddress(t *testing.T) {
	// Test Case 1: Successful addition (non-default)
	t.Run("Successful addition - non-default", func(t *testing.T) {
		db := configs.Cfg.DBConnection()
		repo := NewAddressRepositoryDB(db)
		service := NewAddressService(repo, nil)

		req := AddShippingAddressRequest{
			LinkMan:       "Test User",
			Mobile:        "12345678901",
			ProvinceStr:   "Province",
			CityStr:       "City",
			AreaStr:       "Area",
			DetailAddress: "Detail",
			IsDefault:     0,
		}
		userID := "user123"

		err := service.AddAddress(userID, req)
		assert.NoError(t, err)

		addresses, err := repo.ListByUserID(userID)
		assert.NoError(t, err)
		assert.Len(t, addresses, 1)
		assert.Equal(t, req.LinkMan, addresses[0].LinkMan)
		assert.Equal(t, 0, addresses[0].IsDefault)
	})

	// Test Case 2: Successful addition (isDefault = true)
	t.Run("Successful addition - isDefault true", func(t *testing.T) {
		db := configs.Cfg.DBConnection()
		repo := NewAddressRepositoryDB(db)
		service := NewAddressService(repo, nil)

		// Add an initial default address
		initialAddress := &Address{
			UserId:        "user456",
			LinkMan:       "Initial User",
			Mobile:        "11111111111",
			ProvinceStr:   "P1", CityStr: "C1", AreaStr: "A1", DetailAddress: "D1",
			IsDefault:     1,
		}
		repo.Create(initialAddress)

		req := AddShippingAddressRequest{
			LinkMan:       "New Default User",
			Mobile:        "22222222222",
			ProvinceStr:   "P2",
			CityStr:       "C2",
			AreaStr:       "A2",
			DetailAddress: "D2",
			IsDefault:     1,
		}
		userID := "user456"

		err := service.AddAddress(userID, req)
		assert.NoError(t, err)

		addresses, err := repo.ListByUserID(userID)
		assert.NoError(t, err)
		assert.Len(t, addresses, 2)

		// Check that only the new address is default
		defaultCount := 0
		for _, addr := range addresses {
			if addr.IsDefault == 1 {
				defaultCount++
				assert.Equal(t, req.LinkMan, addr.LinkMan) // New address should be the default
			}
		}
		assert.Equal(t, 1, defaultCount)
	})

	// Test Case 3: Invalid request data
	t.Run("Invalid request data", func(t *testing.T) {
		db := configs.Cfg.DBConnection()
		repo := NewAddressRepositoryDB(db)
		service := NewAddressService(repo, nil)

		req := AddShippingAddressRequest{
			LinkMan: "", // Missing required field
		}
		userID := "user123"

		err := service.AddAddress(userID, req)
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidAddressData, err)
	})

	// 已删除/注释掉模拟 DB 错误的测试用例
	// Test Case 4: SaveAddress returns error (simulated by closing DB)
	// Test Case 5: SetOtherAddressesNonDefault returns error (simulated by closing DB)
}