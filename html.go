package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shankusu2017/utils"
	"log"
	"net/http"
	"os"
	"time"
)

type statement struct {
	Text   string `json:"Text"`
	Weight int    `json:"Weight"`
}

type stateManagerT struct {
	text    []statement
	refresh time.Time
}

var (
	stateMgr stateManagerT
)

func refreshText() error {
	path := "cfg/text.json"
	f, err := os.Stat(path)
	if err != nil {
		log.Fatalf("FATAL 3483c387 %s", err.Error())
	}

	if f.ModTime() == stateMgr.refresh {
		return nil
	}
	stateMgr.refresh = f.ModTime()

	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("FATAL b6bd480a %s", err.Error())
	}

	err = json.Unmarshal(b, &stateMgr.text)
	if err != nil {
		log.Printf("WARN f235bfc2 %s", err.Error())
		return err
	}

	return nil
}

func popText() string {
	var w []int
	for _, item := range stateMgr.text {
		w = append(w, item.Weight)
	}
	i, _ := utils.RandWeigh(w)
	return stateMgr.text[i].Text
}

func rspHome(c *gin.Context) {
	refreshText()

	title := popText()

	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": title,
	})
}
