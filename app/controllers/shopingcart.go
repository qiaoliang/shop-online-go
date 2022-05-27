package controllers

import (
	"bookstore/app/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PutIntoCart(c *gin.Context) {
	token := c.PostForm("token")

	goodsId := c.PostForm("goodsId")

	number := c.PostForm("number")
	id64, err1 := strconv.ParseUint(goodsId, 10, 32)

	vlm64, err2 := strconv.ParseUint(number, 10, 32)
	var result *repository.Cart
	if err1 != nil || err2 != nil {
		fmt.Println("error format of params for " + number)
		result = nil
	} else {
		result = repository.GetCartsInstance().AddOrderIntoCart(token, uint(id64), uint(vlm64))
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

type CartInfo struct {
	Token           string     `json:"token"`
	Infos           string     `json:"cartInfo"`
	NewItemQuantity uint       `json:"number"`
	Items           []CartItem `json:"items"`
}

func (ci *CartInfo) update(key string, quantity uint) {
	for i, _ := range ci.Items {
		item := &ci.Items[i]
		if item.Key == key {
			item.Quantity = quantity
			return
		}
	}
}

type CartItem struct {
	Key             string   `json:"key"`
	Pic             string   `json:"pic"`
	Status          uint     `json:"status"` // === 1 【失效】
	Name            string   `json:"name"`
	Sku             []string `json:"sku"`
	Price           uint     `json:"price"`
	Quantity        uint     `json:"number"`
	Selected        string   `json:"selected"`
	OptionValueName string   `json:"optionValueName"`
}

func UpdateShoppingCart(c *gin.Context) {
	token := c.Param("token")
	key := c.Param("Key")
	number, _ := strconv.Atoi(c.Param("number"))
	result := cartInfo(token)
	quantity := uint(number)
	result.update(key, quantity)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
func GetShopingCart(c *gin.Context) {

	token := c.Param("token")
	fmt.Printf("Token is :%v\n", token)
	if len(token) == 0 {
		token = "I_am_Stub_Token"
	}
	fmt.Printf("token is : %v\n", token)
	result := cartInfo(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func cartInfo(token string) CartInfo {
	sku := []string{"sku1", "sku3"}
	items := make([]CartItem, 0)
	item := CartItem{"key1", "http://localhost:9090/pic/goods/g7227946-01.jpeg", 0, "CD1.0", sku, 66.0, 3, "1", "valueName"}
	items = append(items, item)
	cartInfo := CartInfo{token, "iamInfos", 1, items}
	return cartInfo
}
