package goods

import (
	"bookstore/app/configs"
	"bookstore/app/testutils"
	"bookstore/app/utils"
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

func (r *SkuRepoDBTestSuite) Test_Delete() {

	exp := r.cd10()
	//use RandomId to distingish the test data
	id := exp.SkuId + utils.RandomImpl{}.GenStr()
	exp.SkuId = id
	picStr1 := "any.jpeg"
	pic1 := SkuCarouPicture{SkuId: id, PicStr: picStr1}
	picStr2 := "any.jpeg"
	pic2 := SkuCarouPicture{SkuId: id, PicStr: picStr2}
	pics := []SkuCarouPicture{pic1, pic2}
	exp.SkuCarouPictures = pics
	r.repo.Create(exp)
	saved := r.repo.FindWithCarouselPics(id)
	r.NotNil(saved)
	r.repo.Delete(exp)
	found := r.repo.FindWithCarouselPics(id)
	r.Nil(found)

}

func (r *SkuRepoDBTestSuite) Test_Create() {

	exp := r.cd10()
	//user RandomId to distingish the test data
	id := exp.SkuId + utils.RandomImpl{}.GenStr()
	exp.SkuId = id
	picStr1 := "-1.jpeg"
	pic1 := SkuCarouPicture{SkuId: id, PicStr: picStr1}
	picStr2 := "-2.jpeg"
	pic2 := SkuCarouPicture{SkuId: id, PicStr: picStr2}
	pics := []SkuCarouPicture{pic1, pic2}
	exp.SkuCarouPictures = pics

	r.Nil(r.repo.Create(exp))

	saved := r.repo.FindWithCarouselPics(id)
	r.Equal(id, saved.SkuId)
	r.Equal(exp.SkuCarouPictures[0].SkuId, pic1.SkuId)
	r.Equal(exp.SkuCarouPictures[0].PicStr, picStr1)
	r.Equal(exp.SkuCarouPictures[1].SkuId, pic1.SkuId)
	r.Equal(exp.SkuCarouPictures[1].PicStr, picStr2)
	r.repo.Delete(*saved)
	found := r.repo.FindWithCarouselPics(id)
	r.Nil(found)
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
