package goods

import (
	"bookstore/app/configs"
	"bookstore/app/utils"
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

func (st *BookHandlerSuite) setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.PATCH("/books/:id", UpdateBook)
	return router
}
