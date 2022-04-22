package routers

import (
	"github.com/gin-gonic/gin"
	"learngo/goland/yangpanping/bloggin/api/login"
	"learngo/goland/yangpanping/bloggin/middleware"
)

//路由初始化
func GinInit() {
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		//登录
		userGroup.POST("/login", middleware.JwtToken(), login.LoadLogin)
		//注册
		userGroup.POST("/logon", login.LoadLogon)
		//修改密码
		userGroup.POST("/updatapswd", login.UpdataPswd)

	}
	r.Run(":8080")
}
