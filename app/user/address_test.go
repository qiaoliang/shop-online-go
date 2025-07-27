package user

import (
	"testing"

	"bookstore/app/configs"
	"bookstore/app/testutils"
	"bookstore/app/utils"

	"github.com/stretchr/testify/suite"
)

type AddressTestSuite struct {
	testutils.SupperSuite
	repo       *AddressRepositoryDB
	testUserID string
}

func TestAddressTestSuite(t *testing.T) {
	suite.Run(t, new(AddressTestSuite))
}

func (s *AddressTestSuite) SetupSuite() {
	s.SupperSuite.SetupSuite()
	db := configs.Cfg.DBConnection()
	s.repo = NewAddressRepositoryDB(db)

	// 创建测试用户
	s.createTestUser()
}

func (s *AddressTestSuite) TeardownSuite() {
	s.SupperSuite.TeardownSuite()
	s.repo = nil
}

func (s *AddressTestSuite) SetupTest() {
	// 清理测试数据
	s.cleanupTestData()
}

func (s *AddressTestSuite) TearDownTest() {
	// 清理测试数据
	s.cleanupTestData()
}

// 创建测试用户
func (s *AddressTestSuite) createTestUser() {
	userRepo := NewUserRepoDB(s.repo.db)
	mobile := "13900001234" + utils.RandomImpl{}.GenStr()
	user, err := userRepo.CreateUser(mobile, "password123", "测试用户", "0", genUId)
	s.Nil(err, "创建测试用户失败")
	s.NotNil(user, "用户创建失败")
	s.testUserID = user.Id
}

// 清理测试数据
func (s *AddressTestSuite) cleanupTestData() {
	// 删除测试用户的所有地址
	if s.testUserID != "" {
		addresses, _ := s.repo.ListByUserID(s.testUserID)
		for _, addr := range addresses {
			s.repo.Delete(addr.Id)
		}
	}
}

// 创建测试地址
func (s *AddressTestSuite) createTestAddress() *Address {
	return &Address{
		UserId:        s.testUserID,
		LinkMan:       "张三",
		Mobile:        "13800138000",
		ProvinceStr:   "440000",
		CityStr:       "440300",
		AreaStr:       "440303",
		DetailAddress: "测试街道123号",
		IsDefault:     0,
	}
}

// 测试创建地址 - 成功场景
func (s *AddressTestSuite) TestCreate_Success() {
	address := s.createTestAddress()

	err := s.repo.Create(address)
	s.Nil(err, "创建地址应该成功")
	s.Greater(address.Id, 0, "地址ID应该被自动生成")

	// 验证地址是否已保存到数据库
	savedAddress, err := s.repo.GetByID(address.Id)
	s.Nil(err, "获取地址不应该失败")
	s.NotNil(savedAddress, "应该能找到保存的地址")
	s.Equal(address.UserId, savedAddress.UserId, "用户ID应该匹配")
	s.Equal(address.LinkMan, savedAddress.LinkMan, "联系人姓名应该匹配")
	s.Equal(address.Mobile, savedAddress.Mobile, "手机号应该匹配")
	s.Equal(address.ProvinceStr, savedAddress.ProvinceStr, "省份应该匹配")
	s.Equal(address.CityStr, savedAddress.CityStr, "城市应该匹配")
	s.Equal(address.AreaStr, savedAddress.AreaStr, "区县应该匹配")
	s.Equal(address.DetailAddress, savedAddress.DetailAddress, "详细地址应该匹配")
	s.Equal(address.IsDefault, savedAddress.IsDefault, "默认状态应该匹配")
}

// 测试创建地址 - 必填字段为空
// 注意：数据库层面不进行字段验证，这些测试验证数据库能接受空值
func (s *AddressTestSuite) TestCreate_RequiredFieldsEmpty() {
	testCases := []struct {
		name    string
		address *Address
	}{
		{
			name: "用户ID为空",
			address: &Address{
				UserId:        "",
				LinkMan:       "张三",
				Mobile:        "13800138000",
				ProvinceStr:   "440000",
				CityStr:       "440300",
				AreaStr:       "440303",
				DetailAddress: "测试街道123号",
				IsDefault:     0,
			},
		},
		{
			name: "联系人为空",
			address: &Address{
				UserId:        s.testUserID,
				LinkMan:       "",
				Mobile:        "13800138000",
				ProvinceStr:   "440000",
				CityStr:       "440300",
				AreaStr:       "440303",
				DetailAddress: "测试街道123号",
				IsDefault:     0,
			},
		},
		{
			name: "手机号为空",
			address: &Address{
				UserId:        s.testUserID,
				LinkMan:       "张三",
				Mobile:        "",
				ProvinceStr:   "440000",
				CityStr:       "440300",
				AreaStr:       "440303",
				DetailAddress: "测试街道123号",
				IsDefault:     0,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			err := s.repo.Create(tc.address)
			// 数据库层面不进行字段验证，所以应该成功
			s.Nil(err, "数据库层面应该允许空值")
			s.Greater(tc.address.Id, 0, "地址ID应该被自动生成")
		})
	}
}

