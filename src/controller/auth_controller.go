package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufaldymahas/go-gin-jwt/src/config"
	"github.com/naufaldymahas/go-gin-jwt/src/payload"
	"github.com/naufaldymahas/go-gin-jwt/src/service"
)

func Register(c *gin.Context) {
	var auth payload.AuthRegister
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	us := service.UserService(config.InitDB())

	if err := us.CreateUser(auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func Login(c *gin.Context) {
	var auth payload.AuthLogin
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	us := service.UserService(config.InitDB())

	u, ok := us.LoginUser(auth)
	if ok == false {
		c.String(http.StatusUnauthorized, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
