package src

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestClient_Set(t *testing.T) {
	client := NewClient("localhost:8080")
	// fmt.Println("new client")
	go client.HandleResp()
	// fmt.Println("handle resp")
	defer client.Close()
	// fmt.Printf("%v\n", client.Set("name", "zhangsan"))
	// fmt.Println(client.Get("name"))

}

// chan	无缓冲测试
func TestChan(t *testing.T) {
	ch := make(chan int)
	go func() {
		num := <-ch
		fmt.Println(num)
	}()
	ch <- 100
	time.Sleep(time.Second)
}

// 压测my redis
func BenchmarkMyRedisWrite(b *testing.B) {
	c := NewClient("localhost:8080")
	go c.HandleResp()
	defer c.Close()
	//开始计时
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Set("name", "zhangsan")
	}
	// BenchmarkMyRedis-8   	   28090	     40598 ns/op	     642 B/op	      14 allocs/op
}

// 压测my redis
func BenchmarkMyRedisRead(b *testing.B) {
	c := NewClient("localhost:8080")
	go c.HandleResp()
	defer c.Close()
	//开始计时
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Get("name")
	}
	// BenchmarkMyRedisRead-8             27771             44423 ns/op             588
}

// 并发压测(写)
func BenchmarkMyRedisConcurrencyWrite(b *testing.B) {
	c := NewClient("localhost:8080")
	go c.HandleResp()
	defer c.Close()
	//开始计时
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Set("name", "zhangsan")
		}
	})
	// BenchmarkMyRedisConcurrencyWrite-8   	   90667	     12439 ns/op	     612 B/op	      14 allocs/op
}

// 并发压测(读)
func BenchmarkMyRedisConcurrencyRead(b *testing.B) {
	c := NewClient("localhost:8080")
	go c.HandleResp()
	defer c.Close()
	//开始计时
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get("name")
		}
	})
	// BenchmarkMyRedisConcurrencyRead-8   	   89955	     12198 ns/op	     512 B/op	      15 allocs/op
}

// 单元测试
func TestMyRedis2Read(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	// 写入数据
	conn.Write([]byte("*2\r\n$3\r\nget\r\n$4\r\nname\r\n"))
	// 读取数据 (+{val}\r\n)
	resp, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

// 压测my redis 2.0 (读)
func BenchmarkMyRedisWithRESP(b *testing.B) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//开始计时
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn.Write([]byte("*2\r\n$3\r\nget\r\n$4\r\nname\r\n"))
		// 读取数据 (+{val}\r\n)
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
	}
	// BenchmarkMyRedisWithRESP-8         38745             30267 ns/op              32 B/op          2 allocs/op
}

// 并发压测my redis 2.0 (读)
func BenchmarkMyRedisWithRESPConcurrencyRead(b *testing.B) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//开始计时
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			conn.Write([]byte("*2\r\n$3\r\nget\r\n$4\r\nname\r\n"))
			// 读取数据 (+{val}\r\n)
			_, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
		}
	})
}

// 单元测试(写)
func TestMyRedis2Write(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	// 写入数据
	conn.Write([]byte("*3\r\n$3\r\nset\r\n$4\r\nname\r\n$8\r\nzhangsan\r\n"))
	// 读取数据 (+{val}\r\n)
	resp, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

// 压测my redis 2.0 (写)
func BenchmarkMyRedisWithRESPWrite(b *testing.B) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//开始计时
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn.Write([]byte("*3\r\n$3\r\nset\r\n$4\r\nname\r\n$8\r\nzhangsan\r\n"))
		// 读取数据 (+{val}\r\n)
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
	}
	//BenchmarkMyRedisWithRESPWrite-8            33255             33665 ns/op             56 B/op           2 allocs/op
}
