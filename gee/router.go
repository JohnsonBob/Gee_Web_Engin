package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{make(map[string]HandlerFunc)}
}

func (r *router) addRouter(method string, relativePath string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, relativePath)
	key := method + "-" + relativePath
	r.handlers[key] = handler
}

func (r *router) handle(ctx *Context) {
	key := ctx.Method + "-" + ctx.Path
	if handler, ok := r.handlers[key]; ok {
		handler(ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}
