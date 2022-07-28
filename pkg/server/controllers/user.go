package controllers

import (
	"go-wire-architecture/pkg/handler/services"
	"go-wire-architecture/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (uc *UserController) Save(c *gin.Context) {
	var user model.Users

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := uc.userService.Save(c.Request.Context(), user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, gin.H{"data": "user created"})
	}
}
