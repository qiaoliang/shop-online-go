package main

import (
	"bookstore/app/goods"
	"testing"

	"github.com/stretchr/testify/suite"
)

type InitDataTest struct {
	suite.Suite
}

func TestInitDatTestTestSuite(t *testing.T) {
	suite.Run(t, new(InitDataTest))
}
func (suite *InitDataTest) BeforeTest(suiteName, testName string) {}
func (suite *InitDataTest) AfterTest(suiteName, testName string)  {}
func (suite *InitDataTest) SetupSuite()                           {}
func (suite *InitDataTest) TeardownSuite()                        {}
func (suite *InitDataTest) SetupTest()                            {}
func (suite *InitDataTest) Test_initCategory() {
	str := `[{"id":0,"name":"DevOps"},{"id":1,"name":"大数据"}]`
	exp := []goods.Category{{Id: uint(0), Name: "DevOps"}, {Id: uint(1), Name: "大数据"}}
	ret, err := initCategories([]byte(str))
	suite.True(err == nil)
	suite.Equal(2, len(ret))
	suite.EqualValues(exp, ret)
}
func (suite *InitDataTest) Test_initSKUs() {
	str := `[{"id":0,"name":"DevOps"},{"id":1,"name":"大数据"}]`
	exp := []goods.Category{{Id: uint(0), Name: "DevOps"}, {Id: uint(1), Name: "大数据"}}
	ret, err := initCategories([]byte(str))
	suite.True(err == nil)
	suite.Equal(2, len(ret))
	suite.EqualValues(exp, ret)
}
