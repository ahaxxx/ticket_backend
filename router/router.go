package router

import (
	"github.com/gin-gonic/gin"
	"ticket-backend/controller"
	"ticket-backend/middlewave"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	v1 := r.Group("/api/v1")
	{
		v1.POST("user/register", controller.UserRegister)                         // 用户注册接口
		v1.POST("user/login", controller.UserLogin)                               // 用户登录接口
		v1.GET("user/info", middlewave.UserAuthMiddleware(), controller.UserInfo) // 用户信息接口

		v1.POST("passenger/add", middlewave.UserAuthMiddleware(), controller.PassengerAdd)  // 添加乘客
		v1.GET("passenger/list", middlewave.UserAuthMiddleware(), controller.PassengerList) // 获取乘客列表
	}
	return r
}
