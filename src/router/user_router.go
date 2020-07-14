package router

import (
	"github.com/gin-gonic/gin"
	"github.com/naufaldymahas/go-gin-jwt/src/controller"
)

func UserRouter(r *gin.RouterGroup) *gin.RouterGroup {
	user := r.Group("/user")
	{
		user.GET("/", controller.GetUserByEmail)
	}
	return user
}
