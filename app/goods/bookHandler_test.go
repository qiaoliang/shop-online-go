package goods

import (
	"bookstore/app/configs"
	"bookstore/app/testutils"
	"net/url"
	"strconv"
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
	configs.GetConfigInstance(testutils.GetConfigFileForTest())
	configs.Cfg.Upgrade()
}
func (st *BookHandlerSuite) TeardownSuite() {
	configs.Cfg.Downgrade()
}

func (st *BookHandlerSuite) Test_Get_error_when_Update_unexisted_Book() {
	//构建参数
	body := testutils.HttpGet("/books/8888", nil, st.router)

	exp := `{"error":"Record not found"}`
	st.Equal(exp, string(body), "should same.")
}

func (st *BookHandlerSuite) Test_Get_book_when_Book_existed() {

	body := testutils.HttpGet("/books/1", nil, st.router)

	exp := `{"data":{"id":1,"title":"little prince","author":"Antoine"}}`
	st.Equal(exp, string(body), "should same.")
}

func (st *BookHandlerSuite) Test_Delete_book_when_Book_existed() {
	book := Book{7777, "willBeDeleted", "willBeDeleted"}
	configs.Cfg.GormDB().Create(&book)

	body := testutils.HttpDelete("/books/"+strconv.Itoa(book.ID), nil, st.router)
	exp := `{"data":true}`
	st.Equal(exp, string(body), "should same.")
}

func (st *BookHandlerSuite) Test_Get_books_when_Books_existed() {

	body := testutils.HttpGet("/books", nil, st.router)

	exp := `{"data":[{"id":1,"title":"little prince","author":"Antoine"},{"id":2,"title":"Les Trois Mousquetaires","author":"Alexandre Dumas fils"},{"id":3,"title":"Continuous Delivery","author":"Jez"}]}`
	st.Equal(exp, string(body), "should same.")

}

func (st *BookHandlerSuite) Test_updated_when_the_Book_existed() {
	book := Book{7777, "willBeDeleted", "willBeDeleted"}
	configs.Cfg.GormDB().Create(&book)
	data := map[string]interface{}{
		"id":     7777,
		"title":  "new Title",
		"author": "new Author",
	}
	url := "/books/7777"

	body := testutils.HttpPatch1(url, data, st.router)

	exp := `{"data":{"id":7777`
	st.Contains(string(body), exp)
	st.Contains(string(body), "new Title")
	st.Contains(string(body), "new Author")

	configs.Cfg.GormDB().Delete(&book)

}

func (st *BookHandlerSuite) TEST_CREATE_A_BOOK() {
	data := url.Values{}
	data.Set("title", "haha")
	data.Add("author", "wowowo")
	body := testutils.HttpPost(st.router, data, "/books")

	exp := `{"data":{"id":4,"title":"haha","author":"wowowo"}}`
	st.Equal(exp, string(body), "should same.")

}

func (st *BookHandlerSuite) setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.GET("/books", FindBooks)
	router.POST("/books", CreateBook)
	router.GET("/books/:id", FindBook)
	router.PATCH("/books/:id", UpdateBook)
	router.DELETE("/books/:id", DeleteBook)
	return router
}
