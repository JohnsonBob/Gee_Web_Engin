package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Path           string
	Method         string
	StatusCode     int
}

func NewContext(write http.ResponseWriter, res *http.Request) *Context {
	return &Context{write, res, res.URL.Path, res.Method, http.StatusOK}
}

func (context *Context) PostForm(key string) string {
	return context.Request.FormValue(key)
}

func (context *Context) Query(key string) string {
	return context.Request.URL.Query().Get(key)
}

func (context *Context) Status(code int) {
	context.StatusCode = code
	context.ResponseWriter.WriteHeader(code)
}

func (context *Context) SetHeader(key string, value string) {
	context.ResponseWriter.Header().Set(key, value)
}

func (context *Context) String(code int, format string, values ...interface{}) {
	context.SetHeader("Content-Type", "text/plain")
	context.Status(code)
	_, err := context.ResponseWriter.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		http.Error(context.ResponseWriter, err.Error(), 500)
	}
}

func (context *Context) JSON(code int, obj interface{}) {
	context.Status(code)
	encoder := json.NewEncoder(context.ResponseWriter)
	err := encoder.Encode(obj)
	if err != nil {
		http.Error(context.ResponseWriter, err.Error(), 500)
	}
}

func (context *Context) Data(code int, data []byte) {
	context.Status(code)
	_, err := context.ResponseWriter.Write(data)
	if err != nil {
		http.Error(context.ResponseWriter, err.Error(), 500)
	}
}

func (context *Context) HTML(code int, html string) {
	context.SetHeader("Content-Type", "text/html")
	context.Status(code)
	_, err := context.ResponseWriter.Write([]byte(html))
	if err != nil {
		http.Error(context.ResponseWriter, err.Error(), 500)
	}
}
