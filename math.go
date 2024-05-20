package main

import (
	_ "crypto/rand"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
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

// http://localhost:8081/rand?len=16
func rspRand(c *gin.Context) {
	lenArg := c.Query("len")
	lenHex := 8
	if len(lenArg) > 0 {
		lenInt, err := strconv.Atoi(lenArg)
		if err == nil {
			if lenInt > 0 {
				lenHex = lenInt
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"rand": string("0x") + randomString(lenHex),
	})
}
