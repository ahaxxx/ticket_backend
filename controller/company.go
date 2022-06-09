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
//  CreateCompany
//  @Description: 添加公司接口
//  @param c
//
func CompanyCreate(c *gin.Context) {
	// 参数获取
	admin, _ := c.Get("admin")
	name := c.PostForm("name")

	// 参数验证
	if len(name) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "公司名不能为空!")
		return
	}

	if dao.IsCompanyExist(name) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "公司已经存在!")
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

func CompanyList(c *gin.Context) {
	// 参数获取
	admin, _ := c.Get("admin")
	auth := dto.ToAdminDto(admin.(model.Admin)).Auth
	if auth != 1 {
		response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够，只有超级管理员有权限查看公司列表！")
		return
	}
	company := dao.GetCompanyList()
	data := gin.H{
		"companies": company,
	}
	if len(company) == 0 {
		response.Response(c, http.StatusNotFound, 404, nil, "列表为空")
	}
	response.Response(c, http.StatusOK, 200, data, "公司列表获取成功")
}

func CompanyInfo(c *gin.Context) {
	// 参数获取
	admin, _ := c.Get("admin")
	cid, err := strconv.Atoi(c.PostForm("cid"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "cid参数错误!")
		return
	}
	// 权限验证
	auth := dto.ToAdminDto(admin.(model.Admin)).Auth
	if auth != 1 {
		response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够，只有超级管理员有权限查看公司信息！")
		return
	}
	company := dao.GetCompanyByCid(uint(cid))
	data := gin.H{
		"company": company,
	}
	response.Success(c, data, "用户信息获取成功！")
}

func CompanyUpdate(c *gin.Context) {
	// 参数获取
	admin, _ := c.Get("admin")
	cid, err := strconv.Atoi(c.PostForm("cid"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "id参数错误!")
		return
	}
	name := c.PostForm("name")

	// 权限验证
	auth := dto.ToAdminDto(admin.(model.Admin)).Auth
	if auth != 1 {
		response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够，只有超级管理员有权限修改公司信息！")
		return
	}
	update := model.Company{
		Name: name,
	}
	dao.UpdateCompanyById(uint(cid), update)
	response.Response(c, http.StatusOK, 200, nil, "公司信息修改成功！")
}

func CompanyDelete(c *gin.Context) {
	// 参数获取
	admin, _ := c.Get("admin")
	cid, err := strconv.Atoi(c.PostForm("cid"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "id参数错误!")
		return
	}

	// 权限验证
	auth := dto.ToAdminDto(admin.(model.Admin)).Auth
	if auth != 1 {
		response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够，只有超级管理员有权限修改公司信息！")
		return
	}

	// 删除
	dao.DeleteCompanyById(uint(cid))
	response.Response(c, http.StatusOK, 200, nil, "公司信息删除成功！")
}
