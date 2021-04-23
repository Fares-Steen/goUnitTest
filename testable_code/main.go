package main

import (
	"goUnitTest/controllers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", controllers.Ping)
	router.Run(":8080")
}
