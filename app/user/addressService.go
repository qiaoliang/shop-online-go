package user

import (
	"errors"

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
	// Basic validation in service layer
	if userID == "" || req.LinkMan == "" || req.Mobile == "" ||
		req.ProvinceStr == "" || req.CityStr == "" || req.AreaStr == "" ||
		req.DetailAddress == "" {
		return ErrInvalidAddressData
	}

	// 验证用户是否存在（userID 是手机号）
	user := s.userRepo.RetriveUserByMobile(userID)
	if user == nil {
		return ErrUserNotFound
	}

	newAddress := &Address{
		UserId:        user.Id, // 使用用户的真实ID
		LinkMan:       req.LinkMan,
		Mobile:        req.Mobile,
		ProvinceStr:   req.ProvinceStr,
		CityStr:       req.CityStr,
		AreaStr:       req.AreaStr,
		DetailAddress: req.DetailAddress,
		IsDefault:     req.IsDefault,
	}

	// 如果新地址为默认，则将该用户其他地址的 isDefault 设为 0
	if req.IsDefault == 1 {
		addresses, err := s.addressRepo.ListByUserID(user.Id)
		if err != nil {
			return err
		}
		for _, addr := range addresses {
			if addr.IsDefault == 1 {
				addr.IsDefault = 0
				err := s.addressRepo.Update(addr)
				if err != nil {
					return err
				}
			}
		}
	}

	// 创建新地址
	if err := s.addressRepo.Create(newAddress); err != nil {
		return err
	}

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
