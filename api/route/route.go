package route

import (
	"gin-clean-arch/api/middleware"
	"gin-clean-arch/bootstrap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, db *gorm.DB, gin *gin.Engine) {
	//All public APIs
	publicRouter := gin.Group("")
	NewLoginRouter(env, db, publicRouter)

	// All private APIs
	protectedRouter := gin.Group("")
	//Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuth(env.SecretKey, db))
	protectedRouter.Use(middleware.AdminAuth())
	CreateUser(db, protectedRouter)
	DeleteUser(db, protectedRouter)
	FindAllUsers(db, protectedRouter)
	FindUserById(db, protectedRouter)

}
