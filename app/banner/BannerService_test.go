package banner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// 自动迁移表结构
	err = db.AutoMigrate(&Banner{})
	assert.NoError(t, err)

	return db
}

func TestNewBannerService(t *testing.T) {
	db := setupTestDB(t)
	repo := NewBannerRepoDB(db)

	service := NewBannerService(repo)

	assert.NotNil(t, service)
	assert.NotNil(t, service.repo)
}

func TestFetchBanners_EmptyDatabase(t *testing.T) {
	db := setupTestDB(t)
	repo := NewBannerRepoDB(db)
	service := NewBannerService(repo)

	result := service.FetchBanners("home", "test-token")

	assert.Empty(t, result)
}

func TestFetchBanners_WithData(t *testing.T) {
	db := setupTestDB(t)
	repo := NewBannerRepoDB(db)
	service := NewBannerService(repo)

	// 插入测试数据
	testBanners := []Banner{
		{
			BusinessId: 1,
			DateAdd:    "2024-01-01",
			Id:         1,
			LinkUrl:    "https://example.com/1",
			Paixu:      1,
			PicUrl:     "https://example.com/image1.jpg",
			Remark:     "测试banner1",
			Status:     1,
			StatusStr:  "启用",
			Title:      "标题1",
			Type:       "home",
			UserId:     1,
		},
		{
			BusinessId: 2,
			DateAdd:    "2024-01-02",
			Id:         2,
			LinkUrl:    "https://example.com/2",
			Paixu:      2,
			PicUrl:     "https://example.com/image2.jpg",
			Remark:     "测试banner2",
			Status:     1,
			StatusStr:  "启用",
			Title:      "标题2",
			Type:       "category",
			UserId:     2,
		},
	}

	for _, banner := range testBanners {
		err := db.Create(&banner).Error
		assert.NoError(t, err)
	}

	result := service.FetchBanners("home", "test-token")

	// 验证返回了所有banner
	assert.Len(t, result, 2)

	// 验证数据映射正确
	for i, expected := range testBanners {
		assert.Equal(t, expected.BusinessId, result[i].BusinessId)
		assert.Equal(t, expected.DateAdd, result[i].DateAdd)
		assert.Equal(t, expected.Id, result[i].Id)
		assert.Equal(t, expected.LinkUrl, result[i].LinkUrl)
		assert.Equal(t, expected.Paixu, result[i].Paixu)
		assert.Equal(t, expected.PicUrl, result[i].PicUrl)
		assert.Equal(t, expected.Remark, result[i].Remark)
		assert.Equal(t, expected.Status, result[i].Status)
		assert.Equal(t, expected.StatusStr, result[i].StatusStr)
		assert.Equal(t, expected.Title, result[i].Title)
		assert.Equal(t, expected.Type, result[i].Catelogue)
		assert.Equal(t, expected.UserId, result[i].UserId)
	}
}

func TestFetchBanners_DataMapping(t *testing.T) {
	db := setupTestDB(t)
	repo := NewBannerRepoDB(db)
	service := NewBannerService(repo)

	// 插入一个测试banner
	testBanner := Banner{
		BusinessId: 123,
		DateAdd:    "2024-01-01 10:00:00",
		Id:         456,
		LinkUrl:    "https://test.com/link",
		Paixu:      789,
		PicUrl:     "https://test.com/image.jpg",
		Remark:     "测试备注",
		Status:     1,
		StatusStr:  "启用状态",
		Title:      "测试标题",
		Type:       "test-type",
		UserId:     999,
	}

	err := db.Create(&testBanner).Error
	assert.NoError(t, err)

	result := service.FetchBanners("test", "token")

	assert.Len(t, result, 1)

	// 验证数据映射
	bannerVM := result[0]
	assert.Equal(t, testBanner.BusinessId, bannerVM.BusinessId)
	assert.Equal(t, testBanner.DateAdd, bannerVM.DateAdd)
	assert.Equal(t, testBanner.Id, bannerVM.Id)
	assert.Equal(t, testBanner.LinkUrl, bannerVM.LinkUrl)
	assert.Equal(t, testBanner.Paixu, bannerVM.Paixu)
	assert.Equal(t, testBanner.PicUrl, bannerVM.PicUrl)
	assert.Equal(t, testBanner.Remark, bannerVM.Remark)
	assert.Equal(t, testBanner.Status, bannerVM.Status)
	assert.Equal(t, testBanner.StatusStr, bannerVM.StatusStr)
	assert.Equal(t, testBanner.Title, bannerVM.Title)
	assert.Equal(t, testBanner.Type, bannerVM.Catelogue)
	assert.Equal(t, testBanner.UserId, bannerVM.UserId)
}

func TestFetchBanners_WithDifferentTypes(t *testing.T) {
	db := setupTestDB(t)
	repo := NewBannerRepoDB(db)
	service := NewBannerService(repo)

	// 插入不同类型的banner
	testCases := []string{"home", "category", "promotion", ""}

	for i, btype := range testCases {
		testBanner := Banner{
			BusinessId: uint(i + 1),
			DateAdd:    "2024-01-01",
			Id:         uint(i + 1),
			LinkUrl:    "https://example.com",
			Paixu:      1,
			PicUrl:     "https://example.com/image.jpg",
			Remark:     "测试banner",
			Status:     1,
			StatusStr:  "启用",
			Title:      "测试标题",
			Type:       btype,
			UserId:     1,
		}

		err := db.Create(&testBanner).Error
		assert.NoError(t, err)
	}

	result := service.FetchBanners("any", "token")

	// 验证返回了所有banner
	assert.Len(t, result, len(testCases))

	// 验证每个banner的Type字段正确映射到Catelogue
	for i, expectedType := range testCases {
		assert.Equal(t, expectedType, result[i].Catelogue)
	}
}