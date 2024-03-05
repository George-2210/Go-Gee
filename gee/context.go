package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 是一个类型别名，表示一个键值对的 map，用于构建 JSON 响应数据
type H map[string]interface{}

// Context 结构体包含了处理 HTTP 请求和构建 HTTP 响应所需的各种属性和方法
type Context struct {
	// 原始对象
	Writer http.ResponseWriter // ResponseWriter 接口，用于写入响应数据
	Req    *http.Request       // Request 结构体，包含了客户端发起的请求信息
	// 请求信息
	Path   string            // 请求路径
	Method string            // 请求方法
	Params map[string]string // 存储 URL 路径参数的 map
	// 响应信息
	StatusCode int // HTTP 响应状态码
}

// newContext 创建并返回一个新的 Context 实例
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// Param 获取 URL 路径参数的值
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// PostForm 获取表单中指定键的值
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 获取 URL 查询参数的值
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置 HTTP 响应状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置 HTTP 响应头部信息
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 返回文本类型的响应
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 返回 JSON 类型的响应
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 返回二进制数据类型的响应
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 返回 HTML 类型的响应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
