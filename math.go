package main

import (
	_ "crypto/rand"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
)

const signedCharst = "01234567"
const charset = "0123456789abcdef"

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)

	// 限制第一个字符的范围，使产生的十六进制字符在 int 的范围内
	// 可以直接在前面添加 "-" 而不用担心负数溢出的问题(eg: -f3876898)
	sb.WriteByte(charset[rand.Intn(len(signedCharst))])
	for i := 1; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func rspRand(c *gin.Context) {
	c.JSON(200, gin.H{
		"rand": randomString(8),
	})
}
