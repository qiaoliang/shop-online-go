package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate_String(t *testing.T) {
	str := GenerateStr(4)
	assert.Equal(t, 4, len(str))
	str2 := GenerateStr(5)
	assert.Equal(t, 5, len(str2))
	result := make(map[string]string, 100)
	for i := 0; i < 100; i++ {
		str = GenerateStr(10)
		result[str] = str
	}
	assert.Equal(t, 100, len(result))
}
func Test_generate_Aavatar(t *testing.T) {
	str := GenerateAavatarStr()
	assert.Equal(t, 1, len(str))
	seed := "abcdefchigkl"
	assert.Containsf(t, seed, str, "%v should be in %v\n", str, seed)
}
