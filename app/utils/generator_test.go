package utils

import (
	"bookstore/app/testutils"
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GeneratorTestSuite struct {
	testutils.SupperSuite
}

func TestGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(GeneratorTestSuite))
}



func (suite *GeneratorTestSuite) SetupSuite() {
	suite.SupperSuite.SetupSuite()
}

func (suite *GeneratorTestSuite) TestGenerateAvatar() {
	str := RandomMock{}.GenAavatarStr()
	reg, _ := regexp.Compile(`^[a-l]\.jpeg$`)
	suite.True(reg.MatchString(str), "Should be a jpeg file")
}
