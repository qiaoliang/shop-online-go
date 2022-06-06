package goods

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type BookHandlerSuite struct {
	suite.Suite
	router *gin.Engine
}

func TestBookHandlerSuiteSuite(t *testing.T) {
	suite.Run(t, new(BookHandlerSuite))
}

func (st *BookHandlerSuite) SetupSuite() {
	st.router = st.setupTestRouter()
	configs.GetConfigInstance(utils.GetConfigFileForTest())
	configs.Cfg.Upgrade()
}

func (st *BookHandlerSuite) Test_Get_error_when_Update_unexisted_Book() {
	//构建参数
	var jsonStr = []byte(`[{"Title": "newTitle", "Author": "NewAuthor"}]`)
	url := configs.Cfg.Host + ":" + configs.Cfg.Port + "/books/0"
	body := utils.HttpPatch(url, jsonStr, st.router)

	exp := `{"error":"Record Not Found"}`
	st.Equal(exp, string(body), "should same.")
}

func (st *BookHandlerSuite) Test_Get_book_when_Book_existed() {
	//构建参数

	url := configs.Cfg.Host + ":" + configs.Cfg.Port + "/books/1"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	exp := `{"data":{"id":1,"title":"little prince","author":"Antoine"}}`
	st.Equal(exp, string(body), "should same.")
}

func (st *BookHandlerSuite) Test_Get_books_when_Books_existed() {
	//构建参数

	url := configs.Cfg.Host + ":" + configs.Cfg.Port + "/books"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	exp := `{"data":[{"id":1,"title":"little prince","author":"Antoine"},{"id":2,"title":"Les Trois Mousquetaires","author":"Alexandre Dumas fils"},{"id":3,"title":"Continuous Delivery","author":"Jez"}]}`
	st.Equal(exp, string(body), "should same.")
}

func (st *BookHandlerSuite) Should_Get_updated_when_the_Book_existed() {
	//构建参数
	url := configs.Cfg.Host + ":" + configs.Cfg.Port + "/books/1"
	input := UpdateBookInput{"newTitle", "NewAuthor"}
	inputStr, _ := json.Marshal(&input)
	body := utils.HttpPatch(url, inputStr, st.router)

	exp := `{"data":{"id":1,"title":"newTitle","author":"NewAuthor"}}`

	st.Equal(exp, string(body), "should same.")
}

func (st *BookHandlerSuite) setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.GET("/books", FindBooks)
	v1.POST("/books", CreateBook)
	v1.GET("/books/:id", FindBook)
	v1.PATCH("/books/:id", UpdateBook)
	v1.DELETE("/books/:id", DeleteBook)
	return router
}
