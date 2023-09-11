package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()
	r.GET("/ip", rspIP)
	r.GET("/rand", rspRand)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
