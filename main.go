package main

import (
	"bookstore/app/configs"
	"bookstore/app/routers"

	"bookstore/app/addresses"
	ad "bookstore/app/banner"
	"bookstore/app/goods"
	cart "bookstore/app/shoppingcart"
	"bookstore/app/user"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	configs.GetConfigInstance("config.yaml")
	db := configs.Cfg.DBConnection()
	skuRepo := goods.NewSkuRepoDB(db)
	cateRepo := &goods.CategoryRepo{}
	goodsService := goods.NewGoodsService(skuRepo, cateRepo)
	goodsHandler := goods.NewGoodsHandler(goodsService)

	// user
	userRepo := user.NewUserRepoDB(db)
	userService := user.NewUserServiceWithRepo(userRepo)
	userHandler := user.NewUserHandler(userService)

	// cart
	cartRepo := cart.NewCartRepoDB(db)
	cartService := cart.NewCartService(skuRepo, cartRepo)
	cartHandler := cart.NewCartHandler(cartService)

	// banner
	bannerRepo := ad.NewBannerRepoDB(db)
	bannerService := ad.NewBannerService(bannerRepo)
	bannerHandler := ad.NewBannerHandler(bannerService)

	// address
	addressRepo := addresses.NewAddressRepositoryDB(db)
	addressService := addresses.NewAddressService(addressRepo, db)
	addressHandler := addresses.NewAddressHandler(addressService)

	// 依赖注入
	r := gin.Default()
	routers.SetupRouter(r, bannerHandler, userHandler, cartHandler, addressHandler, goodsHandler)

	port := viper.Get("PORT")
	if port == nil {
		port = 9090
	}
	r.Run(":" + viper.GetString("PORT"))
}
