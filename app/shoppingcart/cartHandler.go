package cart

import (
	"fmt"
	"net/http"
	"strconv"

	"bookstore/app/utils"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *CartService
}

func NewCartHandler(service *CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) PutIntoCart(c *gin.Context) {
	// 从认证中间件获取用户信息
	mobile := utils.GetMobileFromContext(c)
	if mobile == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "User not authenticated"})
		return
	}

	gid := c.PostForm("goodsId")
	number := c.PostForm("number")
	vlm64, err2 := strconv.ParseUint(number, 10, 32)
	if gid == "" || number == "" || err2 != nil {
		result := fmt.Sprintf("gid = %v, number = %v\n ", gid, number)
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": result, "msg": "OK"})
		return
	}
	result := h.service.PutItemsInCart(mobile, gid, uint(vlm64))
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func (h *CartHandler) ModifyNumberOfGoodsInCart(c *gin.Context) {
	// 从认证中间件获取用户信息
	mobile := utils.GetMobileFromContext(c)
	if mobile == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "User not authenticated"})
		return
	}

	gid := c.PostForm("key")
	numStr := c.PostForm("number")
	number, _ := strconv.Atoi(numStr)
	h.service.ModifyQuantityOfGoodsInCate(mobile, gid, uint(number))
	result := h.service.GetCartByToken(mobile)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func (h *CartHandler) GetShopingCart(c *gin.Context) {
	// 从认证中间件获取用户信息
	mobile := utils.GetMobileFromContext(c)
	if mobile == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "User not authenticated"})
		return
	}

	cart := h.service.GetCartByToken(mobile)
	var result interface{}
	if cart == nil {
		fmt.Println("没有找到 token 为 " + mobile + " 的购物车")
		result = ""
	} else {
		result = cart
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
