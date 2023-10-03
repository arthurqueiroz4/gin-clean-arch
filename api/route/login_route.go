package route

import (
	"gin-clean-arch/api/controller"
	"gin-clean-arch/bootstrap"
	"gin-clean-arch/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewLoginRouter(env *bootstrap.Env, db *gorm.DB, router *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	loginController := &controller.LoginController{
		UserRepository: userRepository,
		Env:            env,
	}

	router.POST("/login", loginController.Login)
}
