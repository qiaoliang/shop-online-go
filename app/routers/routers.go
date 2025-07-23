package routers

import (
	"net/http"
	"strconv"

	banner "github.com/example/project/app/banner"
	cart "github.com/example/project/app/shoppingcart"

	"github.com/example/project/app/goods"
	"github.com/example/project/app/order"
	"github.com/example/project/app/security"
	"github.com/example/project/app/user"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {
	port := viper.Get("PORT").(int)
	r := gin.Default()
	// This is Demo, and not good on Prod.
	r.Use(allowCrossDomainAccess())

	r.StaticFS("/pic", http.Dir("./static"))

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg": "商城后台已经正常启动。",
		})
	})

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	SetupRouter(r)

	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic("Cannot start service" + err.Error())
	}
}

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/v1")

	v1.GET("/verification/pic/get", security.GetCapChar)
	v1.GET("/verification/pic/check", security.VerifyCapChar)
	v1.GET("/verification/sms/get", security.GetSMSCode)

	// Advertise management
	v1.GET("/banner/list", banner.FetchBanners)

	// User Management
	v1.POST("/user/m/register", user.Register)
	v1.POST("/user/m/login", user.Login)
	v1.GET("/user/detail", user.GetUserDetail)
	v1.GET("/user/modify", user.UpdateUserInfo)
	v1.GET("/user/amount", user.GetUserAmount)
	v1.GET("/user/logout", user.GetUserDetail)

	//User ShippingAddress Management
	v1.POST("/user/shipping-address/list", user.GetDeliveryAddressList)
	v1.GET("/user/shipping-address/default", user.GetDefaultDeliveryAddress)
	v1.POST("/user/shipping-address/add", user.AddDeliveryAddress)

	// Order management
	v1.GET("/order/statistics", order.GetOrderStatistics)
	v1.GET("/discounts/statistics", order.DiscountStatistics)
	v1.GET("/discounts/coupons", order.Coupons)

	//Goods Management
	v1.GET("/shop/goods/category/all", goods.FetchCatalogues)
	v1.GET("/shop/goods/detail", goods.GetGoodsDetail)
	v1.POST("/shop/goods/reputation", goods.FetchItemReputation)
	v1.POST("/goods/list", goods.FetchGoodsList)

	// Shopping Cart Management
	v1.GET("/shopping-cart/info", cart.GetShopingCart)
	v1.POST("/shopping-cart/add", cart.PutIntoCart)
	v1.POST("/shopping-cart/modifyNumber", cart.ModifyNumberOfGoodsInCart)

	r.GET("/books", goods.FindBooks)
	r.POST("/books", goods.CreateBook)
	r.GET("/books/:id", goods.FindBook)
	r.PATCH("/books/:id", goods.UpdateBook)
	r.DELETE("/books/:id", goods.DeleteBook)
}

func allowCrossDomainAccess() gin.HandlerFunc {
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
				panic(err)
			}
		}()

		c.Next()
	}
}
