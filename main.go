package main

import (
	"github.com/gin-gonic/gin"
)

const addr = ":38080"

func main() {
	r := gin.Default()
	setupHandlers(r)
	r.Run(addr)
}

func setupHandlers(r *gin.Engine) {
	r.GET("/rooms/:name", getRoom)
}
