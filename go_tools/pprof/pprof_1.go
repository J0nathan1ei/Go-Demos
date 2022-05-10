package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(3 * time.Second)
			log.Println("hello")
		}
	}()

	// 本地浏览器登录 127.0.0.1:6060/debug/pprof 查看运行信息
	http.ListenAndServe("127.0.0.1:6060", nil)
}
