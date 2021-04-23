package controllers

import (
	"goUnitTest/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	result, err := services.PingService.HandlePing()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(http.StatusOK, result)
}
