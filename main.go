package main

import (
	"github.com/gin-gonic/gin"
	"okki.hu/garric/ppnext/config"
	"okki.hu/garric/ppnext/consts"
	"okki.hu/garric/ppnext/controller"
	"okki.hu/garric/ppnext/model"
)

func main() {

	room := model.NewRoom("demo")
	room.RegisterVote(model.NewVote("user1", 5))
	room.RegisterVote(model.NewVote("user2", 3))
	room.RegisterVote(model.NewVote("user3", 2))
	room.RegisterVote(model.NewVote("user4", 1))

	config.Repository.Save(room)

	r := gin.Default()

	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.LoadHTMLGlob("templates/*")

	r.Use(controller.Auth())

	// public routes
	r.GET("/", controller.ShowLogin)
	r.GET("/login", controller.ShowLogin)
	r.POST("/login", controller.HandleLogin)

	// protected routes
	prot := r.Group("/rooms", controller.Prot())
	prot.GET("/:room", controller.ShowRoom)
	prot.GET("/:room/json", controller.GetRoom)

	active := prot.Group("/", controller.Active())
	active.POST("/:room/", controller.AcceptVote)
	active.POST("/:room/show", nil) // show votes
	active.POST("/:room/next", nil) // next story

	r.Run(consts.Addr)
}
