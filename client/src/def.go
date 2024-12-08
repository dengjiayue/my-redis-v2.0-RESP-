package src

import (
	"net"
	"sync"
)

// 客户端

type Tag struct {
	UniqueTag uint32
	Lock      sync.Mutex
}

func (t *Tag) GetUniqueTag() uint32 {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.UniqueTag++
	return t.UniqueTag
}

type Client struct {
	Conn net.Conn
	M    sync.Map
	Tag  *Tag
}

func NewClient(port string) *Client {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		panic(err)
	}
	return &Client{
		Conn: conn,
		M:    sync.Map{},
		Tag:  &Tag{UniqueTag: 0, Lock: sync.Mutex{}},
	}
}
