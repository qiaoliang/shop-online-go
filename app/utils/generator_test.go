package utils

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate_Aavatar(t *testing.T) {
	str := NewRandom().GenAavatarStr()
	reg, _ := regexp.Compile(`^[a-l]\.jpeg$`)
	assert.True(t, reg.MatchString(str), "Should be a jpeg file")
}
