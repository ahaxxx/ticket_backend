package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"ticket-backend/common"
	"ticket-backend/dao"
	db "ticket-backend/database"
	"ticket-backend/dto"
	"ticket-backend/model"
	"ticket-backend/response"
)

//
//  AdminRegister
//  @Description: 管理员注册
//  @param c
//
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
	cid, err := strconv.Atoi(c.PostForm("cid"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "公司参数错误!")
		return
	}

	status, err := strconv.Atoi(c.PostForm("status"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "公司参数错误!")
		return
	}

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
	if dao.IsAdminPhoneExist(phone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "电话号码已经存在!")
		return
	}

	if dao.IsAdminNameExist(name) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户名已经存在!")
		return
	}

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
		Cid:      uint(cid),
		Status:   uint(status),
	}
	db.DB.Create(&admin)
	// 返回结果
	response.Response(c, http.StatusOK, 200, nil, "注册成功!")
}

//
//  AdminLogin
//  @Description: 管理员登录
//  @param c
//
func AdminLogin(c *gin.Context) {

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
	var admin model.Admin
	db.DB.Where("phone=?", phone).First(&admin)
	if admin.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在!")
		return
	}

	// 判断密码是否正确

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))

	if err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误!")
		return
	}

	// 判断用户状态

	if admin.Status != 1 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "该后台账户被禁用!")
		return
	}

	// 发放token
	token, err := common.ReleaseAdminToken(admin)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "token分发失败!")
		log.Println("token generate error:%v", err)
		return
	}
	// 返回结果
	data := gin.H{
		"token": token,
	}
	response.Success(c, data, "后台登录成功！")
	return
}

//
//  AdminInfo
//  @Description: 获取管理员信息
//  @param c
//
func AdminInfo(c *gin.Context) {
	admin, _ := c.Get("admin")
	data := gin.H{
		"admin": dto.ToAdminDto(admin.(model.Admin)),
	}
	response.Success(c, data, "用户信息获取成功！")
}

//
//  AdminUpdate
//  @Description: 修改管理员信息
//  @param c
//
func AdminUpdate(c *gin.Context) {
	// 获取表单内容
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	auth, err := strconv.Atoi(c.PostForm("auth"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "权限参数错误!")
		return
	}
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
	admin, _ := c.Get("admin")
	id := dto.ToAdminDto(admin.(model.Admin)).Id
	update := model.Admin{
		Name:  name,
		Phone: phone,
		Auth:  uint(auth),
	}
	dao.UpdateAdminById(id, update)
	response.Response(c, http.StatusOK, 200, nil, "管理员信息修改成功！")
}

//
//  AdminPasswordUpdate
//  @Description: 修改管理员密码
//  @param c
//
func AdminPasswordUpdate(c *gin.Context) {
	password := c.PostForm("password")
	admin, _ := c.Get("admin")
	id := dto.ToAdminDto(admin.(model.Admin)).Id
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
	update := model.Admin{
		Password: string(hashPass),
	}
	dao.UpdateAdminById(id, update)
	// 返回结果
	response.Response(c, http.StatusOK, 200, nil, "密码修改成功!")
}
