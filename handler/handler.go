package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "root.go.tmpl", gin.H{
		"title": "root",
	})
}

func FooHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "foo.go.tmpl", gin.H{
		"title": "foo",
	})
}
