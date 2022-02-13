package main

import (
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	engine.GET("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})
	err := engine.Run(":9000")
	if err != nil {
		panic(err)
	}
}