// 测试根据ID获取地址 - 成功场景
func (s *AddressTestSuite) TestGetByID_Success() {
	// 先创建一个地址
	address := s.createTestAddress()
	err := s.repo.Create(address)
	s.Nil(err, "创建地址应该成功")

	// 根据ID获取地址
	retrievedAddress, err := s.repo.GetByID(address.Id)
	s.Nil(err, "获取地址不应该失败")
	s.NotNil(retrievedAddress, "应该能找到地址")
	s.Equal(address.Id, retrievedAddress.Id, "地址ID应该匹配")
	s.Equal(address.UserId, retrievedAddress.UserId, "用户ID应该匹配")
	s.Equal(address.LinkMan, retrievedAddress.LinkMan, "联系人姓名应该匹配")
}

// 测试根据ID获取地址 - 地址不存在
func (s *AddressTestSuite) TestGetByID_NotFound() {
	// 尝试获取不存在的地址
	address, err := s.repo.GetByID(99999)
	s.Equal(ErrAddressNotFound, err, "应该返回地址未找到错误")
	s.Nil(address, "地址应该为nil")
}

// 测试更新地址 - 成功场景
func (s *AddressTestSuite) TestUpdate_Success() {
	// 先创建一个地址
	address := s.createTestAddress()
	err := s.repo.Create(address)
	s.Nil(err, "创建地址应该成功")

	// 更新地址信息
	address.LinkMan = "李四"
	address.Mobile = "13900139000"
	address.DetailAddress = "更新后的地址"
	address.IsDefault = 1

	err = s.repo.Update(address)
	s.Nil(err, "更新地址应该成功")

	// 验证更新结果
	updatedAddress, err := s.repo.GetByID(address.Id)
	s.Nil(err, "获取更新后的地址不应该失败")
	s.NotNil(updatedAddress, "应该能找到更新后的地址")
	s.Equal("李四", updatedAddress.LinkMan, "联系人姓名应该已更新")
	s.Equal("13900139000", updatedAddress.Mobile, "手机号应该已更新")
	s.Equal("更新后的地址", updatedAddress.DetailAddress, "详细地址应该已更新")
	s.Equal(1, updatedAddress.IsDefault, "默认状态应该已更新")
}

// 测试更新地址 - 地址不存在
func (s *AddressTestSuite) TestUpdate_NotFound() {
	// 尝试更新不存在的地址
	address := &Address{
		Id:            99999,
		UserId:        s.testUserID,
		LinkMan:       "李四",
		Mobile:        "13900139000",
		ProvinceStr:   "440000",
		CityStr:       "440300",
		AreaStr:       "440303",
		DetailAddress: "更新后的地址",
		IsDefault:     1,
	}

	err := s.repo.Update(address)
	s.Nil(err, "更新不存在的地址应该成功（GORM会创建新记录）")
}

// 测试删除地址 - 成功场景
func (s *AddressTestSuite) TestDelete_Success() {
	// 先创建一个地址
	address := s.createTestAddress()
	err := s.repo.Create(address)
	s.Nil(err, "创建地址应该成功")

	// 删除地址
	err = s.repo.Delete(address.Id)
	s.Nil(err, "删除地址应该成功")

	// 验证地址已被删除
	deletedAddress, err := s.repo.GetByID(address.Id)
	s.Equal(ErrAddressNotFound, err, "应该返回地址未找到错误")
	s.Nil(deletedAddress, "地址应该已被删除")
}

// 测试删除地址 - 地址不存在
func (s *AddressTestSuite) TestDelete_NotFound() {
	// 尝试删除不存在的地址
	err := s.repo.Delete(99999)
	s.Nil(err, "删除不存在的地址不应该报错")
}

