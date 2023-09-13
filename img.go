package main

import "github.com/gin-gonic/gin"

// 参考链接：https://blog.csdn.net/sunriseYJP/article/details/127067521

func rspImg(c *gin.Context) {
	c.File("./img/IMG_20230708_120210.jpg")
}
