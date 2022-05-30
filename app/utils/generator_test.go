package utils

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate_String(t *testing.T) {
	str := RandomStr(4)
	assert.Equal(t, 4, len(str))
	str2 := RandomStr(5)
	assert.Equal(t, 5, len(str2))
	result := make(map[string]string, 100)
	for i := 0; i < 100; i++ {
		str = RandomStr(10)
		result[str] = str
	}
	assert.Equal(t, 100, len(result))
}
func Test_generate_Aavatar(t *testing.T) {
	str := GenerateAavatarStr()
	reg, _ := regexp.Compile(`^[a-l]\.jpeg$`)
	assert.True(t, reg.MatchString(str), "Should be a jpeg file")
}
