package gee

import (
	"net/http"
	"strings"
)

// router 路由器结构体
type router struct {
	roots    map[string]*node       // 路由树的根节点，根据请求方法分类存储
	handlers map[string]HandlerFunc // 存储路由与处理函数的映射关系
}

// newRouter 返回一个新的路由器实例
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),       // 初始化路由树
		handlers: make(map[string]HandlerFunc), // 初始化处理函数映射
	}
}

// parsePattern 解析路由模式，返回路由模式的各个部分
// 只允许存在一个 * 通配符
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// addRoute 添加路由和处理函数的映射关系
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	// 解析路由模式，得到路由的各个部分
	parts := parsePattern(pattern)

	// 构建路由的唯一键
	key := method + "-" + pattern

	// 检查该请求方法的根节点是否存在，如果不存在则创建
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	// 将路由插入到路由树中
	r.roots[method].insert(pattern, parts, 0)

	// 将路由与处理函数的映射关系存储起来
	r.handlers[key] = handler
}

// getRoute 根据请求方法和路径查找路由
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	// 解析请求路径，得到路径的各个部分
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
				params[part[1:]] = searchParts[index]
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

// getRoutes 获取指定请求方法下的所有路由节点
func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

// handle 处理请求
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
