package src

import (
	"fmt"
)

// 生成成功响应字符串
func SuccessResponse() string {
	return "+ok\r\n"
}

// 错误响应
func ErrorResponse(errMsg string) string {
	return fmt.Sprintf("-%s\r\n", errMsg)
}

// 整数响应
func IntegerResponse(num int) string {
	return fmt.Sprintf(":%d\r\n", num)
}

// 发送Bulk字符串响应(对象/多个字符串响应)
func BulkStringResponse(data []string) string {
	// 构造Bulk字符串响应字符串
	response := fmt.Sprintf("*%d\r\n", len(data))
	for _, s := range data {
		response += fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
	}
	return response
}

// 单个字符串响应
func SingleStringResponse(data string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(data), data)
}
