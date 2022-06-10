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
	"time"
)

func PlaneCreate(c *gin.Context) {
	// 参数获取
	planenum := c.PostForm("planenum")
	status := c.PostForm("status")
	price, err := strconv.ParseFloat(c.PostForm("price"), 32)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "price参数错误!")
		return
	}

	seat, err := strconv.Atoi(c.PostForm("seat"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "seat参数错误!")
		return
	}
	departure := c.PostForm("departure")
	arrival := c.PostForm("arrival")
	take_off_time, err := strconv.Atoi(c.PostForm("takeofftime"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "takeofftime参数错误!")
		return
	}
	admin, _ := c.Get("admin")
	companyId := dto.ToAdminDto(admin.(model.Admin)).Cid

	// 权限验证
	auth := dto.ToAdminDto(admin.(model.Admin)).Auth
	if auth != 2 {
		response.Response(c, http.StatusUnauthorized, 401, nil, "权限错误，只有公司员工有权限创建航班！")
		return
	}

	// 数据验证
	if dao.IsPlaneNumExist(planenum) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "航班号已经存在!")
		return
	}
	if departure == arrival {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "出发地和目的地不能相同!")
		return
	}
	if seat < 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "座位数必须大于等于0!")
		return
	}
	if price < 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "票价必须大于等于0!")
		return
	}
	timeUnix := time.Now().Unix()
	if int64(take_off_time) < timeUnix {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "起飞时间必须大于当前时间!")
		return
	}

	// 数据封装
	plane := &model.Plane{
		PlaneNum:    planenum,
		Status:      status,
		Price:       float32(price),
		Seat:        seat,
		Departure:   departure,
		Arrival:     arrival,
		TakeoffTime: uint(take_off_time),
		CompanyId:   int(companyId),
	}
	db.DB.Create(&plane)
	response.Response(c, http.StatusOK, 200, nil, "航班创建成功！")
}
