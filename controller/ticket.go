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

	// 数据封装
	ticket := model.Ticket{
		PassId:  uint(passid),
		PlaneId: uint(planeid),
		Uid:     uid,
		Price:   price,
		Status:  uint(status),
	}
	db.DB.Create(&ticket)
	// 返回结果
	response.Response(c, http.StatusOK, 200, nil, "订单提交成功!")
}

func TicketInfo(c *gin.Context) {

}
