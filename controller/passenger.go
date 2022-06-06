package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ticket-backend/dao"
	db "ticket-backend/database"
	"ticket-backend/dto"
	"ticket-backend/model"
	"ticket-backend/response"
)

//
//  PassengerAdd
//  @Description: 乘客添加接口
//  @param c
//
func PassengerAdd(c *gin.Context) {
	// 表单数据获取
	user, _ := c.Get("user")
	name := c.PostForm("name")
	idnum := c.PostForm("idnum")
	phone := c.PostForm("phone")
	sex := c.PostForm("sex")
	workunit := c.PostForm("workunit")

	// 数据验证
	if len(idnum) != 18 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "身份证号码必须为18位!")
		return
	}

	if len(phone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "电话号码必须为11位!")
		return
	}

	if !(sex == "男" || sex == "女") {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "性别必须为男或女!")
		return
	}

	uid := dto.ToUserDto(user.(model.User)).Id
	passenger := model.Passenger{
		Name:     name,
		Idnum:    idnum,
		Phone:    phone,
		Sex:      sex,
		Workunit: workunit,
		UID:      uid,
	}
	db.DB.Create(&passenger)
	response.Response(c, http.StatusOK, 200, nil, "乘客添加成功")
}

//
//  PassengerList
//  @Description: 获取乘客列表接口
//  @param c
//
func PassengerList(c *gin.Context) {
	user, _ := c.Get("user")
	uid := dto.ToUserDto(user.(model.User)).Id
	passengers := dao.GetPassengerListByUid(uid)
	data := gin.H{
		"passengers": passengers,
	}
	if len(passengers) == 0 {
		response.Response(c, http.StatusNotFound, 404, nil, "列表为空")
	}
	response.Response(c, http.StatusOK, 200, data, "乘客列表获取成功")
}

//
//  PassengerUpdate
//  @Description: 更新乘客信息
//  @param c
//
func PassengerUpdate(c *gin.Context) {
	// 表单数据获取
	user, _ := c.Get("user")
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "id参数错误!")
		return
	}
	var idn = uint(id)
	name := c.PostForm("name")
	idnum := c.PostForm("idnum")
	phone := c.PostForm("phone")
	sex := c.PostForm("sex")
	workunit := c.PostForm("workunit")

	// 数据验证
	if len(idnum) != 0 {
		if len(idnum) != 18 {
			response.Response(c, http.StatusUnprocessableEntity, 422, nil, "身份证号码必须为18位!")
			return
		}
	}
	if len(phone) != 0 {
		if len(phone) != 11 {
			response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位!")
			return
		}
	}
	if len(sex) != 0 {
		if !(sex == "男" || sex == "女") {
			response.Response(c, http.StatusUnprocessableEntity, 422, nil, "性别必须为男或女!")
			return
		}
	}
	// 权限验证
	passenger := dao.GetPassengerById(idn)
	uid := dto.ToUserDto(user.(model.User)).Id
	if uid != passenger.UID {
		response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够！")
		return
	}
	// 修改记录
	update := model.Passenger{
		Name:     name,
		Idnum:    idnum,
		Phone:    phone,
		Sex:      sex,
		Workunit: workunit,
	}
	dao.UpdatePassengerById(idn, update)
	response.Response(c, http.StatusOK, 200, nil, "乘客信息修改成功！")
}

//
//  PassengerDelete
//  @Description: 删除乘客信息
//  @param c
//
func PassengerDelete(c *gin.Context) {
	// 获取数据
	user, _ := c.Get("user")
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "id参数错误!")
		return
	}
	var idn = uint(id)
	// 权限验证
	passenger := dao.GetPassengerById(idn)
	uid := dto.ToUserDto(user.(model.User)).Id
	if uid != passenger.UID {
		response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够！")
		return
	}
	// 删除记录
	dao.DeletePassengerById(idn)
	response.Response(c, http.StatusOK, 200, nil, "乘客信息删除成功！")
}

func PassengerInfo(c *gin.Context) {
	// 获取数据
	user, _ := c.Get("user")
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "id参数错误!")
		return
	}
	var idn = uint(id)
	// 权限验证
	passenger := dao.GetPassengerById(idn)
	uid := dto.ToUserDto(user.(model.User)).Id
	if uid != passenger.UID {
		response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够！")
		return
	}
	// 删除记录
	info := dao.GetPassengerById(idn)
	data := gin.H{
		"info": info,
	}
	response.Response(c, http.StatusOK, 200, data, "乘客信息查询成功！")
}
