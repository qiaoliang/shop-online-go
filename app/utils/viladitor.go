package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	fp, _ := filepath.Abs(path)
	fmt.Printf("config file path:%s\n", fp)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
