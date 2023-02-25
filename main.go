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
	sessionNames := []string{"session_foo", "session_bar"}
	r.Use(sessions.SessionsMany(sessionNames, store))

	r.LoadHTMLGlob("templates/*.go.tmpl")
	r.GET("/", handler.RootHandler)
	r.GET("/foo", handler.FooHandler)
	r.GET("/bar", handler.BarHandler)

	r.GET("/login", handler.GetLoginHandler)
	r.POST("/login", handler.PostLoginHandler)
	r.GET("/logout", handler.LogoutHandler)

	r.Run(":8090")
}
