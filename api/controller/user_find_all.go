package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (uc *UserController) FindAllUsers(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		page = "0"
	}
	size := c.Query("size")
	if size == "" {
		size = "10"
	}
	allUsers, total, err := uc.UserUsecase.FindAll(page, size)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	pagination := make(map[string]interface{})
	pagination["page"] = page
	pagination["size"] = size
	pagination["total"] = total
	c.JSON(http.StatusOK, gin.H{
		"pagination": pagination,
		"content":    allUsers,
	})
}
