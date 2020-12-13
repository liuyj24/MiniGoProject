package main

import (
	"fmt"
	"goi"
	"html/template"
	"log"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	g := goi.Default()
	g.GET("/", func(c *goi.Context) {
		c.String(http.StatusOK, "hello world!")
	})

	g.GET("/panic", func(c *goi.Context) {
		names := []string{"yijun"}
		c.String(http.StatusOK, names[100])
	})
	g.Run(":9999")
}

func testStaticResource() {
	g := goi.New()
	g.Use(goi.Logger())

	g.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})

	g.LoadHTMLGlob("templates/*")
	g.Static("/assets", "./static")

	stu1 := &student{Name: "yijun", Age: 20}
	stu2 := &student{Name: "Mary", Age: 20}
	g.GET("/", func(c *goi.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	g.GET("/students", func(c *goi.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", goi.M{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	g.GET("/date", func(c *goi.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", goi.M{
			"title": "goi",
			"now":   time.Date(2020, 12, 12, 0, 0, 0, 0, time.UTC),
		})
	})
	log.Fatal(g.Run(":9999"))

}