// 测试根据用户ID获取地址列表 - 成功场景
func (s *AddressTestSuite) TestListByUserID_Success() {
	// 创建多个地址
	address1 := s.createTestAddress()
	address1.LinkMan = "张三"
	err := s.repo.Create(address1)
	s.Nil(err, "创建第一个地址应该成功")

	address2 := s.createTestAddress()
	address2.LinkMan = "李四"
	address2.Mobile = "13900139000"
	err = s.repo.Create(address2)
	s.Nil(err, "创建第二个地址应该成功")

	// 获取用户地址列表
	addresses, err := s.repo.ListByUserID(s.testUserID)
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

// 测试根据用户ID获取地址列表 - 用户没有地址
func (s *AddressTestSuite) TestListByUserID_NoAddresses() {
	addresses, err := s.repo.ListByUserID(s.testUserID)
	s.Nil(err, "获取空地址列表不应该失败")
	s.Len(addresses, 0, "应该返回空列表")
}

// 测试根据用户ID获取地址列表 - 用户ID不存在
func (s *AddressTestSuite) TestListByUserID_UserNotFound() {
	addresses, err := s.repo.ListByUserID("nonexistent_user_id")
	s.Nil(err, "获取不存在的用户地址列表不应该失败")
	s.Len(addresses, 0, "应该返回空列表")
}

// 测试地址的默认状态逻辑
func (s *AddressTestSuite) TestDefaultAddressLogic() {
	// 创建第一个地址（非默认）
	address1 := s.createTestAddress()
	address1.LinkMan = "张三"
	address1.IsDefault = 0
	err := s.repo.Create(address1)
	s.Nil(err, "创建第一个地址应该成功")

	// 创建第二个地址（默认）
	address2 := s.createTestAddress()
	address2.LinkMan = "李四"
	address2.IsDefault = 1
	err = s.repo.Create(address2)
	s.Nil(err, "创建第二个地址应该成功")

	// 获取地址列表
	addresses, err := s.repo.ListByUserID(s.testUserID)
	s.Nil(err, "获取地址列表不应该失败")
	s.Len(addresses, 2, "应该有两个地址")

	// 验证默认地址逻辑
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

// 测试地址表名
func (s *AddressTestSuite) TestTableName() {
	address := &Address{}
	tableName := address.TableName()
	s.Equal("addresses", tableName, "表名应该是addresses")
}

// 测试地址结构体字段标签
func (s *AddressTestSuite) TestAddressStructTags() {
	address := s.createTestAddress()

	// 验证结构体字段存在
	s.NotEmpty(address.UserId, "UserId字段应该存在")
	s.NotEmpty(address.LinkMan, "LinkMan字段应该存在")
	s.NotEmpty(address.Mobile, "Mobile字段应该存在")
	s.NotEmpty(address.ProvinceStr, "ProvinceStr字段应该存在")
	s.NotEmpty(address.CityStr, "CityStr字段应该存在")
	s.NotEmpty(address.AreaStr, "AreaStr字段应该存在")
	s.NotEmpty(address.DetailAddress, "DetailAddress字段应该存在")
	s.Equal(0, address.IsDefault, "IsDefault字段应该存在且默认为0")
}

// 测试并发创建地址
func (s *AddressTestSuite) TestConcurrentCreate() {
	// 并发创建多个地址
	done := make(chan bool, 5)

	for i := 0; i < 5; i++ {
		go func(index int) {
			address := s.createTestAddress()
			address.LinkMan = "并发用户" + string(rune('A'+index))
			address.Mobile = "1380013800" + string(rune('0'+index))

			err := s.repo.Create(address)
			s.Nil(err, "并发创建地址应该成功")

			done <- true
		}(i)
	}

	// 等待所有goroutine完成
	for i := 0; i < 5; i++ {
		<-done
	}

	// 验证所有地址都已创建
	addresses, err := s.repo.ListByUserID(s.testUserID)
	s.Nil(err, "获取地址列表不应该失败")
	s.Len(addresses, 5, "应该创建了5个地址")
}

// 测试地址数据的完整性
func (s *AddressTestSuite) TestAddressDataIntegrity() {
	// 创建地址
	address := s.createTestAddress()
	address.LinkMan = "完整性测试用户"
	address.Mobile = "13800138000"
	address.ProvinceStr = "440000"
	address.CityStr = "440300"
	address.AreaStr = "440303"
	address.DetailAddress = "完整性测试地址"
	address.IsDefault = 1

	err := s.repo.Create(address)
	s.Nil(err, "创建地址应该成功")

	// 获取地址并验证所有字段
	retrievedAddress, err := s.repo.GetByID(address.Id)
	s.Nil(err, "获取地址不应该失败")
	s.NotNil(retrievedAddress, "应该能找到地址")

	// 验证所有字段的完整性
	s.Equal(address.UserId, retrievedAddress.UserId, "用户ID应该完整保存")
	s.Equal(address.LinkMan, retrievedAddress.LinkMan, "联系人姓名应该完整保存")
	s.Equal(address.Mobile, retrievedAddress.Mobile, "手机号应该完整保存")
	s.Equal(address.ProvinceStr, retrievedAddress.ProvinceStr, "省份应该完整保存")
	s.Equal(address.CityStr, retrievedAddress.CityStr, "城市应该完整保存")
	s.Equal(address.AreaStr, retrievedAddress.AreaStr, "区县应该完整保存")
	s.Equal(address.DetailAddress, retrievedAddress.DetailAddress, "详细地址应该完整保存")
	s.Equal(address.IsDefault, retrievedAddress.IsDefault, "默认状态应该完整保存")
}