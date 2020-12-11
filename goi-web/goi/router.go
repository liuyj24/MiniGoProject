package goi

import "net/http"

type router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handlerFunc, ok := r.handlers[key]; ok {
		handlerFunc(c)
	} else {
		c.String(http.StatusNotFound, "404 not found: %s\n", c.Path)
	}
}
