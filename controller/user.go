package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"ticket-backend/common"
	"ticket-backend/dao"
	db "ticket-backend/database"
	"ticket-backend/dto"
	"ticket-backend/model"
	"ticket-backend/response"
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
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "电话号码必须为11位!")
		return
	}

	if len(password) < 6 { // 密码不能小于六位
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码长度必须大于6位!")
		return
	}

	if len(name) == 0 { // 如果没输入用户名，就随机给一个
		name = utils.RandStr(10)
	}

	if dao.IsUserPhoneExist(phone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "电话号码已经存在!")
		return
	}

	if dao.IsUserNameExist(name) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户名已经存在!")
		return
	}
	// 对密码取哈希
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "密码加密错误!")
		return
	}

	// 创建用户
	user := model.User{
		Name:     name,
		Password: string(hashPass),
		Phone:    phone,
	}
	db.DB.Create(&user)
	// 返回结果
	response.Response(c, http.StatusOK, 200, nil, "注册成功!")
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
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "电话号码必须为11位!")
		return
	}

	if len(password) < 6 { // 密码不能小于六位
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码长度必须大于6位!")
		return
	}

	// 判断账户是否存在
	var user model.User
	db.DB.Where("phone=?", phone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在!")
		return
	}

	// 判断密码是否正确

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误!")
		return
	}

	// 发放token
	token, err := common.ReleaseUserToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "token分发失败!")
		log.Println("token generate error:%v", err)
		return
	}
	// 返回结果
	data := gin.H{
		"token": token,
	}
	response.Success(c, data, "用户登录成功！")
	return
}

func UserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	data := gin.H{
		"user": dto.ToUserDto(user.(model.User)),
	}
	response.Success(c, data, "用户信息获取成功！")
}

//
//  UserUpdate
//  @Description: 用户信息修改
//  @param c
//
func UserUpdate(c *gin.Context) {
	// 获取表单内容
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	// 数据合法化验证
	if len(phone) != 0 {
		if len(phone) != 11 {
			response.Response(c, http.StatusUnprocessableEntity, 422, nil, "电话号码必须为11位!")
			return
		}
		if dao.IsUserPhoneExist(phone) {
			response.Response(c, http.StatusUnprocessableEntity, 422, nil, "电话号码已经存在!")
			return
		}
	}
	// 获取用户ID
	user, _ := c.Get("user")
	id := dto.ToUserDto(user.(model.User)).Id
	update := model.User{
		Name:  name,
		Phone: phone,
	}
	dao.UpdateUserById(id, update)
	response.Response(c, http.StatusOK, 200, nil, "用户信息修改成功！")
}

//
//  UserPasswordUpdate
//  @Description: 用户密码修改接口
//
func UserPasswordUpdate(c *gin.Context) {
	password := c.PostForm("password")
	user, _ := c.Get("user")
	id := dto.ToUserDto(user.(model.User)).Id
	if len(password) != 0 {
		if len(password) < 6 { // 密码不能小于六位
			response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码长度必须大于6位!")
			return
		}
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "密码加密错误!")
		return
	}
	update := model.User{
		Password: string(hashPass),
	}
	dao.UpdateUserById(id, update)
	// 返回结果
	response.Response(c, http.StatusOK, 200, nil, "密码修改成功!")
}
