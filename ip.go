package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func rspIP(c *gin.Context) {
	ip := c.RemoteIP()
	c.JSON(http.StatusOK, gin.H{
		"ip": ip,
	})
}
