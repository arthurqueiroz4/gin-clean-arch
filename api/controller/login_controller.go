package controller

import (
	"gin-clean-arch/bootstrap"
	"gin-clean-arch/domain"
	"gin-clean-arch/internal"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	UserRepository domain.UserRepository
	Env            *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := lc.UserRepository.FindByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return
	}
	// if bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(request.Password)) != nil {
	if user.Pass != request.Password {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := internal.CreateAccessToken(user, lc.Env.SecretKey, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := internal.CreateRefreshToken(user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
