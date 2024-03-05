package main

import (
	"net/http"

	"gee"
)

func main() {
	// 创建一个 Gee 实例
	r := gee.New()

	// 处理根路径请求，返回一个简单的 HTML 页面
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	// 处理 /hello 路径的 GET 请求
	r.GET("/hello", func(c *gee.Context) {
		// 期望请求路径为 /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	// 处理 /login 路径的 POST 请求
	r.POST("/login", func(c *gee.Context) {
		// 返回 JSON 格式的响应，包含提交的用户名和密码
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	// 运行 HTTP 服务器，监听在 9999 端口上
	r.Run(":8080")
}
