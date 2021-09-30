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
	room.RegisterVote(model.NewVote("user5", consts.Nothing))
	room.RegisterVote(model.NewVote("user6", consts.Coffee))
	room.RegisterVote(model.NewVote("user7", consts.Large))
	room.RegisterVote(model.NewVote("user8", consts.Question))
	room.Revealed = false
	config.Repository.Save(room)

	r := gin.Default()

	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.StaticFile("/scripts.js", "./assets/scripts.js")
	r.LoadHTMLGlob("templates/*")

	r.Use(controller.Auth())

	// public routes
	r.GET("/", controller.ShowLogin)
	r.GET("/login", controller.ShowLogin)
	r.POST("/login", controller.HandleLogin)

	// protected routes
	prot := r.Group("/rooms", controller.Prot())
	prot.GET("/:room", controller.DisplayRoom)
	prot.GET("/:room/userlist", controller.UserList)
	prot.GET("/:room/results", controller.Results)
	prot.GET("/:room/events", controller.GetEvents)

	active := prot.Group("/", controller.Active())
	active.POST("/:room/vote", controller.AcceptVote)
	active.POST("/:room/reveal", controller.Reveal)
	active.POST("/:room/reset", controller.ResetRoom)

	r.Run(consts.Addr)
}
