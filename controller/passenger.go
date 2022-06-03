package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
