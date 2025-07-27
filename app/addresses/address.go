package addresses

import (
	"errors"

	"gorm.io/gorm"
)

type Address struct {
	Id            int    `json:"id"`
	UserId        string `json:"userId"`
	LinkMan       string `json:"linkMan"`
	Mobile        string `json:"mobile"`
	ProvinceStr   string `json:"provinceStr"`
	CityStr       string `json:"cityStr"`
	AreaStr       string `json:"areaStr"`
	DetailAddress string `json:"detailAddress"`
	IsDefault     int    `json:"isDefault"`
}

type AddressRepository interface {
	Create(address *Address) error
	GetByID(id int) (*Address, error)
	Update(address *Address) error
	Delete(id int) error
	ListByUserID(userId string) ([]*Address, error)
}

// DB 版本实现

type AddressRepositoryDB struct {
	db *gorm.DB
}

func NewAddressRepositoryDB(db *gorm.DB) *AddressRepositoryDB {
	// 自动迁移表结构
	return &AddressRepositoryDB{db: db}
}

func (r *AddressRepositoryDB) Create(address *Address) error {
	return r.db.Create(address).Error
}

func (r *AddressRepositoryDB) GetByID(id int) (*Address, error) {
	var addr Address
	if err := r.db.First(&addr, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAddressNotFound
		}
		return nil, err
	}
	return &addr, nil
}

func (r *AddressRepositoryDB) Update(address *Address) error {
	return r.db.Save(address).Error
}

func (r *AddressRepositoryDB) Delete(id int) error {
	return r.db.Delete(&Address{}, id).Error
}

func (r *AddressRepositoryDB) ListByUserID(userId string) ([]*Address, error) {
	var result []*Address
	if err := r.db.Where("user_id = ?", userId).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

var ErrAddressNotFound = errors.New("address not found")
