package main

import "github.com/gin-gonic/gin"

func monitorGet(c *gin.Context) {
	nodeLst := GetNodeAll()

	c.JSON(200, nodeLst)
}
