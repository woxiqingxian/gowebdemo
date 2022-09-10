package main

import (
	"gowebdemo/core"
)

func main() {
	core.InitComponent()

	// RunHTTPServer 支持注册回调函数
	// core.RunHTTPServer(func() {
	// 	// server shutdown callbak function
	// 	// 服务器接收到关闭信号，即将关闭，提供关闭前的回调
	// 	// fmt.Println("RunHTTPServer callback start")
	// 	// time.Sleep(3 * time.Second)
	// 	// fmt.Println("RunHTTPServer callback end")
	// })
	core.RunHTTPServer()

}
