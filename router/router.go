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
		v1.POST("user/register", controller.UserRegister)                                       // 用户注册接口
		v1.POST("user/login", controller.UserLogin)                                             // 用户登录接口
		v1.GET("user/info", middlewave.UserAuthMiddleware(), controller.UserInfo)               // 用户信息接口
		v1.PUT("user", middlewave.UserAuthMiddleware(), controller.UserUpdate)                  // 用户信息修改接口
		v1.PUT("user/password", middlewave.UserAuthMiddleware(), controller.UserPasswordUpdate) // 用户密码修改接口

		v1.POST("passenger", middlewave.UserAuthMiddleware(), controller.PassengerAdd)       // 添加乘客
		v1.POST("passenger/list", middlewave.UserAuthMiddleware(), controller.PassengerList) // 获取乘客列表
		v1.POST("passenger/info", middlewave.UserAuthMiddleware(), controller.PassengerInfo) // 获取乘客列表
		v1.PUT("passenger", middlewave.UserAuthMiddleware(), controller.PassengerUpdate)     // 修改乘客信息
		v1.DELETE("passenger", middlewave.UserAuthMiddleware(), controller.PassengerDelete)  // 删除乘客

		v1.POST("admin/register", middlewave.AdminAuthMiddleware(), controller.AdminRegister)      // 注册后台用户
		v1.POST("admin/login", controller.AdminLogin)                                              // 后台用户登录
		v1.GET("admin", middlewave.AdminAuthMiddleware(), controller.AdminInfo)                    // 用户信息接口
		v1.PUT("admin", middlewave.AdminAuthMiddleware(), controller.AdminUpdate)                  // 后台密码修改接口
		v1.PUT("admin/password", middlewave.AdminAuthMiddleware(), controller.AdminPasswordUpdate) // 后台密码修改接口

		v1.POST("company", middlewave.AdminAuthMiddleware(), controller.CreateCompany)
	}
	return r
}
