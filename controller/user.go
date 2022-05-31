package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ticket-backed/utils"
)

func UserRegister(c *gin.Context) {

	// 获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	// 数据验证
	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The phone number must be 11 digits!",
		})
		return
	}

	if len(password) < 6 { // 密码不能小于六位
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The password must be longer than 6 digits!",
		})
		return
	}

	if len(name) == 0 { // 如果没输入用户名，就随机给一个
		name = utils.RandStr(10)
	}
	log.Println(name, password, phone)
	// 创建用户

	// 返回结果

	c.JSON(utils.NewSucc("Register Success!", gin.H{
		"msg": "Register Success!",
	}))
}
