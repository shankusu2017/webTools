package main

import (
	_ "crypto/rand"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
)

const charset = "0123456789abcdef"

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func rspRand(c *gin.Context) {
	c.JSON(200, gin.H{
		"rand": randomString(8),
	})
}
