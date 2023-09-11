package main

import "github.com/gin-gonic/gin"

func rspIP(c *gin.Context) {
	ip := c.RemoteIP()
	c.JSON(200, gin.H{
		"ip": ip,
	})
}
