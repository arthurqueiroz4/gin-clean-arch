package middleware

import (
	"gin-clean-arch/domain"
	"gin-clean-arch/internal"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAuth(secret string) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.Request.Header.Get("Authorization")
		token := strings.Split(authHeader, " ")
		if len(token) == 2 {
			authToken := token[1]
			authorized, err := internal.IsAuthorized(authToken, secret)
			if authorized {
				userId, err := internal.ExtractIDFromToken(authToken, secret)
				if err != nil {
					context.JSON(http.StatusUnauthorized, domain.ErrorResponse{err.Error()})
					context.Abort()
					return
				}
				context.Set("x-user-id", userId)
				context.Next()
				return
			}
			context.JSON(http.StatusUnauthorized, domain.ErrorResponse{err.Error()})
			context.Abort()
		}
		context.JSON(http.StatusUnauthorized, domain.ErrorResponse{"Not authorized"})
		context.Abort()
	}
}
