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

func TicketCreate(c *gin.Context) {
	// 参数获取
	user, _ := c.Get("user")
	passid, err := strconv.Atoi(c.PostForm("passid"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "passid参数错误!")
		return
	}
	planeid, err := strconv.Atoi(c.PostForm("planeid"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "planeid参数错误!")
		return
	}
	uid := dto.ToUserDto(user.(model.User)).Id
	plane := dao.GetPlaneById(uint(planeid))
	price := plane.Price
	status := 1
	if plane.Seat <= 0 {
		response.Response(c, http.StatusNotFound, 404, nil, "该航班已售罄!")
		return
	}
	Seat := plane.Seat - 1
	// 数据封装
	ticket := model.Ticket{
		PassId:  uint(passid),
		PlaneId: uint(planeid),
		Uid:     uid,
		Price:   price,
		Status:  uint(status),
	}

	update := model.Plane{
		Seat: Seat,
	}
	db.DB.Create(&ticket)
	dao.UpdatePlaneById(uint(planeid), update)
	// 返回结果
	response.Response(c, http.StatusOK, 200, nil, "订单提交成功!")
}

func TicketList(c *gin.Context) {
	user, _ := c.Get("user")
	uid := dto.ToUserDto(user.(model.User)).Id
	tickets := dao.GetTicketListByUid(uid)
	data := gin.H{
		"tickets": tickets,
	}
	if len(tickets) == 0 {
		response.Response(c, http.StatusNotFound, 404, nil, "列表为空")
	}
	response.Response(c, http.StatusOK, 200, data, "航班列表获取成功")
}

func TicketConfirm(c *gin.Context) {
	// 获取参数
	admin, _ := c.Get("admin")
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "id参数错误!")
		return
	}
	// 权限验证
	cid := dto.ToAdminDto(admin.(model.Admin)).Cid
	ticket := dao.GetTicketById(uint(id))
	plane := dao.GetPlaneById(ticket.PlaneId)
	if cid != uint(plane.CompanyId) {
		response.Response(c, http.StatusUnprocessableEntity, 401, nil, "权限错误!")
		return
	}
	// 数据封装
	update := model.Ticket{
		Status: uint(2),
	}
	// 数据更新
	dao.UpdateTicketById(uint(id), update)
	response.Response(c, http.StatusOK, 200, nil, "订单确认出票成功!")
}

func TicketCancel(c *gin.Context) {
	user, _ := c.Get("user")
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "id参数错误!")
		return
	}
	// 权限验证
	uid := dto.ToUserDto(user.(model.User)).Id
	ticket := dao.GetTicketById(uint(id))
	userid := ticket.Uid
	if uid != userid {
		response.Response(c, http.StatusUnprocessableEntity, 401, nil, "权限错误!")
		return
	}
	// 数据封装
	update := model.Ticket{
		Status: uint(3),
	}
	// 数据更新
	dao.UpdateTicketById(uint(id), update)
	response.Response(c, http.StatusOK, 200, nil, "订单取消成功!")
}
