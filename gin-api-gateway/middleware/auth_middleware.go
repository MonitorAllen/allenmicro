package middleware

import (
	"allenmicro/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" || tokenString[0:7] != "Bearer " {
			context.JSON(http.StatusUnauthorized, gin.H{
				"message": "用户验证失败",
			})
			context.Abort()
			return
		}

		token, claims, err := common.ParseToken(tokenString[7:])
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"message": "用户验证失败",
			})
			context.Abort()
			return
		}

		context.Set("user", claims.Username)
		context.Next()
	}
}
