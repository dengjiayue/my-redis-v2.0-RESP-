package my_redis_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	PORT1 = "16379"
	PORT2 = "8080"
)

type Client struct {
	Conn *redis.Client
}

func NewClient() *Client {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:" + PORT1,
		PoolSize:     1,
		MinIdleConns: 1,
	})
	return &Client{
		Conn: client,
	}
}

// 请求建立redis - tcp连接
func NewTcpConn() net.Conn {
	// 建立连接
	conn, err := net.Dial("tcp", "localhost:"+PORT1)
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return nil
	}
	fmt.Println("tcp连接成功")
	return conn
}

// 读取conn数据
func ReadData(conn net.Conn) ([]byte, error) {
	//构建1024字节的缓冲区
	buf := make([]byte, 1024)
	//读取数据
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return nil, err
	}
	//清除空格
	buf = buf[:n]
	return buf, nil

	// reader := io.Reader(conn)
	// data, err := ioutil.ReadAll(reader)
	// if err != nil {
	// 	fmt.Println("Error reading from connection:", err)
	// 	return nil, err
	// }
	// return data, nil
}

// 字符串写
func (c *Client) Set() {
	ctx := context.Background()
	data, err := c.Conn.Set(ctx, "name", "tom", 0).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func (c *Client) Set2() {
	ctx := context.Background()
	data, err := c.Conn.Set(ctx, "test", "tom", 1*time.Minute).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
}

func (c *Client) Get2() {
	ctx := context.Background()
	data, err := c.Conn.Get(ctx, "test").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
}

func (c *Client) Hset2() {
	ctx := context.Background()
	data, err := c.Conn.HSet(ctx, "test", "name", "tom").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
}

// 字符串读
func (c *Client) Hgetall2() {
	ctx := context.Background()
	data, err := c.Conn.HGetAll(ctx, "test").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
}

// 使用tcp发送请求
func TSet() {
	conn := NewTcpConn()
	defer conn.Close()
	conn.Write([]byte("*3\r\n$3\r\nset\r\n$4\r\nname\r\n$3\r\ntom\r\n"))
	// 读取响应数据
	data, err := ReadData(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

// 字符串读
func (c *Client) Get() {
	ctx := context.Background()
	data, err := c.Conn.Get(ctx, "name").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

// 字符串读
func TGet() {
	conn := NewTcpConn()
	defer conn.Close()
	conn.Write([]byte("*2\r\n$3\r\nget\r\n$4\r\nname\r\n"))
	// 读取响应数据
	data, err := ReadData(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

// 哈希写
func (c *Client) Hset() {
	ctx := context.Background()
	data, err := c.Conn.HSet(ctx, "user", "name", "tom", "age", 18).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

// 哈希写
func THset() {
	conn := NewTcpConn()
	defer conn.Close()
	conn.Write([]byte("*6\r\n$4\r\nhset\r\n$4\r\nuser\r\n$4\r\nname\r\n$3\r\ntom\r\n$3\r\nage\r\n$2\r\n19\r\n"))
	// 读取响应数据
	data, err := ReadData(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

// 哈希读
func (c *Client) Hget() {
	ctx := context.Background()
	data, err := c.Conn.HGet(ctx, "user", "name").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

// 哈希读
func THget() {
	conn := NewTcpConn()
	defer conn.Close()
	conn.Write([]byte("*3\r\n$4\r\nhget\r\n$4\r\nuser\r\n$4\r\nname\r\n"))
	// 读取响应数据
	data, err := ReadData(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

// 哈希整个读
func (c *Client) Hgetall() {
	ctx := context.Background()
	data, err := c.Conn.HGetAll(ctx, "user").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

// 哈希整个读
func THgetall() {
	conn := NewTcpConn()
	defer conn.Close()
	conn.Write([]byte("*2\r\n$7\r\nhgetall\r\n$4\r\nuser\r\n"))
	// 读取响应数据
	data, err := ReadData(conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}

// 列表写
func (c *Client) Lpush() {
	ctx := context.Background()
	data, err := c.Conn.LPush(ctx, "list", "a", "b", "c").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

// 列表读
func TLpush() {
	conn := NewTcpConn()
	defer conn.Close()
	conn.Write([]byte("*3\r\n$5\r\nlpush\r\n$4\r\nlist\r\n$1\r\na\r\n"))
	// 读取响应数据
	data, err := ReadData(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

// 列表读
func (c *Client) Rpop() {
	ctx := context.Background()
	data, err := c.Conn.RPop(ctx, "list").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

// 列表读
func TRpop() {
	conn := NewTcpConn()
	defer conn.Close()
	conn.Write([]byte("*2\r\n$4\r\nrpop\r\n$4\r\nlist\r\n"))
	// 读取响应数据
	data, err := ReadData(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

// 集合写
func (c *Client) Sadd() {
	ctx := context.Background()
	data, err := c.Conn.SAdd(ctx, "set", "a", "b", "c").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

// 集合写
func TSadd() {
	conn := NewTcpConn()
	defer conn.Close()
	conn.Write([]byte("*3\r\n$4\r\nsadd\r\n$4\r\nsets\r\n$1\r\na\r\n"))
	// 读取响应数据
	data, err := ReadData(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

// 集合读
func (c *Client) Smembers() {
	ctx := context.Background()
	data, err := c.Conn.SMembers(ctx, "set").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

// 集合读
func TSmembers() {
	conn := NewTcpConn()
	defer conn.Close()
	conn.Write([]byte("*2\r\n$8\r\nsmembers\r\n$4\r\nsets\r\n"))
	// 读取响应数据
	data, err := ReadData(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

// 有序集合写
func (c *Client) Zadd() {
	ctx := context.Background()
	data, err := c.Conn.ZAdd(ctx, "zset", &redis.Z{
		Score:  1,
		Member: "a",
	}, &redis.Z{
		Score:  2,
		Member: "b",
	}).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

// 有序集合写
func TZadd() {
	conn := NewTcpConn()
	defer conn.Close()
	conn.Write([]byte("*6\r\n$4\r\nzadd\r\n$5\r\nzset2\r\n$1\r\n1\r\n$1\r\nc\r\n$1\r\n2\r\n$1\r\nd\r\n"))
	// 读取响应数据
	data, err := ReadData(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

// 有序集合读
func (c *Client) Zrange() {
	ctx := context.Background()
	data, err := c.Conn.ZRange(ctx, "zset", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

// 有序集合读
func TZrange() {
	conn := NewTcpConn()
	defer conn.Close()
	conn.Write([]byte("*4\r\n$6\r\nzrange\r\n$5\r\nzset2\r\n$1\r\n0\r\n$2\r\n-1\r\n"))
	// 读取响应数据
	data, err := ReadData(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
