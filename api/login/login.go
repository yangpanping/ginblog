package login

import (
	"github.com/gin-gonic/gin"
	"learngo/goland/yangpanping/bloggin/middleware"
	"learngo/goland/yangpanping/bloggin/models"
	"net/http"
)

//登录
func LoadLogin(c *gin.Context) {
	var u models.Userinform
	u.Usname = c.PostForm("usname")
	u.Uspswd = c.PostForm("uspswd")
	//检查用户是否存在
	u = models.CheckUser(u)
	if u.Usid == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":  "user no exist",
			"username": u.Usname,
			"uspswd":   u.Uspswd,
		})
	} else {
		//产生一个token
		token, _ := middleware.GenToken(u)
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": u.Usname,
			"uspswd":   u.Uspswd,
			"token":    token,
		})
	}

}

//注册
func LoadLogon(c *gin.Context) {
	var u models.Userinform
	u.Usname = c.PostForm("usname")
	u.Uspswd = c.PostForm("uspswd")
	u.Usroot = 3 //默认权限为3
	//把用户添加到数据库中
	code := models.InsertUser(u)
	if code != 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":  "logon fail,usname is exist",
			"username": u.Usname,
			"uspswd":   u.Uspswd,
		})
	} else {
		//产生一个token
		token, _ := middleware.GenToken(u)
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": u.Usname,
			"uspswd":   u.Uspswd,
			"token":    token,
		})
	}
}

//修改密码
func UpdataPswd(c *gin.Context) {

}
