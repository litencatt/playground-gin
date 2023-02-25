package main

import (
	"playground-gin/handler"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("sample_session", store))

	r.LoadHTMLGlob("templates/*.go.tmpl")
	r.GET("/", handler.RootHandler)
	r.GET("/foo", handler.FooHandler)

	r.Run(":8090")
}
