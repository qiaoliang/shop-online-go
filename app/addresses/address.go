package addresses

import "errors"

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

type AddressRepositoryMem struct {
	data map[int]*Address
	lastID int
}

func NewAddressRepositoryMem() *AddressRepositoryMem {
	return &AddressRepositoryMem{
		data: make(map[int]*Address),
		lastID: 0,
	}
}

func (r *AddressRepositoryMem) Create(address *Address) error {
	r.lastID++
	address.Id = r.lastID
	r.data[address.Id] = address
	return nil
}

func (r *AddressRepositoryMem) GetByID(id int) (*Address, error) {
	addr, ok := r.data[id]
	if !ok {
		return nil, ErrAddressNotFound
	}
	return addr, nil
}

func (r *AddressRepositoryMem) Update(address *Address) error {
	if _, ok := r.data[address.Id]; !ok {
		return ErrAddressNotFound
	}
	r.data[address.Id] = address
	return nil
}

func (r *AddressRepositoryMem) Delete(id int) error {
	if _, ok := r.data[id]; !ok {
		return ErrAddressNotFound
	}
	delete(r.data, id)
	return nil
}

func (r *AddressRepositoryMem) ListByUserID(userId string) ([]*Address, error) {
	var result []*Address
	for _, addr := range r.data {
		if addr.UserId == userId {
			result = append(result, addr)
		}
	}
	return result, nil
}

var ErrAddressNotFound = errors.New("address not found")
