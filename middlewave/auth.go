package middlewave

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"ticket-backend/common"
	db "ticket-backend/database"
	"ticket-backend/model"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不够",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseUserTokenString(tokenString)
		if err != nil || token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不够",
			})
			c.Abort()
			return
		}
		// 验证通过之后获取claim中的userID
		userId := claims.UserId
		var user model.User
		db.DB.First(&user, userId)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不够",
			})
			c.Abort()
			return
		}
		// 用户存在，将user信息写入context
		c.Set("user", user)
		c.Next()
	}
}
