package main

import (
	"playground-gin/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.go.tmpl")
	r.GET("/", handler.RootHandler)

	r.Run(":8090")
}
