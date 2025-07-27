package addresses

import (
	"net/http"

	"bookstore/app/common/models"
	"bookstore/app/utils"

	"github.com/gin-gonic/gin"
)

// AddShippingAddressRequest defines the request body for adding a shipping address.
type AddShippingAddressRequest struct {
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

	// 从认证中间件获取用户ID
	userID := utils.GetUserIDFromContext(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, models.JsonResult{
			Code: "401",
			Msg:  "User not authenticated",
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
