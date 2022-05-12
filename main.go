package main

import (
	"net/http"
	"nin/nin"
)

func main() {
	r := nin.New()
	r.GET("/", func(c *nin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Nin</h1>")
	})

	r.GET("/hello", func(c *nin.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *nin.Context) {
		c.JSON(http.StatusOK, nin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
