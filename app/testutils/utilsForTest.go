package testutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

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
	return "../../config-test.yaml"
}

func HttpPatch1(reqURL string, data map[string]interface{}, r *gin.Engine) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	req, err := http.NewRequest("PATCH", reqURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err.Error()
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	body := doIt(r, w, req, reqURL, nil)
	return string(body)
}

func HttpGet(reqURL string, params map[string]string, r *gin.Engine) string {
	values := ""
	for key, val := range params {
		values += "&" + key + "=" + val
	}
	if len(values) != 0 {
		temp := values[1:]
		values = "?" + temp
	}
	reqURL = reqURL + values
	httpMethod := "GET"
	req, _ := http.NewRequest(httpMethod, reqURL, nil)
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
	HTTPMethod := "POST"
	return HttpMethod(HTTPMethod, reqURL, data, r)
}
func HttpDelete(reqURL string, data url.Values, r *gin.Engine) string {
	HTTPMethod := "DELETE"
	return HttpMethod(HTTPMethod, reqURL, data, r)
}

func HttpMethod(HTTPMethod string, reqURL string, data url.Values, r *gin.Engine) string {
	req, _ := http.NewRequest(HTTPMethod, reqURL, bytes.NewBufferString(data.Encode()))
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
