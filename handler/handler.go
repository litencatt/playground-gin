package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("session_value", 1)
	session.Save()
	c.HTML(http.StatusOK, "root.go.tmpl", gin.H{
		"title": "root",
	})
}

func FooHandler(c *gin.Context) {
	session := sessions.Default(c)
	val := session.Get("session_value")
	var v int
	if val != nil {
		v = val.(int) + 1
	}
	session.Set("session_value", v)
	session.Save()
	c.HTML(http.StatusOK, "foo.go.tmpl", gin.H{
		"title": "foo",
		"value": val,
	})
}

func BarHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	val := session.Get("session_value")
	c.HTML(http.StatusOK, "foo.go.tmpl", gin.H{
		"title": "foo",
		"value": val,
	})
}
