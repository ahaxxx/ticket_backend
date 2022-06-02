package router

import (
	"github.com/gin-gonic/gin"
	"ticket-backend/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	v1 := r.Group("/api/v1")
	{
		v1.POST("user/register", controller.UserRegister) // 用户注册接口
		v1.POST("user/login", controller.UserLogin)       // 用户注册接口
	}
	return r
}
