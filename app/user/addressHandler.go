package user

import (
	"net/http"

	"bookstore/app/common/models"
	"bookstore/app/utils"

	"github.com/gin-gonic/gin"
)

// AddShippingAddressRequest defines the request body for adding a shipping address.
type AddShippingAddressRequest struct {
	Token         string `json:"token" binding:"required"`         // 用户认证token（必需）
	LinkMan       string `json:"linkMan" binding:"required"`
	Mobile        string `json:"mobile" binding:"required"`
	ProvinceStr   string `json:"provinceStr" binding:"required"`
	CityStr       string `json:"cityStr" binding:"required"`
	AreaStr       string `json:"areaStr" binding:"required"`
	DetailAddress string `json:"detailAddress" binding:"required"`
	IsDefault     int    `json:"isDefault"`
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
	var req AddShippingAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.JsonResult{
			Code: "400",
			Msg:  "Invalid request parameters",
		})
		return
	}

	// 使用请求体中的token作为用户ID
	userID := req.Token
	if userID == "" {
		c.JSON(http.StatusUnauthorized, models.JsonResult{
			Code: "401",
			Msg:  "User token is required",
		})
		return
	}

	err := h.addressService.AddAddress(userID, req)
	if err != nil {
		// TODO: Handle specific error types from service layer (Subtask 2.5)
		c.JSON(http.StatusInternalServerError, models.JsonResult{
			Code: "500",
			Msg:  "Failed to add address",
		})
		return
	}

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
