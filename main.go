package main

import "my_redis/src"

func main() {
	server := src.NewServer(100)
	// 优雅的关闭
	defer server.Close()
	go server.Run(":8080")
	server.Deal()
}
