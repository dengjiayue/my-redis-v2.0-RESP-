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
