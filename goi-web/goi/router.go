package goi

import (
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*node
}

func NewRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node)}
}

//将url按‘/’进行划分
func parsePattern(pattern string) []string {
	strs := strings.Split(pattern, "/")
	result := make([]string, 0)
	for _, str := range strs {
		if str != "" {
			result = append(result, str)
			if str[0] == '*' { //我感觉这里有问题，会不会用contains更好
				break
			}
		}
	}
	return result
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {

	//解析url，添加到前缀树中
	parts := parsePattern(pattern)
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)

	//添加handler
	key := method + "-" + pattern
	r.handlers[key] = handler
}

//这个暂时不知道什么意思，目前看来好像是一个比较重要的入口方法
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index] //这里就包含了:
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path)
	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 not found: %s\n", c.Path)
	}
}
