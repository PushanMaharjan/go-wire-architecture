package user

import (
	"github.com/gin-gonic/gin"
	"go-wire-architecture/internal/domain"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (uc *Handler) Save(c *gin.Context) {
	var u domain.Users

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// u, err := uc.userService.Save(c.Request.Context(), u)
	//
	// if err != nil {
	//   c.AbortWithStatus(http.StatusNotFound)
	// } else {

	c.JSON(200, gin.H{"data": "u created"})

	// }
}
