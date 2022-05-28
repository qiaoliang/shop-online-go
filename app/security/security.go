package security

import (
	"fmt"
	"image/color"
	"net/http"

	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
)

func GetSMSCode(c *gin.Context) {

	mobile, _ := c.GetQuery("mobile")
	key, _ := c.GetQuery("key")
	picCode, _ := c.GetQuery("picCode")
	token, _ := c.GetQuery("token")

	fmt.Printf("mobile = %v, key = %v, picCode = %v, token = '%v'\n", mobile, key, picCode, token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": 1, "msg": "OK"})
}
func VerifyCapChar(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": 1, "msg": "OK"})
}
func GetCapChar(c *gin.Context) {
	//key := c.DefaultQuery("key", "noKey") // key is a timestump from client
	pic := "http://localhost:9090/pic/captcha.jpeg"
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": pic, "msg": "OK"})
}

func GenerateCapTCha() {
	cap := captcha.New()
	// 可以设置多个字体 或使用cap.AddFont("xx.ttf")追加更多
	cap.SetFont("resources/comic.ttf")
	// 设置验证码大小
	cap.SetSize(128, 64)
	// 设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// 设置背景色 可以多个 随机替换背景色 默认白色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	//img, str := cap.Create(4, captcha.NUM)
	//img1, str1 := cap.Create(6, captcha.ALL)
	//png.Encode(w, img)
}
