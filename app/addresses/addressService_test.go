package addresses

import (
	"testing"

	"bookstore/app/configs"
	"bookstore/app/testutils"

	"github.com/stretchr/testify/suite"
)

type AddressServiceSuite struct {
	testutils.SupperSuite
	repo    AddressRepository
	service AddressService
}

func TestAddressServiceSuite(t *testing.T) {
	suite.Run(t, new(AddressServiceSuite))
}

func (suite *AddressServiceSuite) SetupSuite() {
	suite.SupperSuite.SetupSuite()
	db := configs.Cfg.DBConnection()
	suite.repo = NewAddressRepositoryDB(db)
	suite.service = NewAddressService(suite.repo, db)
}

func (suite *AddressServiceSuite) SetupTest() {
	// 清理测试数据
	addresses, _ := suite.repo.ListByUserID("user123")
	for _, addr := range addresses {
		suite.repo.Delete(addr.Id)
	}
	addresses, _ = suite.repo.ListByUserID("user456")
	for _, addr := range addresses {
		suite.repo.Delete(addr.Id)
	}
}

func (suite *AddressServiceSuite) TestSuccessfulAdditionNonDefault() {
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

	err := suite.service.AddAddress(userID, req)
	suite.NoError(err)

	addresses, err := suite.repo.ListByUserID(userID)
	suite.NoError(err)
	suite.Len(addresses, 1)
	suite.Equal(req.LinkMan, addresses[0].LinkMan)
	suite.Equal(0, addresses[0].IsDefault)
}

func (suite *AddressServiceSuite) TestSuccessfulAdditionIsDefaultTrue() {
	// Add an initial default address
	initialAddress := &Address{
		UserId:        "user456",
		LinkMan:       "Initial User",
		Mobile:        "11111111111",
		ProvinceStr:   "P1", CityStr: "C1", AreaStr: "A1", DetailAddress: "D1",
		IsDefault:     1,
	}
	suite.repo.Create(initialAddress)

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

	err := suite.service.AddAddress(userID, req)
	suite.NoError(err)

	addresses, err := suite.repo.ListByUserID(userID)
	suite.NoError(err)
	suite.Len(addresses, 2)

	// Check that only the new address is default
	defaultCount := 0
	for _, addr := range addresses {
		if addr.IsDefault == 1 {
			defaultCount++
			suite.Equal(req.LinkMan, addr.LinkMan) // New address should be the default
		}
	}
	suite.Equal(1, defaultCount)
}

func (suite *AddressServiceSuite) TestInvalidRequestData() {
	req := AddShippingAddressRequest{
		LinkMan: "", // Missing required field
	}
	userID := "user123"

	err := suite.service.AddAddress(userID, req)
	suite.Error(err)
	suite.Equal(ErrInvalidAddressData, err)
}