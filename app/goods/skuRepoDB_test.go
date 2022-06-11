package goods

import (
	"bookstore/app/configs"
	"bookstore/app/testutils"
	"bookstore/app/utils"
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
	r.repo = getSkuRepoDB(configs.Cfg.DBConnection())
}
func (r *SkuRepoDBTestSuite) TeardownSuite() {
	r.SupperSuite.TeardownSuite()
	r.repo = nil
}

func (r *SkuRepoDBTestSuite) SetupTest() {}
func (r *SkuRepoDBTestSuite) Test_Find() {
	exp := r.cd10()
	result := r.repo.First(exp.SkuId)
	r.NotNil(result)
	r.EqualValues(&exp, result)
}
func (r *SkuRepoDBTestSuite) Test_Find_with_association() {
	exp := r.cd10WithPics()

	result := r.repo.FindWithCarouselPics(exp.SkuId)
	r.NotNil(result)
	r.Equal(&exp, result)
}

func (r *SkuRepoDBTestSuite) Test_Delete() {
	//arrange
	sku := r.prepare_a_sku()
	//act
	r.repo.Delete(sku)
	//assert
	found := r.repo.FindWithCarouselPics(sku.SkuId)
	r.Nil(found)

}
func (r *SkuRepoDBTestSuite) Test_Update() {

	exp := r.prepare_a_sku()
	exp.MinPrice = "0.0"
	exp.CategoryId = 888
	r.repo.Update(exp.SkuId, exp)

	found := r.repo.FindWithCarouselPics(exp.SkuId)
	r.Equal("0.0", found.MinPrice)
	r.Equal(uint(888), found.CategoryId)

	//clean up
	r.repo.Delete(*found)
}

func (r *SkuRepoDBTestSuite) Test_FindAll() {
	exp := r.repo.FindAll()
	r.NotNil(exp)
	r.Equal(8, len(exp))
	r.Equal(2, len(exp[0].SkuCarouPictures))
	r.Equal(2, len(exp[1].SkuCarouPictures))
	r.Equal(2, len(exp[4].SkuCarouPictures))
	r.Equal(2, len(exp[7].SkuCarouPictures))
	r.Equal("-01.jpeg", exp[7].SkuCarouPictures[0].PicStr)
	r.Equal("-02.jpeg", exp[7].SkuCarouPictures[1].PicStr)

}

func (r *SkuRepoDBTestSuite) Test_create_without_association() {
	//arange
	exp := r.aSkuWithoutPics()
	//act
	r.repo.Create(exp)
	//assert
	saved := r.repo.First(exp.SkuId)
	r.NotNil(saved)
	r.Equal(exp.SkuId, saved.SkuId)
	r.Equal(0, len(exp.SkuCarouPictures))
	//clean up
	r.repo.Delete(*saved)
}
func (r *SkuRepoDBTestSuite) Test_create_with_association() {
	//arange
	exp := r.anySkuWithPics()
	//act
	r.repo.Create(exp)
	//assert
	saved := r.repo.FindWithCarouselPics(exp.SkuId)
	r.NotNil(saved)
	r.Equal(exp.SkuId, saved.SkuId)
	r.Equal(exp.SkuCarouPictures[0].SkuId, saved.SkuCarouPictures[0].SkuId)
	r.Equal(exp.SkuCarouPictures[0].PicStr, saved.SkuCarouPictures[0].PicStr)
	r.Equal(exp.SkuCarouPictures[1].SkuId, saved.SkuCarouPictures[1].SkuId)
	r.Equal(exp.SkuCarouPictures[1].PicStr, saved.SkuCarouPictures[1].PicStr)
	//clean up
	r.repo.Delete(*saved)
}
func (r *SkuRepoDBTestSuite) prepare_a_sku() SKU {
	exp := r.anySkuWithPics()
	r.Nil(r.repo.Create(exp))
	return exp
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
		StatusStr:       SalingStatus(ONSAIL).String(),
		Aftersale:       AfterSaleType(BOTH),
	}
}
func (r *SkuRepoDBTestSuite) aSkuWithoutPics() SKU {
	ret := r.cd10()
	ret.SkuId = ret.SkuId + utils.RandomImpl{}.GenStr()
	return ret
}
func (r *SkuRepoDBTestSuite) anySkuWithPics() SKU {
	ret := r.cd10()
	ret.SkuId = ret.SkuId + utils.RandomImpl{}.GenStr()
	picStr1 := "-1.jpeg"
	pic1 := SkuCarouPicture{SkuId: ret.SkuId, PicStr: picStr1}
	picStr2 := "-2.jpeg"
	pic2 := SkuCarouPicture{SkuId: ret.SkuId, PicStr: picStr2}
	pics := []SkuCarouPicture{pic1, pic2}
	ret.SkuCarouPictures = pics
	return ret
}
func (r *SkuRepoDBTestSuite) cd10WithPics() SKU {
	return prepareSku_Cd10_With_Pics()
}

func prepareSku_Cd10_With_Pics() SKU {
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
		StatusStr:        SalingStatus(ONSAIL).String(),
		Aftersale:        AfterSaleType(BOTH),
		SkuCarouPictures: pics,
	}
}
