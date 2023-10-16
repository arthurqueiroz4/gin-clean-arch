package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (uc *UserController) FindById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
		c.Abort()
		return
	}
	userFound, err := uc.UserUsecase.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, userFound)
}
