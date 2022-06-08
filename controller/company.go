package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "ticket-backend/database"
	"ticket-backend/dto"
	"ticket-backend/model"
	"ticket-backend/response"
)

func CreateCompany(c *gin.Context) {
	// 参数获取
	admin, _ := c.Get("admin")
	name := c.PostForm("name")

	// 参数验证
	if len(name) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "公司名不能为空!")
		return
	}
	// 权限验证
	auth := dto.ToAdminDto(admin.(model.Admin)).Auth

	if auth != 1 {
		response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够，只有超级管理员有权限创建公司！")
		return
	}

	company := model.Company{
		Name: name,
	}
	db.DB.Create(&company)
	response.Response(c, http.StatusOK, 200, nil, "公司添加成功")
}
