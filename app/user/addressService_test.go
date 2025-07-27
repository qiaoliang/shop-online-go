package user

import (
	"testing"

	"bookstore/app/configs"
	"bookstore/app/testutils"
	"bookstore/app/utils"

	"gorm.io/gorm"

	"github.com/stretchr/testify/suite"
)

type AddressServiceTestSuite struct {
	testutils.SupperSuite
	addressService AddressService
	addressRepo    AddressRepository
	userRepo       UserRepo
	db             *gorm.DB
	testUserID     string
}

func TestAddressServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AddressServiceTestSuite))
}

func (s *AddressServiceTestSuite) SetupSuite() {
	s.SupperSuite.SetupSuite()
	s.db = configs.Cfg.DBConnection()
	s.addressRepo = NewAddressRepositoryDB(s.db)
	s.userRepo = NewUserRepoDB(s.db)
	s.addressService = NewAddressService(s.addressRepo, s.userRepo, s.db)

	// 创建测试用户
	s.createTestUser()
}

func (s *AddressServiceTestSuite) TeardownSuite() {
	s.SupperSuite.TeardownSuite()
	s.addressService = nil
	s.addressRepo = nil
	s.userRepo = nil
}

func (s *AddressServiceTestSuite) SetupTest() {
	// 清理测试数据
	s.cleanupTestData()
}

func (s *AddressServiceTestSuite) TearDownTest() {
	// 清理测试数据
	s.cleanupTestData()
}

// 创建测试用户
func (s *AddressServiceTestSuite) createTestUser() {
	mobile := "13900001234" + utils.RandomImpl{}.GenStr()
	user, err := s.userRepo.CreateUser(mobile, "password123", "测试用户", "0", genUId)
	s.Nil(err, "创建测试用户失败")
	s.NotNil(user, "用户创建失败")
	s.testUserID = user.Id
}

// 清理测试数据
func (s *AddressServiceTestSuite) cleanupTestData() {
	// 删除测试用户的所有地址
	if s.testUserID != "" {
		addresses, _ := s.addressRepo.ListByUserID(s.testUserID)
		for _, addr := range addresses {
			s.addressRepo.Delete(addr.Id)
		}
	}
}

