package routers

import (
	"net/http"
	"strconv"

	"bookstore/app/controllers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {
	port := viper.Get("PORT").(int)
	r := gin.Default()

	r.Use(Cors())

	r.StaticFS("/pic",http.Dir("./static"))

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v1")
	v1.GET("/banner/list",controllers.FetchBanners)
	v1.POST("/goods/list",controllers.FetchGoodsList)
	v1.POST("/user/m/login",controllers.UserLogin)


	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic("Cannot start service" + err.Error())
	}
}

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method
        origin := c.Request.Header.Get("Origin") //请求头部
        if origin != "" {
            //接收客户端发送的origin （重要！）
            c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
            //服务器支持的所有跨域请求的方法
            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
            //允许跨域设置可以返回其他子段，可以自定义字段
            c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
            // 允许浏览器（客户端）可以解析的头部 （重要）
            c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
            //设置缓存时间
            c.Header("Access-Control-Max-Age", "172800")
            //允许客户端传递校验信息比如 cookie (重要)
            c.Header("Access-Control-Allow-Credentials", "true")
        }

        //允许类型校验
        if method == "OPTIONS" {
            c.JSON(http.StatusOK, "ok!")
        }

        defer func() {
            if err := recover(); err != nil {

            }
        }()

        c.Next()
    }
 }
