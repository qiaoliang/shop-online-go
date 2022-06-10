package goods

import (
	"bookstore/app/configs"
	"bookstore/app/testutils"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SkuRepoDBTestSuite struct {
	testutils.SupperSuite
	repo SkuRepoIf
}

func TestSkuRepoDBTestSuite(t *testing.T) {
	suite.Run(t, new(SkuRepoDBTestSuite))
}

func (r *SkuRepoDBTestSuite) BeforeTest(suiteName, testName string) {}

func (r *SkuRepoDBTestSuite) AfterTest(suiteName, testName string) {}

func (r *SkuRepoDBTestSuite) SetupSuite() {
	r.SupperSuite.SetupSuite()
	skuRepo = nil
	r.repo = getSkuRepoDB(configs.Cfg.GormDB())
}
func (r *SkuRepoDBTestSuite) TeardownSuite() {
	r.SupperSuite.TeardownSuite()
	r.repo = nil
}

func (r *SkuRepoDBTestSuite) SetupTest() {}
func (r *SkuRepoDBTestSuite) Test_Find() {
	exp := r.cd10()
	result := r.repo.Find(exp.SkuId)

	r.NotNil(result)
	r.EqualValues(&exp, result)
}
func (r *SkuRepoDBTestSuite) Test_Find_with_association() {
	exp := r.cd10WithPics()
	result := r.repo.Find(exp.SkuId)
	log.Printf("%v\n", result)
	result.StatusStr = SalingStatus(result.Status).String()

	result = r.repo.FindWithCarouselPics(exp.SkuId)
	r.NotNil(result)
	r.Equal(&exp, result)
}

func (r *SkuRepoDBTestSuite) cd10() SKU {
	return SKU{
		SkuId:           "g7225946",
		Name:            "持续交付1.0",
		CategoryId:      0,
		RecommendStatus: "1",
		PicStr:          "g7225946.jpeg",
		Unit:            "册",
		Stock:           110,
		MinPrice:        "66.0",
		OriginalPrice:   "99.0",
		Logistics:       "1",
		Content:         "DevOps 的第一本书",
		Status:          SalingStatus(ONSAIL),
		Aftersale:       AfterSaleType(BOTH),
	}
}

func (r *SkuRepoDBTestSuite) cd10WithPics() SKU {
	Pic1 := SkuCarouPicture{
		Id:     0,
		SkuId:  "g7225946",
		PicStr: "-01.jpeg",
	}
	Pic2 := SkuCarouPicture{
		Id:     0,
		SkuId:  "g7225946",
		PicStr: "-02.jpeg",
	}
	pics := []SkuCarouPicture{Pic1, Pic2}
	return SKU{
		SkuId:            "g7225946",
		Name:             "持续交付1.0",
		CategoryId:       0,
		RecommendStatus:  "1",
		PicStr:           "g7225946.jpeg",
		Unit:             "册",
		Stock:            110,
		MinPrice:         "66.0",
		OriginalPrice:    "99.0",
		Logistics:        "1",
		Content:          "DevOps 的第一本书",
		Status:           SalingStatus(ONSAIL),
		Aftersale:        AfterSaleType(BOTH),
		SkuCarouPictures: pics,
	}
}
