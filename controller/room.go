package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"okki.hu/garric/ppnext/config"
)

func ShowRoom(c *gin.Context) {
	name := c.Param("room")
	room, err := config.Repository.Load(name)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	h := gin.H{
		"room":    room,
		"options": config.VoteOptions,
	}
	c.HTML(http.StatusOK, "room.html", h)
}

func AcceptVote(c *gin.Context) {
	var v int
	err := c.ShouldBind(&v)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.Status(http.StatusOK)
}
