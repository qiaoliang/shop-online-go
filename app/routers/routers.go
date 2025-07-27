package routers

import (
	"log"
	"net/http"

	banner "bookstore/app/banner"
	cart "bookstore/app/shoppingcart"

	"bookstore/app/goods"
	"bookstore/app/order"
	"bookstore/app/security"
	"bookstore/app/user"

	"github.com/gin-gonic/gin"
)

// 依赖注入说明：所有 handler 需在 main.go 实例化后传入 SetupRouter
func SetupRouter(r *gin.Engine, bannerHandler *banner.BannerHandler, userHandler *user.UserHandler, cartHandler *cart.CartHandler, addressHandler *user.AddressHandler, goodsHandler *goods.GoodsHandler, authMiddleware *security.AuthMiddleware) {

	// 添加跨域访问中间件
	r.Use(allowCrossDomainAccess())

	// 配置静态文件路由
	r.StaticFS("/pic", http.Dir("./static"))

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

	// 公开接口 - 不需要认证
	v1.GET("/verification/pic/get", security.GetCapChar)
	v1.GET("/verification/pic/check", security.VerifyCapChar)
	v1.GET("/verification/sms/get", security.GetSMSCode)

	// 用户认证相关接口
	v1.POST("/user/m/register", userHandler.Register)
	v1.POST("/user/m/login", userHandler.Login)

	// 商品相关接口 - 公开
	v1.GET("/shop/goods/category/all", goodsHandler.FetchCatalogues)
	v1.GET("/shop/goods/detail", goodsHandler.GetGoodsDetail)
	v1.POST("/shop/goods/reputation", goodsHandler.FetchItemReputation)
	v1.POST("/goods/list", goodsHandler.FetchGoodsList)

	// 广告相关接口 - 公开
	v1.GET("/banner/list", bannerHandler.FetchBanners)

	// 需要认证的接口组
	authenticated := v1.Group("")
	authenticated.Use(authMiddleware.Authenticate()) // 使用可选认证中间件

	// 用户相关接口 - 需要认证
	authenticated.GET("/user/detail", userHandler.GetUserDetail)
	authenticated.GET("/user/modify", userHandler.UpdateUserInfo)
	authenticated.GET("/user/amount", userHandler.GetUserAmount)
	authenticated.GET("/user/logout", userHandler.GetUserDetail)

	// 购物车相关接口 - 需要认证
	authenticated.GET("/shopping-cart/info", cartHandler.GetShopingCart)
	authenticated.POST("/shopping-cart/add", cartHandler.PutIntoCart)
	authenticated.POST("/shopping-cart/modifyNumber", cartHandler.ModifyNumberOfGoodsInCart)

	// 收货地址相关接口 - 需要认证
	authenticated.GET("/user/shipping-address/list", func(c *gin.Context) {
		log.Printf("[DEBUG] Router: 路由到获取地址列表 - 路径: %s, 方法: %s, 用户ID: %s",
			c.Request.URL.Path, c.Request.Method, c.GetString("userID"))
		addressHandler.GetAddressList(c)
	})
	authenticated.GET("/user/shipping-address/default", func(c *gin.Context) {
		log.Printf("[DEBUG] Router: 路由到获取默认地址 - 路径: %s, 方法: %s, 用户ID: %s",
			c.Request.URL.Path, c.Request.Method, c.GetString("userID"))
		addressHandler.GetDefaultAddress(c)
	})
	authenticated.POST("/user/shipping-address/add", func(c *gin.Context) {
		log.Printf("[DEBUG] Router: 路由到添加地址 - 路径: %s, 方法: %s, 用户ID: %s",
			c.Request.URL.Path, c.Request.Method, c.GetString("userID"))
		addressHandler.AddAddress(c)
	})

	// 订单相关接口 - 需要认证
	authenticated.GET("/order/statistics", order.GetOrderStatistics)
	authenticated.GET("/discounts/statistics", order.DiscountStatistics)
	authenticated.GET("/discounts/coupons", order.Coupons)

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
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token, session, Content-Type")
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
			c.Abort() // 阻止继续执行后续中间件
			return
		}

		defer func() {
			if err := recover(); err != nil {
				panic(err)
			}
		}()

		c.Next()
	}
}
