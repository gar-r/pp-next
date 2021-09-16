package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRoom(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, name)
}
