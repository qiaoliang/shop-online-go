package addresses

import (
	"errors"

	"gorm.io/gorm"
)

var ErrInvalidAddressData = errors.New("invalid address data")

// AddressService defines the interface for address related business logic.
type AddressService interface {
	AddAddress(userID string, req AddShippingAddressRequest) error
}

// addressService implements AddressService interface.
type addressService struct {
	addressRepo AddressRepository
	db          *gorm.DB // Add DB dependency for transactions
}

// NewAddressService creates a new AddressService.
func NewAddressService(repo AddressRepository, db *gorm.DB) AddressService {
	return &addressService{
		addressRepo: repo,
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

	newAddress := &Address{
		UserId:        userID,
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
		addresses, err := s.addressRepo.ListByUserID(userID)
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
