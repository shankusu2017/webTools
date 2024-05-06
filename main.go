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

	r.LoadHTMLGlob("templates/*")
	r.GET("/", rspHome)

	r.GET("/img", rspImg)
	r.POST("/monitor/post", monitorPost)
	r.GET("/monitor", monitorGet)

	r.Run(":80") // 监听并在 0.0.0.0:80 上启动服务
}
