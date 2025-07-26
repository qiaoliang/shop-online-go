package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 新增类型：UserHandler，持有 UserService

type UserHandler struct {
	us *UserService
}

func NewUserHandler(us *UserService) *UserHandler {
	return &UserHandler{us: us}
}

func (h *UserHandler) Login(c *gin.Context) {
	deviceId := c.PostForm("deviceId")
	deviceName := c.PostForm("deviceName")
	mobile := c.PostForm("mobile")
	pwd := c.PostForm("pwd")
	user, err := h.us.login(deviceId, deviceName, mobile, pwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": user, "msg": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": "ok"})
}
func Logout(c *gin.Context) {
	token, _ := c.GetQuery("token")
	GetUserService().logout(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(nil), "msg": "OK"})
}
func (h *UserHandler) Register(c *gin.Context) {
	autoLogin := c.PostForm("autoLogin")
	code := c.PostForm("code")
	mobile := c.PostForm("mobile")
	nick := c.PostForm("nick")
	pwd := c.PostForm("pwd")
	if !checkVerifyCode(code) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "验证码失败，需要返回消息", "msg": "OK"})
		return
	}
	user, error := h.us.RegisterNewUser(mobile, pwd, nick, autoLogin)
	msg := "OK"
	if error != nil {
		msg = error.Error()
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": msg})
}
func (h *UserHandler) UpdateUserInfo(c *gin.Context) {
	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "NoToken", "msg": "OK"})
		return
	}
	nick := c.Query("nick")
	avatarUrl := c.Query("avatarUrl")
	province := c.Query("province")
	city := c.Query("city")
	fmt.Println(token, nick, avatarUrl, province, city)

	// 创建用户数据对象
	userData := User{
		Nickname:  nick,
		AvatarUrl: avatarUrl,
		Province:  province,
		City:      city,
	}

	// 使用新添加的UpdateUserByToken方法更新用户信息
	user := h.us.UpdateUserByToken(token, userData)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": "OK"})
}

func (h *UserHandler) GetDeliveryAddressList(c *gin.Context) {
	token := c.PostForm("token")
	address := h.us.FindUserByToken(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": address, "msg": "OK"})
}

func (h *UserHandler) GetDefaultDeliveryAddress(c *gin.Context) {
	token := c.PostForm("token")
	address := h.us.GetDefaultDeliveryAddress(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": address, "msg": "OK"})
}

func AddDeliveryAddress(c *gin.Context) {
	//TODO:  AddDeliveryAddress
	token := c.PostForm("token")
	address := GetUserService().GetDefaultDeliveryAddress(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": address, "msg": "OK"})
}

func (h *UserHandler) GetUserAmount(c *gin.Context) {
	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "NoToken", "msg": "OK"})
		return
	}
	result := fetchUserAmount(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}


func fetchUserAmount(token string) interface{} {
	//TODO:
	return map[string]string{"token": "fetchUserAmount", "amount": "amount 0"}

}
func (h *UserHandler) GetUserDetail(c *gin.Context) {
	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "NoToken", "msg": "OK"})
		return
	}
	user := h.us.FindUserByToken(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": "OK"})
}
func checkVerifyCode(code string) bool {
	//TODO 需要校验注册的图片验证码
	return true
}
