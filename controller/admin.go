package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	db "ticket-backend/database"
	"ticket-backend/model"
	"ticket-backend/response"
)

func AdminRegister(c *gin.Context) {

	// 获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	auth, err := strconv.Atoi(c.PostForm("auth"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "权限参数错误!")
		return
	}
	company := c.PostForm("company")

	// 数据验证
	if len(phone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "电话号码必须为11位!")
		return
	}

	if len(password) < 6 { // 密码不能小于六位
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码长度必须大于6位!")
		return
	}

	if len(name) == 0 { // 如果没输入用户名，就随机给一个
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "请输入用户名!")
		return
	}
	// 数据存在性验证
	/**
	if dao.IsPhoneExist(phone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "电话号码已经存在!")
		return
	}

	if dao.IsNameExist(name) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户名已经存在!")
		return
	}
	**/
	// 对密码取哈希
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "密码加密错误!")
		return
	}

	// 创建后台用户
	admin := model.Admin{
		Name:     name,
		Password: string(hashPass),
		Phone:    phone,
		Auth:     uint(auth),
		Company:  company,
	}
	db.DB.Create(&admin)
	// 返回结果
	response.Response(c, http.StatusOK, 200, nil, "注册成功!")
}
