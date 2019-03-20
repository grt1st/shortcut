package main

import (
	"github.com/gin-gonic/gin"
	"github.com/grt1st/short_links/core"
	"github.com/grt1st/short_links/handles"
	"github.com/grt1st/short_links/models"
)

func main() {
	// init
	core.Init()
	models.InitSql()
	server()
}

func server() {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	//router.HTMLRender = handles.CreateRender()

	// static can be handle by nginx
	router.Static("/static", "./static")

	// ping
	router.GET("/ping", handles.Ping)
	// homepage
	router.GET("/", handles.Index)
	// 404
	router.GET("/404", handles.E404)
	// save
	router.POST("/new", handles.NewShortcut)
	// shortcut
	router.GET("/s/:shortcut", handles.GetShortcut)

	router.Run(core.Config.Host)

}
