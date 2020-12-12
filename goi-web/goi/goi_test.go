package goi

import (
	"log"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := NewRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern fail!")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	node, params := r.getRoute("GET", "/hello/yijun")

	if node == nil {
		t.Fatal("something wrong, the node is nil...")
	}
	if node.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}
	if params["name"] != "yijun" {
		t.Fatal("name should be equal to 'yijun'")
	}
	log.Printf("match path: %s, params['name']: %s\n", node.pattern, params["name"])
}
