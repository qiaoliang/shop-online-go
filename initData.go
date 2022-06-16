package main

import (
	"bookstore/app/goods"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func initCategories(str []byte) ([]goods.Category, error) {
	var categories []goods.Category
	err := json.Unmarshal(str, &categories)
	return categories, err
}
func initSKUs(str []byte) ([]goods.SKU, error) {
	var ret []goods.SKU
	err := json.Unmarshal(str, &ret)
	return ret, err
}
func readFromFile(fn string) []byte {
	jsonFile, err := os.Open(fn)
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func fileList(dirPth string) (files []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if !fi.IsDir() {
			ok := strings.HasSuffix(fi.Name(), ".json")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}
	return files, nil
}
