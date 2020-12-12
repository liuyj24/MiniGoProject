package main

import (
	"goi"
	"log"
	"net/http"
)

func main() {
	g := goi.New()
	g.GET("/index", func(c *goi.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := g.Group("/v1")
	{
		v1.GET("/", func(c *goi.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Goi</h1>")
		})
		v1.GET("/hello", func(c *goi.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := g.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *goi.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Params["name"], c.Path)
		})
		v2.POST("/login", func(c *goi.Context) {
			c.JSON(http.StatusOK, goi.M{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	//g.GET("/", func(c *goi.Context) {
	//	c.HTML(http.StatusOK, "<h1>Hello Goi</h1>")
	//})
	//
	//g.GET("/hello", func(c *goi.Context) {
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//})
	//g.POST("/login", func(c *goi.Context) {
	//	c.JSON(http.StatusOK, goi.M{
	//		"username": c.PostForm("username"),
	//		"password": c.PostForm("password"),
	//	})
	//})
	g.GET("/assets/*filepath", func(c *goi.Context) {
		c.JSON(http.StatusOK, goi.M{"filepath": c.Param("filepath")})
	})
	log.Fatal(g.Run(":9999"))
}
