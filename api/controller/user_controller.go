package controller

import (
	"fmt"
	"gin-clean-arch/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var jsonUser domain.JSONUser
	var err error
	if err = c.ShouldBindJSON(&jsonUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalid"})
		return
	}
	newUser := domain.User{
		Name:  jsonUser.Name,
		Email: jsonUser.Email,
		Pass:  jsonUser.Pass,
	}
	fmt.Println("USER PARSED", newUser)
	userCreated, err := uc.UserUsecase.Create(newUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, userCreated)
}
