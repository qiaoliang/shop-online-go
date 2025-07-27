package user

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type Address struct {
	Id            int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserId        string `json:"userId" gorm:"column:user_id"`
	LinkMan       string `json:"linkMan" gorm:"column:link_man"`
	Mobile        string `json:"mobile" gorm:"column:mobile"`
	ProvinceStr   string `json:"provinceId" gorm:"column:province_str"`
	CityStr       string `json:"cityId" gorm:"column:city_str"`
	AreaStr       string `json:"districtId" gorm:"column:area_str"`
	DetailAddress string `json:"address" gorm:"column:detail_address"`
	IsDefault     int    `json:"isDefault" gorm:"column:is_default"`
}

// TableName 指定表名
func (Address) TableName() string {
	return "addresses"
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
	log.Printf("[DEBUG] AddressRepositoryDB.Create: 开始创建地址到数据库 - UserId: %s, LinkMan: %s, Mobile: %s",
		address.UserId, address.LinkMan, address.Mobile)

	err := r.db.Create(address).Error
	if err != nil {
		log.Printf("[DEBUG] AddressRepositoryDB.Create: 数据库创建失败 - %v", err)
		return err
	}

	log.Printf("[DEBUG] AddressRepositoryDB.Create: 地址创建成功，数据库ID: %d", address.Id)
	return nil
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
	log.Printf("[DEBUG] AddressRepositoryDB.Update: 开始更新地址 - ID: %d, UserId: %s", address.Id, address.UserId)

	err := r.db.Save(address).Error
	if err != nil {
		log.Printf("[DEBUG] AddressRepositoryDB.Update: 地址更新失败 - ID: %d, 错误: %v", address.Id, err)
		return err
	}

	log.Printf("[DEBUG] AddressRepositoryDB.Update: 地址更新成功 - ID: %d", address.Id)
	return nil
}

func (r *AddressRepositoryDB) Delete(id int) error {
	return r.db.Delete(&Address{}, id).Error
}

func (r *AddressRepositoryDB) ListByUserID(userId string) ([]*Address, error) {
	log.Printf("[DEBUG] AddressRepositoryDB.ListByUserID: 开始查询用户地址列表 - UserId: %s", userId)

	var result []*Address
	if err := r.db.Where("user_id = ?", userId).Find(&result).Error; err != nil {
		log.Printf("[DEBUG] AddressRepositoryDB.ListByUserID: 查询用户地址列表失败 - UserId: %s, 错误: %v", userId, err)
		return nil, err
	}

	log.Printf("[DEBUG] AddressRepositoryDB.ListByUserID: 查询用户地址列表成功 - UserId: %s, 地址数量: %d", userId, len(result))
	return result, nil
}

var ErrAddressNotFound = errors.New("address not found")
