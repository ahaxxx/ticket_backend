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
	// 数据上传
	db.DB.Create(&plane)
	response.Response(c, http.StatusOK, 200, nil, "航班创建成功！")
}

func PlaneUpdate(c *gin.Context) {
	// 参数获取
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "id参数错误!")
		return
	}
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

	// 参数验证
	if dao.IsPlaneNumExist(planenum) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "航班号已经存在!")
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
	// 权限验证
	plane := dao.GetPlaneById(uint(id))
	cid := dto.ToAdminDto(admin.(model.Admin)).Cid
	if cid != uint(plane.CompanyId) {
		response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够！")
		return
	}
	// 数据封装
	update := model.Plane{
		PlaneNum:    planenum,
		Status:      status,
		Price:       float32(price),
		Seat:        seat,
		Departure:   departure,
		Arrival:     arrival,
		TakeoffTime: uint(take_off_time),
	}
	// 修改记录
	dao.UpdatePlaneById(uint(id), update)
	response.Response(c, http.StatusOK, 200, nil, "航班信息修改成功！")
}

func PlaneList(c *gin.Context) {
	planes := dao.GetPlaneList()
	data := gin.H{
		"planes": planes,
	}
	if len(planes) == 0 {
		response.Response(c, http.StatusNotFound, 404, nil, "列表为空")
	}
	response.Response(c, http.StatusOK, 200, data, "航班列表获取成功")
}

func PlaneSearch(c *gin.Context) {
	departure := c.PostForm("departure")
	arrival := c.PostForm("arrival")
	take_off_time, err := strconv.Atoi(c.PostForm("takeofftime"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "takeofftime参数错误!")
		return
	}
	plane := dao.SearchPlaneByDAT(departure, arrival, uint(take_off_time))
	data := gin.H{
		"plane": plane,
	}
	response.Response(c, http.StatusOK, 200, data, "航班信息查询成功！")
}

func PlaneInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("pid"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "id参数错误!")
		return
	}
	var pid = uint(id)

	plane := dao.GetPlaneById(pid)
	data := gin.H{
		"plane": plane,
	}
	response.Response(c, http.StatusOK, 200, data, "航班信息查询成功！")
}

func PlaneDelete(c *gin.Context) {
	admin, _ := c.Get("admin")
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "id参数错误!")
		return
	}
	var idn = uint(id)
	// 权限验证
	plane := dao.GetPlaneById(idn)
	cid := dto.ToAdminDto(admin.(model.Admin)).Cid
	if cid != uint(plane.CompanyId) {
		response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够！")
		return
	}
	// 删除记录
	dao.DeletePlaneById(idn)
	response.Response(c, http.StatusOK, 200, nil, "航班信息删除成功！")
}
