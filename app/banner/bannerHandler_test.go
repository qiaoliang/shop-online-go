package ad

import (
	"testing"

	"github.com/example/project/app/testutils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type BannerHandlerSuite struct {
	testutils.SupperSuite
	router *gin.Engine
}

func TestBannerHandlerSuite(t *testing.T) {
	suite.Run(t, new(BannerHandlerSuite))
}

func (st *BannerHandlerSuite) SetupSuite() {
	st.SupperSuite.SetupSuite()
	st.router = setupTestRouter()
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/v1")

	v1.GET("/banner/list", FetchBanners)
	return router
}
func (s *BannerHandlerSuite) Test_should_get_default_list_when_no_param_existed() {
	noParam := make(map[string]string, 0)
	noParam["type"] = ""
	noParam["token"] = ""
	body := testutils.HttpGet("/v1/banner/list", noParam, s.router)

	exp := `{"code":0,"data":[{"businessId":0,"dateAdd":"2022-05-05 11:26:09","id":222083,"linkUrl":"https://gitee.com/sagittatius","paixu":0,"picUrl":"http://localhost:9090/pic/banners/b0001.jpeg","remark":"跳转gitee sagittatius","status":0,"statusStr":"any","title":"any","type":"any","userId":1605},{"businessId":1,"dateAdd":"2022-05-05 11:26:09","id":222084,"linkUrl":"https://gitee.com/sagittatius","paixu":0,"picUrl":"http://localhost:9090/pic/banners/b0002.jpeg","remark":"跳转gitee sagittatius","status":0,"statusStr":"any","title":"any","type":"any","userId":1606}],"msg":"OK"}`
	s.Equal(exp, string(body), "should same.")
}
