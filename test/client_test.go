package my_redis_test

import (
	"testing"
)

//单元测试

// 写

// string
func TestWriteString(t *testing.T) {
	c := NewClient()
	defer c.Conn.Close()
	c.Set()
}

// string 读
// 源码写,输出响应数据
func TestSendTcp(t *testing.T) {
	// TSet()
	//+OK
	TGet()
	//$3\r\nfoo\r\n
	// THget()
	//$3\r\nbar\r\n
	// THset()
	//:0
	// THgetall()
	//*4\r\n$4\r\nname\r\n$3\r\ntom\r\n$3\r\nage\r\n$2\r\n18\r\n
	// TLpush()
	//:3
	// TRpop()
	//$1\r\na\r\n
	// TSadd()
	//:1
	// TSmembers()
	//*1\r\na\r\n
	// TZadd()
	//:2
	// TZrange()
	//*2\r\n$1\r\nc\r\n$1\r\nd\r\n
}

// hash
func TestWriteHash(t *testing.T) {
	c := NewClient()
	defer c.Conn.Close()
	c.Hset()
}

// list
func TestWriteList(t *testing.T) {
	c := NewClient()
	defer c.Conn.Close()
	c.Lpush()
}

// sadd
func TestWriteSet(t *testing.T) {
	c := NewClient()
	defer c.Conn.Close()
	c.Sadd()
}

// zadd
func TestWriteZset(t *testing.T) {
	c := NewClient()
	defer c.Conn.Close()
	c.Zadd()
}

// 读
// string
func TestReadString(t *testing.T) {
	c := NewClient()
	defer c.Conn.Close()
	c.Get()
}

// hash
func TestReadHash(t *testing.T) {
	c := NewClient()
	defer c.Conn.Close()
	c.Hget()
}

// hash read all
func TestReadHashAll(t *testing.T) {
	c := NewClient()
	defer c.Conn.Close()
	c.Hgetall()
}

// list
func TestReadList(t *testing.T) {
	c := NewClient()
	defer c.Conn.Close()
	c.Rpop()
}

// sadd
func TestReadSet(t *testing.T) {
	c := NewClient()
	defer c.Conn.Close()
	c.Smembers()
}

// zadd
func TestReadZset(t *testing.T) {
	c := NewClient()
	defer c.Conn.Close()
	c.Zrange()
}
