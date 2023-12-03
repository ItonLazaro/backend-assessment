package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	token "example/todo-go/src/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := token.TokenValid(context)

		if err != nil {
			context.String(http.StatusUnauthorized, "Unauthorized")
			context.Abort()
			return
		}
		context.Next()
	}
}