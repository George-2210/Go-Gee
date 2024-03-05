package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 定义了 gee 框架使用的请求处理函数
type HandlerFunc func(*Context)

// Engine 实现了 ServeHTTP 接口
type (
	RouterGroup struct {
		prefix      string        // 前缀
		middlewares []HandlerFunc // 支持中间件
		parent      *RouterGroup  // 支持嵌套
		engine      *Engine       // 所有组共享一个 Engine 实例
	}

	Engine struct {
		*RouterGroup                // 指向顶层的 RouterGroup
		router       *router        // 路由器
		groups       []*RouterGroup // 存储所有的路由组
	}
)

// New 是 gee.Engine 的构造函数
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 定义了创建新 RouterGroup 的方法
// 所有的路由组共享同一个 Engine 实例
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

// addRoute 添加路由规则
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET 定义了添加 GET 请求的方法
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 定义了添加 POST 请求的方法
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run 定义了启动 HTTP 服务器的方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 实现了 http.Handler 接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
