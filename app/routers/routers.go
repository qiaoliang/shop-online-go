package routers

import (
	"net/http"

	banner "bookstore/app/banner"
	cart "bookstore/app/shoppingcart"

	"bookstore/app/addresses"

	"bookstore/app/goods"
	"bookstore/app/order"
	"bookstore/app/security"
	"bookstore/app/user"

	"github.com/gin-gonic/gin"
)

// 依赖注入说明：所有 handler 需在 main.go 实例化后传入 SetupRouter
func SetupRouter(r *gin.Engine, bannerHandler *banner.BannerHandler, userHandler *user.UserHandler, cartHandler *cart.CartHandler, addressHandler *addresses.AddressHandler, goodsHandler *goods.GoodsHandler) {

	// 添加跨域访问中间件
	r.Use(allowCrossDomainAccess())

	// 添加根路径处理函数，访问 http://localhost:9090 时返回
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg": "服务已启动",
		})
	})

	// 添加ping路径处理函数
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	v1.GET("/verification/pic/get", security.GetCapChar)
	v1.GET("/verification/pic/check", security.VerifyCapChar)
	v1.GET("/verification/sms/get", security.GetSMSCode)
	// Goods Management
	v1.GET("/shop/goods/category/all", goodsHandler.FetchCatalogues)
	v1.GET("/shop/goods/detail", goodsHandler.GetGoodsDetail)
	v1.POST("/shop/goods/reputation", goodsHandler.FetchItemReputation)
	v1.POST("/goods/list", goodsHandler.FetchGoodsList)
	// Shopping Cart Management
	v1.GET("/shopping-cart/info", cartHandler.GetShopingCart)
	v1.POST("/shopping-cart/add", cartHandler.PutIntoCart)
	v1.POST("/shopping-cart/modifyNumber", cartHandler.ModifyNumberOfGoodsInCart)
	// User Management
	v1.POST("/user/m/register", userHandler.Register)
	v1.POST("/user/m/login", userHandler.Login)
	v1.GET("/user/detail", userHandler.GetUserDetail)
	v1.GET("/user/modify", userHandler.UpdateUserInfo)
	v1.GET("/user/amount", userHandler.GetUserAmount)
	v1.GET("/user/logout", userHandler.GetUserDetail)
	//User ShippingAddress Management
	v1.POST("/user/shipping-address/list", userHandler.GetDeliveryAddressList)
	v1.GET("/user/shipping-address/default", userHandler.GetDefaultDeliveryAddress)
	v1.POST("/user/shipping-address/add", addressHandler.AddAddress)
	// Advertise management
	v1.GET("/banner/list", bannerHandler.FetchBanners)
	// Order management
	v1.GET("/order/statistics", order.GetOrderStatistics)
	v1.GET("/discounts/statistics", order.DiscountStatistics)
	v1.GET("/discounts/coupons", order.Coupons)
	//books 相关路由保持不变
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
