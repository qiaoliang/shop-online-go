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

	fmt.Printf("mobile = %v,  pwd=%v\n\n\n", mobile, pwd)
	found := GetUserService().login(deviceId, deviceName, mobile, pwd)
	var result map[string]string
	msg := "OK"
	code := 0
	if found == nil {
		result = map[string]string{"token": "UserLogin"}
		msg = "用户名或密码错误"
	} else {
		result = map[string]string{"token": found.Mobile}
	}
	fmt.Println(result["token"])
	c.JSON(http.StatusOK, gin.H{"code": code, "data": &result, "msg": msg})
}
func Logout(c *gin.Context) {
	token, _ := c.GetQuery("token")
	GetUserService().logout(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": "", "msg": "OK"})
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

	result := updateUser(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}
func Register(c *gin.Context) {
	autoLogin := c.PostForm("autoLogin") //true
	code := c.PostForm("code")           // "5916"
	mobile := c.PostForm("mobile")       //  "13911057997"
	nick := c.PostForm("nick")           //  "熔岩巨兽"
	pwd := c.PostForm("pwd")             //  "F1ref0x0820"

	fmt.Printf("autoLogin = %v, code = %v, mobile = %v, nick = '%v, pwd = '%v'\n",
		autoLogin, code, mobile, nick, pwd)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": "", "msg": "OK"})

}
func GetUserDetail(c *gin.Context) {

	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "NoToken", "msg": "OK"})
		return
	}

	result := fetchUserData(token)
	//base = {}, userLevel = {}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
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

func updateUser(token string) interface{} {
	return map[string]string{"token": "updateUser", "userInfo": "UserInfo 0"}

}

func fetchUserAmount(token string) interface{} {
	return map[string]string{"token": "fetchUserAmount", "amount": "amount 0"}

}
func fetchUserData(token string) interface{} {
	base := "1"

	userLevel := "1"
	// id?: number;
	// [key: string]: any;

	userInfo := base
	// id?: number;
	//nick?: string;
	//avatar?: string;
	//[key: string]: any;
	UserState := "0"
	//token: string;
	//userInfo: NonNullable<UserInfo>;
	//userLevel: NonNullable<UserLevel>;
	return map[string]string{"token": token, "userLevel": userLevel, "userInfo": userInfo, "UserState": UserState}

}
