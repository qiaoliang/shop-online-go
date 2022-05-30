package utils

import (
	"math/rand"
)

type Util struct {
}

func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
func GenerateAavatarStr() string {
	str := "abcdefghijkl"
	bytes := []byte(str)
	var result []byte
	result = append(result, bytes[rand.Intn(len(bytes))])
	return string(result) + ".jpeg"
}
