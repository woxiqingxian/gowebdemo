package main

import (
	"fmt"
	"gowebdemo/core"
)

func main() {
	core.InitServer()

	core.RunHttpServer(func() {
		// server shutdown callbak function
		// 服务器接收到关闭信号，即将关闭，提供关闭前的回调
		fmt.Println("RunHttpServer callback start")
		// time.Sleep(3 * time.Second)
		fmt.Println("RunHttpServer callback end")
	})
}
