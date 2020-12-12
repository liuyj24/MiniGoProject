package goi

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type (
	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}

	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}
)

func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, postfix string, handler HandlerFunc) {
	pattern := group.prefix + postfix
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

//处理HTTP请求：根据request和response创建上下文，交给router处理
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := newContext(w, req)
	engine.router.handle(context)
}

//监听端口，启动框架
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
