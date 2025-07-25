package addresses

import (
	"testing"
)

func TestAddressRepositoryMem_CRUD(t *testing.T) {
	repo := NewAddressRepositoryMem()
	addr := &Address{
		UserId:        "user1",
		LinkMan:       "张三",
		Mobile:        "13800000000",
		ProvinceStr:   "北京",
		CityStr:       "北京",
		AreaStr:       "朝阳区",
		DetailAddress: "望京SOHO",
		IsDefault:     1,
	}

	// Create
	err := repo.Create(addr)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if addr.Id == 0 {
		t.Fatalf("Id not set after Create")
	}

	// GetByID
	got, err := repo.GetByID(addr.Id)
	if err != nil {
		t.Fatalf("GetByID failed: %v", err)
	}
	if got.UserId != "user1" {
		t.Errorf("GetByID: want user1, got %s", got.UserId)
	}

	// Update
	addr.LinkMan = "李四"
	err = repo.Update(addr)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	got, _ = repo.GetByID(addr.Id)
	if got.LinkMan != "李四" {
		t.Errorf("Update: want 李四, got %s", got.LinkMan)
	}

	// ListByUserID
	other := &Address{UserId: "user2", LinkMan: "王五"}
	repo.Create(other)
	list, err := repo.ListByUserID("user1")
	if err != nil {
		t.Fatalf("ListByUserID failed: %v", err)
	}
	if len(list) != 1 || list[0].UserId != "user1" {
		t.Errorf("ListByUserID: want 1 user1, got %v", list)
	}

	// Delete
	err = repo.Delete(addr.Id)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	_, err = repo.GetByID(addr.Id)
	if err != ErrAddressNotFound {
		t.Errorf("GetByID after Delete: want ErrAddressNotFound, got %v", err)
	}
}