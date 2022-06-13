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

		v1.POST("company", middlewave.AdminAuthMiddleware(), controller.CompanyCreate)    // 公司创建接口
		v1.GET("company/list", middlewave.AdminAuthMiddleware(), controller.CompanyList)  // 获取公司列表接口，仅有超级管理员有权限
		v1.POST("company/info", middlewave.AdminAuthMiddleware(), controller.CompanyInfo) // 获取公司信息接口
		v1.PUT("company", middlewave.AdminAuthMiddleware(), controller.CompanyUpdate)     // 修改公司信息接口
		v1.DELETE("company", middlewave.AdminAuthMiddleware(), controller.CompanyDelete)  // 删除公司接口

		v1.POST("plane", middlewave.AdminAuthMiddleware(), controller.PlaneCreate)   // 创建航班接口，仅有后台用户有权限创建
		v1.POST("plane/info", controller.PlaneInfo)                                  // 查看航班信息
		v1.POST("plane/search", controller.PlaneSearch)                              // 查询航班接口
		v1.GET("plane/list", controller.PlaneList)                                   // 获取航班列表接口
		v1.DELETE("plane", middlewave.AdminAuthMiddleware(), controller.PlaneDelete) // 删除航班接口
		v1.PUT("plane", middlewave.AdminAuthMiddleware(), controller.PlaneUpdate)    // 修改航班信息接口

		v1.POST("ticket", middlewave.UserAuthMiddleware(), controller.TicketCreate)           // 创建订单接口
		v1.GET("ticket/list", middlewave.UserAuthMiddleware(), controller.TicketList)         // 获取订单列表
		v1.POST("ticket/confirm", middlewave.AdminAuthMiddleware(), controller.TicketConfirm) // 确认订单出票
		v1.POST("ticket/cancel", middlewave.UserAuthMiddleware(), controller.TicketCancel)    // 订单取消
		// 需要添加航班模块，机票订单模块
		// 技术难点是需要学习一下多表联动问题
	}
	return r
}
