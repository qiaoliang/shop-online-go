package utils

import (
	"path/filepath"
	"runtime"
)

func GetConfigFileForTest() string {
	_, filename, _, _ := runtime.Caller(0)
	path, _ := filepath.Abs(filename)
	path = filepath.Dir(path) + "/../../config.yaml"
	return path
}
