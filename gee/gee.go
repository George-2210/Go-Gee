package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 定义了 gee 使用的请求处理函数类型
type HandlerFunc func(*Context)

// Engine 实现了 ServeHTTP 接口
type Engine struct {
	router *router
}

// New 是 gee.Engine 的构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method, pattern, handler)
}

// GET 定义了添加 GET 请求的方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 定义了添加 POST 请求的方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 定义了启动 HTTP 服务器的方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
