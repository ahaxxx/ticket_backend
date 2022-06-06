package middlewave

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"ticket-backend/common"
	db "ticket-backend/database"
	"ticket-backend/model"
	"ticket-backend/response"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够！")
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseUserTokenString(tokenString)
		if err != nil || !token.Valid {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够！")
			c.Abort()
			return
		}
		// 验证通过之后获取claim中的userID
		userId := claims.UserId
		var user model.User
		db.DB.First(&user, userId)
		if user.ID == 0 {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够！")
			c.Abort()
			return
		}
		// 用户存在，将user信息写入context
		c.Set("user", user)
		c.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够！")
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseAdminTokenString(tokenString)
		if err != nil || !token.Valid {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够！")
			c.Abort()
			return
		}
		// 验证通过之后获取claim中的userID
		adminId := claims.AdminId
		var admin model.Admin
		db.DB.First(&admin, adminId)
		if admin.ID == 0 {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不够！")
			c.Abort()
			return
		}
		// 用户存在，将user信息写入context
		c.Set("admin", admin)
		c.Next()
	}
}
