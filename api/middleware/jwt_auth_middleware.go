package middleware

import (
	"gin-clean-arch/domain"
	"gin-clean-arch/internal"
	"gin-clean-arch/repository"
	"gin-clean-arch/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

func JwtAuth(secret string, db *gorm.DB) gin.HandlerFunc {

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	return func(context *gin.Context) {
		authHeader := context.Request.Header.Get("Authorization")
		token := strings.Split(authHeader, " ")
		if len(token) == 2 {
			authToken := token[1]
			authorized, err := internal.IsAuthorized(authToken, secret)
			if authorized {
				userIdString, err := internal.ExtractIDFromToken(authToken, secret)
				if err != nil {
					context.JSON(http.StatusUnauthorized, domain.ErrorResponse{err.Error()})
					context.Abort()
					return
				}

				userContext, err := getUserById(userIdString, userUsecase)
				if err != nil {
					context.JSON(http.StatusUnauthorized, domain.ErrorResponse{err.Error()})
					context.Abort()
					return
				}
				context.Set("userContext", userContext)
				context.Next()
				return
			}
			context.JSON(http.StatusUnauthorized, domain.ErrorResponse{err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusUnauthorized, domain.ErrorResponse{"Not authorized"})
		context.Abort()
		return
	}
}

func getUserById(id string, userUsecase domain.UserUsecase) (*domain.User, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	userFound, err := userUsecase.FindById(uint(idInt))
	if err != nil {
		return nil, err
	}
	return userFound, nil
}
