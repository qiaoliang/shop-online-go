package user

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

var ErrInvalidAddressData = errors.New("invalid address data")
var ErrUserNotFound = errors.New("user not found")

// AddressService defines the interface for address related business logic.
type AddressService interface {
	AddAddress(userID string, req AddShippingAddressRequest) error
	GetAddressList(userID string) ([]*Address, error)
	GetDefaultAddress(userID string) (*Address, error)
}

// addressService implements AddressService interface.
type addressService struct {
	addressRepo AddressRepository
	userRepo    UserRepo // 添加用户仓库依赖
	db          *gorm.DB // Add DB dependency for transactions
}

// NewAddressService creates a new AddressService.
func NewAddressService(repo AddressRepository, userRepo UserRepo, db *gorm.DB) AddressService {
	return &addressService{
		addressRepo: repo,
		userRepo:    userRepo,
		db:          db,
	}
}

// AddAddress adds a new shipping address for a user.
func (s *addressService) AddAddress(userID string, req AddShippingAddressRequest) error {
	log.Printf("[DEBUG] AddressService.AddAddress: 开始处理添加地址业务逻辑，用户ID: %s", userID)

	// Basic validation in service layer
	if userID == "" || req.LinkMan == "" || req.Mobile == "" ||
		req.Address == "" {
		log.Printf("[DEBUG] AddressService.AddAddress: 参数验证失败 - userID: %s, LinkMan: %s, Mobile: %s, Address: %s",
			userID, req.LinkMan, req.Mobile, req.Address)
		return ErrInvalidAddressData
	}

	log.Printf("[DEBUG] AddressService.AddAddress: 参数验证通过")

	// 验证用户是否存在（userID 是用户ID）
	user := s.userRepo.RetriveUserByID(userID)
	if user == nil {
		log.Printf("[DEBUG] AddressService.AddAddress: 用户不存在，用户ID: %s", userID)
		return ErrUserNotFound
	}

	log.Printf("[DEBUG] AddressService.AddAddress: 用户验证成功，用户ID: %s, 用户名: %s", userID, user.Nickname)

	// 直接使用字符串类型的ID
	provinceStr := req.ProvinceId
	cityStr := req.CityId
	areaStr := req.DistrictId

	// 将bool转换为int
	isDefault := 0
	if req.IsDefault {
		isDefault = 1
	}

	log.Printf("[DEBUG] AddressService.AddAddress: 准备创建地址对象 - ProvinceStr: %s, CityStr: %s, AreaStr: %s, IsDefault: %d",
		provinceStr, cityStr, areaStr, isDefault)

	newAddress := &Address{
		UserId:        user.Id, // 使用用户的真实ID
		LinkMan:       req.LinkMan,
		Mobile:        req.Mobile,
		ProvinceStr:   provinceStr,
		CityStr:       cityStr,
		AreaStr:       areaStr,
		DetailAddress: req.Address,
		IsDefault:     isDefault,
	}

	log.Printf("[DEBUG] AddressService.AddAddress: 地址对象创建完成，准备保存到数据库")

	// 如果新地址为默认，则将该用户其他地址的 isDefault 设为 0
	if isDefault == 1 {
		log.Printf("[DEBUG] AddressService.AddAddress: 新地址设置为默认地址，开始更新其他地址的默认状态")
		addresses, err := s.addressRepo.ListByUserID(user.Id)
		if err != nil {
			log.Printf("[DEBUG] AddressService.AddAddress: 获取用户现有地址列表失败 - %v", err)
			return err
		}

		log.Printf("[DEBUG] AddressService.AddAddress: 获取到用户现有地址数量: %d", len(addresses))

		for _, addr := range addresses {
			if addr.IsDefault == 1 {
				log.Printf("[DEBUG] AddressService.AddAddress: 更新地址ID %d 为非默认地址", addr.Id)
				addr.IsDefault = 0
				err := s.addressRepo.Update(addr)
				if err != nil {
					log.Printf("[DEBUG] AddressService.AddAddress: 更新地址默认状态失败，地址ID: %d, 错误: %v", addr.Id, err)
					return err
				}
			}
		}
		log.Printf("[DEBUG] AddressService.AddAddress: 其他地址默认状态更新完成")
	}

	// 创建新地址
	log.Printf("[DEBUG] AddressService.AddAddress: 开始创建新地址到数据库")
	if err := s.addressRepo.Create(newAddress); err != nil {
		log.Printf("[DEBUG] AddressService.AddAddress: 创建新地址失败 - %v", err)
		return err
	}

	log.Printf("[DEBUG] AddressService.AddAddress: 新地址创建成功，地址ID: %d", newAddress.Id)

	return nil
}

// GetAddressList retrieves all addresses for a user.
func (s *addressService) GetAddressList(userID string) ([]*Address, error) {
	if userID == "" {
		return nil, ErrInvalidAddressData
	}

	return s.addressRepo.ListByUserID(userID)
}

// GetDefaultAddress retrieves the default address for a user.
func (s *addressService) GetDefaultAddress(userID string) (*Address, error) {
	if userID == "" {
		return nil, ErrInvalidAddressData
	}

	addresses, err := s.addressRepo.ListByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 查找默认地址
	for _, addr := range addresses {
		if addr.IsDefault == 1 {
			return addr, nil
		}
	}

	// 没有找到默认地址，返回 nil
	return nil, nil
}
