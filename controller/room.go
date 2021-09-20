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
		"room": room,
	}
	c.HTML(http.StatusOK, "room.html", h)
}
