package route

import (
	"gin-clean-arch/api/controller"
	"gin-clean-arch/bootstrap"
	"gin-clean-arch/repository"
	"gin-clean-arch/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewLoginRouter(env *bootstrap.Env, db *gorm.DB, router *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	loginController := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(userRepository, env),
		UserUsecase:  usecase.NewUserUsecase(userRepository),
	}

	router.POST("/login", loginController.Login)
}
