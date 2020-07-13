package router

import (
	"github.com/gin-gonic/gin"
	"github.com/naufaldymahas/go-gin-jwt/src/controller"
)

func AuthRouter(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
	}
}
