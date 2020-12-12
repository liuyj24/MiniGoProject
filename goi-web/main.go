package main

import (
	"goi"
	"log"
	"net/http"
)

func main() {
	g := goi.New()
	g.GET("/", func(c *goi.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Goi</h1>")
	})

	g.GET("/hello", func(c *goi.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	g.POST("/login", func(c *goi.Context) {
		c.JSON(http.StatusOK, goi.M{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	log.Fatal(g.Run(":9999"))
}
