package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	deviceId := c.PostForm("deviceId")     //"16533880163937665988"
	deviceName := c.PostForm("deviceName") //"PC"
	mobile := c.PostForm("mobile")         //"13911057997"
	pwd := c.PostForm("pwd")               //"1212121212"

	user := GetUserService().login(deviceId, deviceName, mobile, pwd)

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": "ok"})
}
func Logout(c *gin.Context) {
	token, _ := c.GetQuery("token")
	GetUserService().logout(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(nil), "msg": "OK"})
}
func UpdateUserInfo(c *gin.Context) {
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
	//TODO:
	user := updateUser(token)

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": "OK"})
}

func GetDeliveryAddressList(c *gin.Context) {
	//TODO: GetDeliveryAddressList
	token := c.PostForm("token")
	address := GetUserService().FindUserByToken(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": address, "msg": "OK"})
}

func GetDefaultDeliveryAddress(c *gin.Context) {
	//TODO: Get Default Address
	token := c.PostForm("token")
	address := GetUserService().GetDefaultDeliveryAddress(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": address, "msg": "OK"})
}

func AddDeliveryAddress(c *gin.Context) {
	//TODO:  AddDeliveryAddress
	token := c.PostForm("token")
	address := GetUserService().GetDefaultDeliveryAddress(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": address, "msg": "OK"})
}

func Register(c *gin.Context) {
	autoLogin := c.PostForm("autoLogin") //true
	code := c.PostForm("code")           // "5916"
	mobile := c.PostForm("mobile")       //  "13911057997"
	nick := c.PostForm("nick")           //  "熔岩巨兽"
	pwd := c.PostForm("pwd")             //  "F1ref0x0820"

	if !checkVerifyCode(code) {
		//TODO: 验证码失败，需要返回消息
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "验证码失败，需要返回消息", "msg": "OK"})
		return
	}
	fmt.Printf("autoLogin = %v, code = %v, mobile = %v, nick = '%v, pwd = '%v'\n",
		autoLogin, code, mobile, nick, pwd)
	user := GetUserService().RegisterNewUser(mobile, nick, pwd)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": "OK"})
}

func GetUserAmount(c *gin.Context) {
	//TODO:
	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "NoToken", "msg": "OK"})
		return
	}
	result := fetchUserAmount(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func updateUser(token string) *User {
	//TODO:
	return nil

}

func fetchUserAmount(token string) interface{} {
	//TODO:
	return map[string]string{"token": "fetchUserAmount", "amount": "amount 0"}

}
func GetUserDetail(c *gin.Context) {
	//TODO:
	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "NoToken", "msg": "OK"})
		return
	}

	user := GetUserService().FindUserByToken(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": "OK"})
}
func checkVerifyCode(code string) bool {
	//TODO 需要校验注册的图片验证码
	return true
}
