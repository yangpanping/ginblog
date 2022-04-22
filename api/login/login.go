package login

import (
	"fmt"
	"learngo/goland/yangpanping/bloggin/middleware"

	"github.com/gin-gonic/gin"
	"learngo/goland/yangpanping/bloggin/models"
	"net/http"
)

//登录
func LoadLogin(c *gin.Context) {
	var u models.Userinform
	u.Usname = c.PostForm("usname")
	u.Uspswd = c.PostForm("uspswd")
	//检查用户是否存在
	root := models.CheckUser(u)
	//产生一个token
	token, _ := middleware.GenToken(u.Usname)
	//解析一个token
	info, _ := middleware.ParseToken(token)
	fmt.Println(info)
	//输出json结果给调用方
	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"username": u.Usname,
		"uspswd":   u.Uspswd,
		"root":     root,
		"token":    token,
	})
}

//注册
func LoadLogon(c *gin.Context) {

}

//修改密码
func UpdataPswd(c *gin.Context) {

}
