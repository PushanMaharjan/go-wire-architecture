package controllers

import (
	"go-wire-architecture/pkg/handler/services"
	"go-wire-architecture/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SecondUserController struct {
	userService services.UserService
}

func NewSecondUserController(service services.UserService) *SecondUserController {
	return &SecondUserController{
		userService: service,
	}
}

func (uc *SecondUserController) Save(c *gin.Context) {
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
