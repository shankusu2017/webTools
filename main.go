package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shankusu2017/url"
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
	r.POST("/v1/monitor/post", monitorPost)
	r.GET("/v1/monitor", monitorGet)

	r.GET(fmt.Sprintf("%s", url.URL_REPEATER_SERVER))
	r.POST(fmt.Sprintf("%s", url.URL_EVENT_POST))

	r.Run(":80") // 监听并在 0.0.0.0:80 上启动服务
}
