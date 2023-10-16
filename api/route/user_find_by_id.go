package route

import (
	"gin-clean-arch/api/controller"
	"gin-clean-arch/repository"
	"gin-clean-arch/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindUserById(db *gorm.DB, router *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	userController := controller.UserController{UserUsecase: userUsecase}

	router.GET("/:id", userController.FindById)
}
