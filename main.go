package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naufaldymahas/go-gin-jwt/src/config"
	"github.com/naufaldymahas/go-gin-jwt/src/model"
	"github.com/naufaldymahas/go-gin-jwt/src/router"
)

func migrate() {
	db := config.InitDB()
	defer db.Close()

	db.Debug().AutoMigrate(&model.User{})
}

func main() {
	migrate()

	r := gin.Default()

	router.AuthRouter(r)

	r.Run()
}
