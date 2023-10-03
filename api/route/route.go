package route

import (
	"gin-clean-arch/api/middleware"
	"gin-clean-arch/bootstrap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	//All public APIs
	NewLoginRouter(env, db, publicRouter)
	protectedRouter := gin.Group("")
	//Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuth(env.SecretKey))
	// All private APIs

}
