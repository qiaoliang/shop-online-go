package utils

import (
	"bookstore/app/configs"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	configs.GetConfigInstance("../../config-test.yaml")
	os.Exit(m.Run())
}

func Test_generate_Aavatar(t *testing.T) {
	str := RandomMock{}.GenAavatarStr()
	reg, _ := regexp.Compile(`^[a-l]\.jpeg$`)
	assert.True(t, reg.MatchString(str), "Should be a jpeg file")
}
