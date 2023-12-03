package middlewares

import (
	"net/http"

	token "example/todo-go/src/utils"

	"github.com/gin-gonic/gin"
)

// Check if JWT is valid
// else returns "Unauthorized"
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
