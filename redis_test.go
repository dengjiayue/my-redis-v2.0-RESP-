package main

import (
	"my_redis/src"
	"testing"
)

func TestRedisServer(t *testing.T) {
	server := src.NewServer(100)
	// 优雅的关闭
	defer server.Close()
	go server.Run(":8081")
	server.Deal()

}
