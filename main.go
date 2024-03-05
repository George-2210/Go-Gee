package main

import (
	"fmt"
	"log"
	"net/http"
)

// 标准库启动Web服务

func main() {
	// 注册路由处理函数
	http.HandleFunc("/", indexHandler)      // 处理根路径
	http.HandleFunc("/hello", helloHandler) // 处理 /hello 路径
	// 启动HTTP服务器，监听8080端口
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 处理根路径的函数
func indexHandler(w http.ResponseWriter, req *http.Request) {
	// 输出请求路径
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// 处理 /hello 路径的函数
func helloHandler(w http.ResponseWriter, req *http.Request) {
	// 输出请求头信息
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

// 请求头信息
// Header["Connection"] = ["keep-alive"]
// Header["Sec-Ch-Ua-Mobile"] = ["?0"]
// Header["Sec-Fetch-Site"] = ["none"]
// Header["Sec-Fetch-Mode"] = ["navigate"]
// Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"]
// Header["Sec-Fetch-Dest"] = ["document"]
// Header["Accept-Language"] = ["zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6"]
// Header["Sec-Ch-Ua"] = ["\"Microsoft Edge\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\""]
// Header["Sec-Ch-Ua-Platform"] = ["\"Windows\""]
// Header["User-Agent"] = ["Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0"]
// Header["Sec-Fetch-User"] = ["?1"]
// Header["Upgrade-Insecure-Requests"] = ["1"]
// Header["Accept-Encoding"] = ["gzip, deflate, br"]
