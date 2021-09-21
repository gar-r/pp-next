package main

import (
	"github.com/gin-gonic/gin"
	"okki.hu/garric/ppnext/config"
	"okki.hu/garric/ppnext/controller"
	"okki.hu/garric/ppnext/mw"
)

func main() {
	r := gin.Default()

	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.LoadHTMLGlob("templates/*")

	r.Use(mw.Auth())

	// public routes
	r.GET("/", controller.ShowLogin)
	r.GET("/login", controller.ShowLogin)
	r.POST("/login", controller.HandleLogin)

	// protected routes
	prot := r.Group("/rooms", mw.Prot())
	prot.GET("/:room", controller.ShowRoom)

	active := prot.Group("/", mw.Active())
	active.POST("/:room/", controller.AcceptVote)
	active.POST("/:room/show", nil) // show votes
	active.POST("/:room/next", nil) // next story

	r.Run(config.Addr)
}
