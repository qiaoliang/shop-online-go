package user

import (
	"log"
	"net/http"

	"bookstore/app/common/models"
	"bookstore/app/utils"

	"github.com/gin-gonic/gin"
)

// AddShippingAddressRequest defines the request body for adding a shipping address.
type AddShippingAddressRequest struct {
	LinkMan     string `json:"linkMan" binding:"required"`     // 联系人姓名
	Mobile      string `json:"mobile" binding:"required"`      // 手机号
	Address     string `json:"address" binding:"required"`     // 详细地址
	IsDefault   bool   `json:"isDefault"`                      // 是否默认地址
	ProvinceId  string `json:"provinceId" binding:"required"`  // 省份ID
	CityId      string `json:"cityId" binding:"required"`      // 城市ID
	DistrictId  string `json:"districtId" binding:"required"`  // 区县ID
}

// AddressHandler handles address related HTTP requests.
type AddressHandler struct {
	addressService AddressService
}

// NewAddressHandler creates a new AddressHandler.
func NewAddressHandler(service AddressService) *AddressHandler {
	return &AddressHandler{
		addressService: service,
	}
}

// AddAddress handles the POST /v1/user/shipping-address/add API.
func (h *AddressHandler) AddAddress(c *gin.Context) {
	log.Printf("[DEBUG] AddAddress: 开始处理添加收货地址请求")

	var req AddShippingAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[DEBUG] AddAddress: 请求参数绑定失败 - %v", err)
		c.JSON(http.StatusBadRequest, models.JsonResult{
			Code: "400",
			Msg:  "Invalid request parameters",
		})
		return
	}

	log.Printf("[DEBUG] AddAddress: 请求参数绑定成功 - LinkMan: %s, Mobile: %s, ProvinceId: %s, CityId: %s, DistrictId: %s, Address: %s, IsDefault: %t",
		req.LinkMan, req.Mobile, req.ProvinceId, req.CityId, req.DistrictId, req.Address, req.IsDefault)

	// 从认证中间件获取用户ID
	userID := utils.GetUserIDFromContext(c)
	if userID == "" {
		log.Printf("[DEBUG] AddAddress: 用户未认证，无法获取用户ID")
		c.JSON(http.StatusUnauthorized, models.JsonResult{
			Code: "401",
			Msg:  "User not authenticated",
		})
		return
	}

	log.Printf("[DEBUG] AddAddress: 用户认证成功，用户ID: %s", userID)

	err := h.addressService.AddAddress(userID, req)
	if err != nil {
		log.Printf("[DEBUG] AddAddress: 服务层处理失败 - %v", err)
		// TODO: Handle specific error types from service layer (Subtask 2.5)
		c.JSON(http.StatusInternalServerError, models.JsonResult{
			Code: "500",
			Msg:  "Failed to add address",
		})
		return
	}

	log.Printf("[DEBUG] AddAddress: 地址添加成功，用户ID: %s", userID)

	c.JSON(http.StatusOK, models.JsonResult{
		Code: "200",
		Msg:  "Address added successfully",
	})
}

// GetAddressList handles the GET /v1/user/shipping-address/list API.
func (h *AddressHandler) GetAddressList(c *gin.Context) {
	// 从认证中间件获取用户ID
	userID := utils.GetUserIDFromContext(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, models.JsonResult{
			Code: "401",
			Msg:  "User not authenticated",
		})
		return
	}

	addresses, err := h.addressService.GetAddressList(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResult{
			Code: "500",
			Msg:  "Failed to get address list",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": addresses,
		"msg":  "Address list retrieved successfully",
	})
}

// GetDefaultAddress handles the GET /v1/user/shipping-address/default API.
func (h *AddressHandler) GetDefaultAddress(c *gin.Context) {
	// 从认证中间件获取用户ID
	userID := utils.GetUserIDFromContext(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, models.JsonResult{
			Code: "401",
			Msg:  "User not authenticated",
		})
		return
	}

	address, err := h.addressService.GetDefaultAddress(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResult{
			Code: "500",
			Msg:  "Failed to get default address",
		})
		return
	}

	// 如果没有默认地址，返回空数据
	if address == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"data": nil,
			"msg":  "No default address found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": address,
		"msg":  "Default address retrieved successfully",
	})
}
