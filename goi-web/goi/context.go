package goi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type M map[string]interface{}

type Context struct {
	//origin objects
	Writer http.ResponseWriter
	Req    *http.Request

	//request info
	Path   string
	Method string
	Params map[string]string

	//response info
	StatusCode int

	//middleware
	handlers []HandlerFunc
	index    int

	//engine pointer
	engine *Engine
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, M{"message": err})
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

//获取post表单中的值
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//获取url中的参数值
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//设置响应状态码
func (c *Context) SetStatus(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//设置响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

//设置纯文本格式的响应
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

//设置Json格式的响应
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//设置二进制的响应
func (c *Context) Data(code int, data []byte) {
	c.SetStatus(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(http.StatusInternalServerError, err.Error())
	}
}
