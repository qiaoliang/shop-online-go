package utils

import (
	"math/rand"

	"github.com/rs/xid"
)

type ShopRandom interface {
	GenStr() string
	GenAavatarStr() string
}

type RandomMock struct{}

func (i RandomMock) GenStr() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < 10; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
func (i RandomMock) GenAavatarStr() string {
	return avatar()
}

type RandomImpl struct{}

func (i RandomImpl) GenStr() string {
	return xid.New().String()
}
func (i RandomImpl) GenAavatarStr() string {
	return avatar()
}

func avatar() string {
	str := "abcdefghijkl"
	bytes := []byte(str)
	var result []byte
	result = append(result, bytes[rand.Intn(len(bytes))])
	return string(result) + ".jpeg"
}
