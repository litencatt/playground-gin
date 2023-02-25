package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	sessionFoo := sessions.DefaultMany(c, "session_foo")
	sessionFoo.Set("value", 1)
	sessionFoo.Save()

	sessionBar := sessions.DefaultMany(c, "session_bar")
	sessionBar.Set("value", 1)
	sessionBar.Save()

	c.HTML(http.StatusOK, "root.go.tmpl", gin.H{
		"title": "root",
	})
}

func FooHandler(c *gin.Context) {
	sessionFoo := sessions.DefaultMany(c, "session_foo")
	val := sessionFoo.Get("value")
	var v int
	if val != nil {
		v = val.(int) + 1
	}
	sessionFoo.Set("value", v)
	sessionFoo.Save()
	c.HTML(http.StatusOK, "foo.go.tmpl", gin.H{
		"title": "foo",
		"value": val,
	})
}

func BarHandler(c *gin.Context) {
	sessionBar := sessions.DefaultMany(c, "session_bar")
	val := sessionBar.Get("value")
	c.HTML(http.StatusOK, "bar.go.tmpl", gin.H{
		"title": "bar",
		"value": val,
	})
}
