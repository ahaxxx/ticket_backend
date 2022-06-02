package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"ticket-backend/dao"
	db "ticket-backend/database"
	"ticket-backend/model"
	"ticket-backend/utils"
)

//
//  UserRegister
//  @Description:用户注册接口实现
//  @param c
//
func UserRegister(c *gin.Context) {

	// 获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	// 数据验证
	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "电话号码必须为11位!",
		})
		return
	}

	if len(password) < 6 { // 密码不能小于六位
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码长度必须大于6位!",
		})
		return
	}

	if len(name) == 0 { // 如果没输入用户名，就随机给一个
		name = utils.RandStr(10)
	}
	log.Println(name, password, phone)

	if dao.IsPhoneExist(phone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "电话号码已经存在!",
		})
		return
	}

	if dao.IsNameExist(name) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户名已经存在!",
		})
		return
	}
	// 对密码取哈希
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "密码加密错误",
		})
	}

	// 创建用户
	user := model.User{
		Name:     name,
		Password: string(hashPass),
		Phone:    phone,
	}
	db.DB.Create(&user)
	// 返回结果

	c.JSON(utils.NewSucc("Register Success!", gin.H{
		"msg": "注册成功!",
	}))
}

//
//  UserLogin
//  @Description:用户登录接口
//  @param c
//
func UserLogin(c *gin.Context) {

	// 获取参数
	phone := c.PostForm("phone")
	password := c.PostForm("password")

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

	// 判断账户是否存在
	if dao.IsPhoneExist(phone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "Phone existed!",
		})
		return
	}

	// 判断密码是否正确

	// 发放token

	// 返回结果
}
