package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type statement struct {
	text   string `json:"text"`
	weight int    `json:"weight"`
}

func rspHome(c *gin.Context) {
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "不要因为走的太远，而忘了为什么出发。",
	})
}
