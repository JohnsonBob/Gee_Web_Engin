package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(writer http.ResponseWriter, request *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRouter(method string, relativePath string, handler HandlerFunc) {
	key := method + "-" + relativePath
	engine.router[key] = handler
}

func (engine *Engine) GET(relativePath string, handlerFunc HandlerFunc) {
	engine.addRouter("GET", relativePath, handlerFunc)
}

func (engine *Engine) POST(relativePath string, handlerFunc HandlerFunc) {
	engine.addRouter("POST", relativePath, handlerFunc)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := request.Method + "-" + request.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(writer, request)
	} else {
		_, _ = fmt.Fprintf(writer, "404 NOT FOUND: %s\n", request.URL)
	}
}
