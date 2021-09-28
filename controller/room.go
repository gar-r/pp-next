package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"okki.hu/garric/ppnext/config"
	"okki.hu/garric/ppnext/consts"
)

func ShowRoom(c *gin.Context) {
	user := c.MustGet("user")
	name := c.Param("room")
	room, err := config.Repository.Load(name)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	h := gin.H{
		"room":    room,
		"user":    user,
		"options": config.VoteOptions,
		"support": consts.Support,
	}
	c.HTML(http.StatusOK, "room.html", h)
}

func UserList(c *gin.Context) {
	name := c.Param("room")
	room, err := config.Repository.Load(name)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.HTML(http.StatusOK, "user-list.html", room)
}

func AcceptVote(c *gin.Context) {
	var v int
	err := c.ShouldBind(&v)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.Status(http.StatusOK)
}
