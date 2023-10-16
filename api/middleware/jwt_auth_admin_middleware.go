package middleware

import (
	"gin-clean-arch/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("userContext")
		if !exists {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Internal error"})
			c.Abort()
			return
		}
		user, ok := value.(*domain.User)

		if !ok {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Internal error"})
			c.Abort()
			return
		}

		if user.Role != "ADMIN" {
			c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: "Not authorized"})
			c.Abort()
			return
		}
	}
}