// 创建测试地址请求
func (s *AddressServiceTestSuite) createTestAddressRequest() AddShippingAddressRequest {
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
func (s *AddressServiceTestSuite) TestAddAddress_Success() {
	req := s.createTestAddressRequest()

	err := s.addressService.AddAddress(s.testUserID, req)
	s.Nil(err, "添加地址应该成功")

	// 验证地址是否已创建
	addresses, err := s.addressRepo.ListByUserID(s.testUserID)
	s.Nil(err, "获取地址列表不应该失败")
	s.Len(addresses, 1, "应该有一个地址")
	s.Equal(req.LinkMan, addresses[0].LinkMan, "联系人姓名应该匹配")
	s.Equal(req.Mobile, addresses[0].Mobile, "手机号应该匹配")
	s.Equal(req.Address, addresses[0].DetailAddress, "详细地址应该匹配")
	s.Equal(req.ProvinceId, addresses[0].ProvinceStr, "省份ID应该匹配")
	s.Equal(req.CityId, addresses[0].CityStr, "城市ID应该匹配")
	s.Equal(req.DistrictId, addresses[0].AreaStr, "区县ID应该匹配")
	s.Equal(0, addresses[0].IsDefault, "非默认地址的IsDefault应该为0")
}

// 测试添加默认地址 - 成功场景
func (s *AddressServiceTestSuite) TestAddAddress_DefaultAddress_Success() {
	req := s.createTestAddressRequest()
	req.IsDefault = true

	err := s.addressService.AddAddress(s.testUserID, req)
	s.Nil(err, "添加默认地址应该成功")

	// 验证地址是否已创建为默认地址
	addresses, err := s.addressRepo.ListByUserID(s.testUserID)
	s.Nil(err, "获取地址列表不应该失败")
	s.Len(addresses, 1, "应该有一个地址")
	s.Equal(1, addresses[0].IsDefault, "默认地址的IsDefault应该为1")
}

// 测试添加多个地址时设置默认地址的逻辑
func (s *AddressServiceTestSuite) TestAddAddress_MultipleAddresses_DefaultLogic() {
	// 添加第一个地址（非默认）
	req1 := s.createTestAddressRequest()
	req1.LinkMan = "张三"
	req1.IsDefault = false

	err := s.addressService.AddAddress(s.testUserID, req1)
	s.Nil(err, "添加第一个地址应该成功")

	// 添加第二个地址（默认）
	req2 := s.createTestAddressRequest()
	req2.LinkMan = "李四"
	req2.IsDefault = true

	err = s.addressService.AddAddress(s.testUserID, req2)
	s.Nil(err, "添加第二个地址应该成功")

	// 验证地址列表
	addresses, err := s.addressRepo.ListByUserID(s.testUserID)
	s.Nil(err, "获取地址列表不应该失败")
	s.Len(addresses, 2, "应该有两个地址")

	// 验证默认地址逻辑：第一个地址应该变为非默认，第二个地址应该是默认
	var defaultAddress *Address
	var nonDefaultAddress *Address
	for _, addr := range addresses {
		if addr.LinkMan == "张三" {
			nonDefaultAddress = addr
		} else if addr.LinkMan == "李四" {
			defaultAddress = addr
		}
	}

	s.NotNil(nonDefaultAddress, "应该找到张三的地址")
	s.NotNil(defaultAddress, "应该找到李四的地址")
	s.Equal(0, nonDefaultAddress.IsDefault, "张三的地址应该不是默认地址")
	s.Equal(1, defaultAddress.IsDefault, "李四的地址应该是默认地址")
}

// 测试添加地址 - 无效用户ID
func (s *AddressServiceTestSuite) TestAddAddress_InvalidUserID() {
	req := s.createTestAddressRequest()

	err := s.addressService.AddAddress("", req)
	s.Equal(ErrInvalidAddressData, err, "空用户ID应该返回无效地址数据错误")
}

// 测试添加地址 - 用户不存在
func (s *AddressServiceTestSuite) TestAddAddress_UserNotFound() {
	req := s.createTestAddressRequest()

	err := s.addressService.AddAddress("nonexistent_user_id", req)
	s.Equal(ErrUserNotFound, err, "不存在的用户ID应该返回用户未找到错误")
}

// 测试添加地址 - 无效地址数据
func (s *AddressServiceTestSuite) TestAddAddress_InvalidAddressData() {
	testCases := []struct {
		name string
		req  AddShippingAddressRequest
	}{
		{
			name: "空联系人",
			req: AddShippingAddressRequest{
				LinkMan:     "",
				Mobile:      "13800138000",
				Address:     "测试地址",
				ProvinceId:  "440000",
				CityId:      "440300",
				DistrictId:  "440303",
			},
		},
		{
			name: "空手机号",
			req: AddShippingAddressRequest{
				LinkMan:     "张三",
				Mobile:      "",
				Address:     "测试地址",
				ProvinceId:  "440000",
				CityId:      "440300",
				DistrictId:  "440303",
			},
		},
		{
			name: "空详细地址",
			req: AddShippingAddressRequest{
				LinkMan:     "张三",
				Mobile:      "13800138000",
				Address:     "",
				ProvinceId:  "440000",
				CityId:      "440300",
				DistrictId:  "440303",
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			err := s.addressService.AddAddress(s.testUserID, tc.req)
			s.Equal(ErrInvalidAddressData, err, "无效地址数据应该返回错误")
		})
	}
}

// 测试获取地址列表 - 成功场景
func (s *AddressServiceTestSuite) TestGetAddressList_Success() {
	// 先添加一些测试地址
	req1 := s.createTestAddressRequest()
	req1.LinkMan = "张三"
	s.addressService.AddAddress(s.testUserID, req1)

	req2 := s.createTestAddressRequest()
	req2.LinkMan = "李四"
	s.addressService.AddAddress(s.testUserID, req2)

	// 获取地址列表
	addresses, err := s.addressService.GetAddressList(s.testUserID)
	s.Nil(err, "获取地址列表不应该失败")
	s.Len(addresses, 2, "应该有两个地址")

	// 验证地址内容
	names := make(map[string]bool)
	for _, addr := range addresses {
		names[addr.LinkMan] = true
	}
	s.True(names["张三"], "应该包含张三的地址")
	s.True(names["李四"], "应该包含李四的地址")
}

// 测试获取地址列表 - 空用户ID
func (s *AddressServiceTestSuite) TestGetAddressList_InvalidUserID() {
	addresses, err := s.addressService.GetAddressList("")
	s.Equal(ErrInvalidAddressData, err, "空用户ID应该返回无效地址数据错误")
	s.Nil(addresses, "地址列表应该为nil")
}

// 测试获取地址列表 - 用户没有地址
func (s *AddressServiceTestSuite) TestGetAddressList_NoAddresses() {
	addresses, err := s.addressService.GetAddressList(s.testUserID)
	s.Nil(err, "获取空地址列表不应该失败")
	s.Len(addresses, 0, "应该返回空列表")
}

// 测试获取默认地址 - 成功场景
func (s *AddressServiceTestSuite) TestGetDefaultAddress_Success() {
	// 添加一个默认地址
	req := s.createTestAddressRequest()
	req.IsDefault = true
	s.addressService.AddAddress(s.testUserID, req)

	// 获取默认地址
	defaultAddress, err := s.addressService.GetDefaultAddress(s.testUserID)
	s.Nil(err, "获取默认地址不应该失败")
	s.NotNil(defaultAddress, "应该找到默认地址")
	s.Equal(1, defaultAddress.IsDefault, "返回的地址应该是默认地址")
}

// 测试获取默认地址 - 空用户ID
func (s *AddressServiceTestSuite) TestGetDefaultAddress_InvalidUserID() {
	defaultAddress, err := s.addressService.GetDefaultAddress("")
	s.Equal(ErrInvalidAddressData, err, "空用户ID应该返回无效地址数据错误")
	s.Nil(defaultAddress, "默认地址应该为nil")
}

// 测试获取默认地址 - 没有默认地址
func (s *AddressServiceTestSuite) TestGetDefaultAddress_NoDefaultAddress() {
	// 添加一个非默认地址
	req := s.createTestAddressRequest()
	req.IsDefault = false
	s.addressService.AddAddress(s.testUserID, req)

	// 获取默认地址
	defaultAddress, err := s.addressService.GetDefaultAddress(s.testUserID)
	s.Nil(err, "获取默认地址不应该失败")
	s.Nil(defaultAddress, "没有默认地址时应该返回nil")
}

// 测试获取默认地址 - 多个地址但没有默认地址
func (s *AddressServiceTestSuite) TestGetDefaultAddress_MultipleAddresses_NoDefault() {
	// 添加多个非默认地址
	req1 := s.createTestAddressRequest()
	req1.LinkMan = "张三"
	req1.IsDefault = false
	s.addressService.AddAddress(s.testUserID, req1)

	req2 := s.createTestAddressRequest()
	req2.LinkMan = "李四"
	req2.IsDefault = false
	s.addressService.AddAddress(s.testUserID, req2)

	// 获取默认地址
	defaultAddress, err := s.addressService.GetDefaultAddress(s.testUserID)
	s.Nil(err, "获取默认地址不应该失败")
	s.Nil(defaultAddress, "没有默认地址时应该返回nil")
}

// 测试获取默认地址 - 多个地址中有默认地址
func (s *AddressServiceTestSuite) TestGetDefaultAddress_MultipleAddresses_WithDefault() {
	// 添加多个地址，其中一个是默认地址
	req1 := s.createTestAddressRequest()
	req1.LinkMan = "张三"
	req1.IsDefault = false
	s.addressService.AddAddress(s.testUserID, req1)

	req2 := s.createTestAddressRequest()
	req2.LinkMan = "李四"
	req2.IsDefault = true
	s.addressService.AddAddress(s.testUserID, req2)

	req3 := s.createTestAddressRequest()
	req3.LinkMan = "王五"
	req3.IsDefault = false
	s.addressService.AddAddress(s.testUserID, req3)

	// 获取默认地址
	defaultAddress, err := s.addressService.GetDefaultAddress(s.testUserID)
	s.Nil(err, "获取默认地址不应该失败")
	s.NotNil(defaultAddress, "应该找到默认地址")
	s.Equal("李四", defaultAddress.LinkMan, "应该找到李四的默认地址")
	s.Equal(1, defaultAddress.IsDefault, "返回的地址应该是默认地址")
}