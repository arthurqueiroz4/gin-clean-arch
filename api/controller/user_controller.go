package controller

import "gin-clean-arch/domain"

type UserController struct {
	UserUsecase domain.UserUsecase
}
