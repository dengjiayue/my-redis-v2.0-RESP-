package src

import (
	"fmt"
	"net"
)

// 发送成功响应
func SendSuccessResponse(conn net.Conn, data string) error {
	fmt.Println("data:", data)
	// 构造响应字符串
	response := fmt.Sprintf("+%s\r\n", data)

	// 写入到连接中
	_, err := conn.Write([]byte(response))
	if err != nil {
		return fmt.Errorf("failed to send success response: %v", err)
	}

	return nil
}

// 发送错误响应
func SendErrorResponse(conn net.Conn, errMsg string) error {
	// 构造错误响应字符串
	response := fmt.Sprintf("-Err %s\r\n", errMsg)

	// 写入到连接中
	_, err := conn.Write([]byte(response))
	if err != nil {
		return fmt.Errorf("failed to send error response: %v", err)
	}

	return nil
}

// 发送整数响应
func SendIntegerResponse(conn net.Conn, num int) error {
	// 构造整数响应字符串
	response := fmt.Sprintf(":%d\r\n", num)
	// 写入到连接中
	_, err := conn.Write([]byte(response))
	if err != nil {
		return fmt.Errorf("failed to send integer response: %v", err)
	}
	return nil
}

// 发送Bulk字符串响应(对象/多个字符串响应)
func SendBulkStringResponse(conn net.Conn, data []string) error {
	// 构造Bulk字符串响应字符串
	response := fmt.Sprintf("*%d\r\n", len(data))
	for _, s := range data {
		response += fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
	}
	// 写入到连接中
	_, err := conn.Write([]byte(response))
	if err != nil {
		return fmt.Errorf("failed to send bulk string response: %v", err)
	}
	return nil
}

// 单个字符串响应
func SendSingleStringResponse(conn net.Conn, data string) error {
	// 构造Bulk字符串响应字符串
	response := fmt.Sprintf("$%d\r\n%s\r\n", len(data), data)
	// 写入到连接中
	_, err := conn.Write([]byte(response))
	if err != nil {
		return fmt.Errorf("failed to send bulk string response: %v", err)
	}
	return nil
}
