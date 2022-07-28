package web

import (
    "github.com/gin-gonic/gin"
    "go-wire-architecture/internal/user"
    "net/http"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
    return &UserHandler{
	  
    }
}

func (uc *UserHandler) Save(c *gin.Context) {
    var u user.Users
    
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
