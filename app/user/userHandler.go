package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	result := initUserData()
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func UpdateUserInfo(c *gin.Context) {
	token := c.Param("token")
	nick := c.Param("nick")
	avatarUrl := c.Param("avatarUrl")
	province := c.Param("province")
	city := c.Param("city")
	fmt.Println(token, nick, avatarUrl, province, city)

	result := updateUser(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func GetUserDetail(c *gin.Context) {

	token := c.Param("token")
	result := fetchUserData(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func GetUserAmount(c *gin.Context) {

	token := c.Param("token")
	result := fetchUserAmount(token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": &result, "msg": "OK"})
}

func updateUser(token string) interface{} {
	return map[string]string{"token": "iamToken", "userInfo": "UserInfo 0"}

}

func fetchUserAmount(token string) interface{} {
	return map[string]string{"token": "iamToken", "amount": "amount 0"}

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
	return map[string]string{"token": "iamToken", "userLevel": userLevel, "userInfo": userInfo, "UserState": UserState}

}
func initUserData() interface{} {
	return map[string]string{"token": "iamToken"}
}
