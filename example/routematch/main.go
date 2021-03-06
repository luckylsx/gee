package main

import (
	"net/http"

	"github.com/luckylsx/gee"
)

func main() {
	r := gee.New()
	// r.GET("/", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	// })

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/hello/:name/go", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s, go\n", c.Param("name"), c.Path)
	})

	r.GET("/hello/:name/rust", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s, rust\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
