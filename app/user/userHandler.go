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

	if user == nil {
		user = &User{}
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": "ok"})
}
func Logout(c *gin.Context) {
	token, _ := c.GetQuery("token")
	GetUserService().logout(token)
	user := &User{}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": "OK"})
}
func UpdateUserInfo(c *gin.Context) {
	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "NoToken", "msg": "OK"})
		return
	}

	nick := c.Param("nick")
	avatarUrl := c.Param("avatarUrl")
	province := c.Param("province")
	city := c.Param("city")
	fmt.Println(token, nick, avatarUrl, province, city)

	user := updateUser(token)
	if user == nil {
		user = &User{}
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": "OK"})
}
func Register(c *gin.Context) {
	autoLogin := c.PostForm("autoLogin") //true
	code := c.PostForm("code")           // "5916"
	mobile := c.PostForm("mobile")       //  "13911057997"
	nick := c.PostForm("nick")           //  "熔岩巨兽"
	pwd := c.PostForm("pwd")             //  "F1ref0x0820"

	if !checkVerifyCode(code) {
		//TODO 验证码失败，需要返回消息
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "验证码失败，需要返回消息", "msg": "OK"})
		return
	}
	fmt.Printf("autoLogin = %v, code = %v, mobile = %v, nick = '%v, pwd = '%v'\n",
		autoLogin, code, mobile, nick, pwd)
	user := GetUserService().RegisterNewUser(mobile, nick, pwd)
	if user == nil {
		user = &User{}
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": "OK"})
}

func GetUserAmount(c *gin.Context) {

	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "NoToken", "msg": "OK"})
		return
	}
	result := fetchUserAmount(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func updateUser(token string) *User {
	return &User{}

}

func fetchUserAmount(token string) interface{} {
	return &User{}
	//return map[string]string{"token": "fetchUserAmount", "amount": "amount 0"}

}
func GetUserDetail(c *gin.Context) {

	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "NoToken", "msg": "OK"})
		return
	}

	user := GetUserService().findUserByMobile(token)
	if user == nil {
		user = &User{}
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": userToVM(user), "msg": "OK"})
}
func userToVM(user *User) UserVM {
	return UserVM{
		user.Mobile,
		*user,
		user.UserLevel,
	}
}
func checkVerifyCode(code string) bool {
	//TODO 需要校验注册的图片验证码
	return true
}
