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

func HttpGet(reqURL string, params map[string]string, r *gin.Engine) string {
	values := ""
	for key, val := range params {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp

	reqURL = reqURL + values
	req, _ := http.NewRequest("GET", reqURL, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		fmt.Printf("Http Request Error with reqMethod = %v, reqURL = %v, data = %v\n", "GET", reqURL, params)
	}

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func HttpPost(r *gin.Engine, data url.Values, reqURL string) string {
	req, _ := http.NewRequest("POST", reqURL, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()

	body := doIt(r, w, req, reqURL, data)
	return string(body)
}

func doIt(r *gin.Engine, w *httptest.ResponseRecorder, req *http.Request, reqURL string, data url.Values) []byte {
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		fmt.Printf("Http Request Error with reqMethod = POST, reqURL = %v, data = %v\n", reqURL, data)
	}

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
