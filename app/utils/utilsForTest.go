package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

func GetConfigFileForTest() string {
	_, filename, _, _ := runtime.Caller(0)
	path, _ := filepath.Abs(filename)
	path = filepath.Dir(path) + "/../../config.yaml"
	return path
}

func HttpRequest(r *gin.Engine, data url.Values, reqMethod string, reqURL string) []byte {
	req, _ := http.NewRequest(reqMethod, reqURL, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		fmt.Printf("Http Request Error with reqMethod = %v, reqURL = %v, data = %v\n", reqMethod, reqURL, data)
	}

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
