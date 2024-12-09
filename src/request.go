package src

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

// 读取每一行数据: 先读一行数据的长度($n), 根据长度再读取数据一行数据
// 解析每个数据部分的长度和数据体
func ParseData(reader *bufio.Reader) (string, error) {
	// 读取当前行数据，即以 $X 开头的长度
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// 去掉行尾换行符
	line = strings.TrimSpace(line)

	// 确保这一行以 "$" 开头
	if !strings.HasPrefix(line, "$") {
		return "", fmt.Errorf("invalid data length line: %s", line)
	}

	//读取下一行数据
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read data body: %v", err)
	}
	// 去掉行尾的换行符
	data = strings.TrimSpace(data)

	return data, nil
}

// 解析第一行 "*N" 格式，返回 N 和可能的错误
func ParseArrayLength(reader *bufio.Reader) (int, error) {
	// 读取第一行
	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	// 去掉行尾的换行符
	line = strings.TrimSpace(line)

	// 确保这一行以 "*" 开头
	if !strings.HasPrefix(line, "*") {
		return 0, fmt.Errorf("invalid array length line: %s", line)
	}

	// 提取 N 值，去掉前面的 '*'，然后解析为整数
	nStr := line[1:]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		return 0, fmt.Errorf("invalid array length: %v", err)
	}

	return n, nil
}

// 获取一个请求chan
func (s *Server) GetChan() chan []string {
	// 获取队列的一个请求chan(阻塞等待)
	rsp := <-s.DisengagedQueue
	return rsp
}

// 向通道写入数据, 并将通道放入处理队列
func (s *Server) Request(rsp chan []string, data []string) {
	// 放入处理队列
	s.DealQueue <- rsp
	// 向通道写入数据
	rsp <- data

}

// 启动处理服务,监听10000端口并建立tcp连接
func (s *Server) Run(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()
	log.Printf("Server listening on %s...\n", address)

	// 等待并处理客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("New client connected:", conn.RemoteAddr())

		// 每个连接启动一个 goroutine 来处理
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	// 创建读取器来读取客户端请求
	reader := bufio.NewReader(conn)

	for {
		data, err := getRequstData(reader)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected.")
				return
			}
			fmt.Printf("Error reading data: %v\n", err)
			continue
		}

		fmt.Printf("Request received:%v\n", data)

		if len(data) == 0 {
			continue
		}

		go func(reqdata []string) {
			// 获取一个请求chan
			rsp := s.GetChan()
			// 向通道写入数据并请求
			s.Request(rsp, reqdata)
			// 读取处理结果
			rspData := <-rsp
			//将chan放回空闲队列
			s.DisengagedQueue <- rsp
			// 发送数据
			_, err := conn.Write([]byte(rspData[0]))
			if err != nil {
				fmt.Println("Error sending response:", err)
				return
			}
			// fmt.Println("Response sent:", string(rspData))
		}(append([]string(nil), data...))
	}
}

func getRequstData(reader *bufio.Reader) ([]string, error) {
	// 读消息数量
	count, err := ParseArrayLength(reader)
	if err != nil {
		return nil, err
	}

	// 用于存储所有的消息
	var messages []string

	// 解析每个消息
	for i := 0; i < count; i++ {
		// 读取数据
		msgContent, err := ParseData(reader)
		if err != nil {
			return nil, err
		}

		// 将消息转为字符串
		messages = append(messages, string(msgContent))
	}

	// 打印解析后的消息
	// fmt.Println("解析到的消息:", messages)
	return messages, nil
}
