package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

type JsonResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func JsonToStruct(jsonStr string, data interface{}) {
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(errors.New("Unmarshal error!!! the jsonStr is :" + jsonStr))
	}

}
func GetConfigFileForTest() string {
	_, filename, _, _ := runtime.Caller(0)
	path, _ := filepath.Abs(filename)
	path = filepath.Dir(path) + "/../../config.yaml"
	return path
}

func HttpRequest(r *gin.Engine, data url.Values, reqMethod string, reqURL string) string {
	req, _ := http.NewRequest(reqMethod, reqURL, bytes.NewBufferString(data.Encode()))
	if reqMethod == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		fmt.Printf("Http Request Error with reqMethod = %v, reqURL = %v, data = %v\n", reqMethod, reqURL, data)
	}

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
