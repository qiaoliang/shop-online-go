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
		origin := c.Request.Header.Get("Origin") //????????????
		if origin != "" {
			//????????????????????????origin ???????????????
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//?????????????????????????????????????????????
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//??????????????????????????????????????????????????????????????????
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// ??????????????????????????????????????????????????? ????????????
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//??????????????????
			c.Header("Access-Control-Max-Age", "172800")
			//??????????????????????????????????????? cookie (??????)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//??????????????????
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
