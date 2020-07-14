package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufaldymahas/go-gin-jwt/src/config"
	"github.com/naufaldymahas/go-gin-jwt/src/service"
)

func GetUserByEmail(c *gin.Context) {
	email := c.Query("email")
	us := service.UserService(config.InitDB())
	defer us.DB.Close()

	user, err := us.FindByEmail(email)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
