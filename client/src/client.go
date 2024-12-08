package src

import (
	"encoding/json"
	"fmt"
	"io"
	"my_redis/public"
)

// 客户端

type ReqData struct {
	UniqueTag uint32
	ReqData   []byte
}

func (c *Client) request(data []byte) uint32 {
	tag := c.Tag.GetUniqueTag()
	c.M.Store(tag, make(chan []byte, 1))

	// 发送数据
	go func(d []byte) {
		req := Encode(tag, d)
		_, err := c.Conn.Write(req)
		if err != nil {
			fmt.Println("Error writing to connection:", err)
		}
	}(data)

	return tag

}

func (c *Client) Get(key string) string {
	// 构造请求数据
	req := new(public.GetReq)
	req.Key = key
	// 序列化
	data, err := json.Marshal(req)
	if err != nil {
		return err.Error()
	}
	// 封装请求方法数据
	reqdata := make([]byte, 1+len(data))
	copy(reqdata[1:], data)
	reqdata[0] = public.GET
	// 发送请求
	tag := c.request(reqdata)
	// fmt.Println("req success")
	// fmt.Println("wait resp")
	// 等待结果
	ch, ok := c.M.Load(tag)
	if !ok {
		return "请求失败"
	}
	resp := <-ch.(chan []byte)
	// 解析结果
	return string(resp)
}

func (c *Client) Set(key, value string) string {
	// 构造请求数据
	req := new(public.SetReq)
	req.Key = key
	req.Value = value
	// 序列化
	data, err := json.Marshal(req)
	if err != nil {
		return err.Error()
	}
	// 封装请求方法数据
	reqdata := make([]byte, 1+len(data))
	reqdata[0] = public.SET
	copy(reqdata[1:], data)
	// fmt.Println("reqdata", reqdata)
	// 发送请求
	tag := c.request(reqdata)
	// fmt.Println("req success")
	// 等待结果
	// fmt.Println("wait resp")
	ch, ok := c.M.Load(tag)
	if !ok {
		return "请求失败"
	}
	resp := <-ch.(chan []byte)
	// 解析结果
	return string(resp)
}

// 处理conn的返回数据并写入对应的map[tag]chan中
func (c *Client) HandleResp() {
	Reader := io.Reader(c.Conn)
	for {
		//读8个字节,uncode len,tag
		resp := make([]byte, 8)
		_, err := Reader.Read(resp)
		if err != nil {
			return
		}
		// 解析数据
		l, tag := Uncode(resp)
		// 读数据
		resp = make([]byte, l)
		_, err = Reader.Read(resp)
		if err != nil {
			return
		}
		go func(t uint32) {
			// 写入chan
			C, ok := c.M.Load(t)
			if !ok {
				return
			}
			ch := C.(chan []byte)
			defer close(ch)
			ch <- resp
			c.M.Delete(t)
		}(tag)
	}
}

// 编码数据
func Encode(tag uint32, data []byte) []byte {
	resp := make([]byte, 8+len(data))
	l := public.Uint32ToBytes(uint32(len(data)))
	t := public.Uint32ToBytes(tag)
	copy(resp, l)
	copy(resp[4:], t)
	copy(resp[8:], data)
	return resp
}

// 解码数据
func Uncode(data []byte) (uint32, uint32) {
	l := public.BytesToUint32(data[:4])
	t := public.BytesToUint32(data[4:8])
	return l, t
}

// 优雅的关闭
func (c *Client) Close() {
	c.Conn.Close()
	// 关闭所有的chan
	c.M.Range(func(key, value interface{}) bool {
		ch := value.(chan []byte)
		close(ch)
		return true
	})

}
