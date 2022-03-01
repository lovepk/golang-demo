package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo_list/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code = 200
		token := context.GetHeader("authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = 403 // 无权限 token错误
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 401
			}
		}
		if code != 200 {
			context.JSON(200, gin.H{
				"status": code,
				"msg": "token 解析错误",
			})
			context.Abort()
			return
		}
		context.Next()
	}
}