package route

import (
	"gin-clean-arch/api/controller"
	"gin-clean-arch/repository"
	"gin-clean-arch/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUser(db *gorm.DB, router *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	uc := usecase.NewUserUsecase(userRepository)

	userController := &controller.UserController{
		UserUsecase: uc,
	}

	router.DELETE("/:id", userController.DeleteUser)
}
