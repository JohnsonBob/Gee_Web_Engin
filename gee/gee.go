package gee

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) GET(relativePath string, handlerFunc HandlerFunc) {
	engine.router.addRouter("GET", relativePath, handlerFunc)
}

func (engine *Engine) POST(relativePath string, handlerFunc HandlerFunc) {
	engine.router.addRouter("POST", relativePath, handlerFunc)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	engine.router.handle(NewContext(writer, request))
}
