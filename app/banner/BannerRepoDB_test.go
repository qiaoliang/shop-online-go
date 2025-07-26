package banner

import (
	"testing"

	"bookstore/app/configs"

	"github.com/stretchr/testify/assert"
)

func TestBannerRepoDB_FindAllBanners(t *testing.T) {
	// 初始化测试配置
	configs.GetConfigInstance("../../config-test.yaml")
	db := configs.Cfg.DBConnection()

	// 创建仓库实例
	repo := NewBannerRepoDB(db)

	// 执行测试
	t.Run("应该返回所有横幅数据", func(t *testing.T) {
		// 调用被测试方法
		banners := repo.FindAllBanners()

		// 验证结果
		assert.NotNil(t, banners, "返回的横幅列表不应为空")
		assert.Equal(t, 2, len(banners), "应该返回2个横幅")

		// 验证第一个横幅的数据
		assert.Equal(t, uint(222083), banners[0].Id)
		assert.Equal(t, "https://baidu.com", banners[0].LinkUrl)
		assert.Equal(t, "http://localhost:9090/pic/banners/b0001.jpeg", banners[0].PicUrl)
		assert.Equal(t, "跳转gitee sagittatius", banners[0].Remark)

		// 验证第二个横幅的数据
		assert.Equal(t, uint(222084), banners[1].Id)
		assert.Equal(t, "https://baidu.com", banners[1].LinkUrl)
		assert.Equal(t, "http://localhost:9090/pic/banners/b0002.jpeg", banners[1].PicUrl)
		assert.Equal(t, "跳转gitee sagittatius", banners[1].Remark)
	})
}