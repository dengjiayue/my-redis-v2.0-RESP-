package src

import (
	"fmt"
	"strconv"
)

// 封装my redis多数据类型支持

// list 类型
// 使用字符串数组实现

// 实现list类型的lpush方法
func (s *Server) Lpush(data []string) error {
	// 检查参数数量
	if len(data) < 3 {
		return fmt.Errorf("invalid number of arguments")
	}
	// 检查key是否存在
	if _, ok := s.MList[data[0]]; !ok {
		// 如果不存在，创建一个新的数组
		s.MList[data[0]] = make([]string, 0)
	}
	// 将元素添加到数组的结尾
	s.MList[data[0]] = append(s.MList[data[0]], data[1:]...)
	// 返回成功
	return nil
}

// 实现list类型的rpop方法
func (s *Server) Rpop(data []string) (string, error) {
	// 检查参数数量
	if len(data) != 1 {
		return "", fmt.Errorf("invalid number of arguments")
	}
	// 检查key是否存在
	if _, ok := s.MList[data[0]]; !ok {
		// 如果不存在，返回空数组
		return "", nil
	}
	// 获取数组的第一个元素
	last := s.MList[data[0]][0]
	// 删除数组的第一个元素
	s.MList[data[0]] = s.MList[data[0]][1:]
	// 返回元素
	return last, nil
}

// 实现list类型的lrange方法
func (s *Server) Lrange(data []string) ([]string, error) {
	// 检查参数数量
	if len(data) != 3 {
		return nil, fmt.Errorf("invalid number of arguments")
	}
	// 检查key是否存在
	if _, ok := s.MList[data[0]]; !ok {
		// 如果不存在，返回空数组
		return nil, nil
	}
	// 检查start和end是否合法
	start, err := strconv.Atoi(data[1])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(data[2])
	if err != nil {
		return nil, err
	}
	if start < 0 || end >= len(s.MList[data[0]]) {
		return nil, fmt.Errorf("invalid start or end index")
	}
	// 返回元素
	return s.MList[data[0]][start : end+1], nil
}

// 实现list类型的llen方法
func (s *Server) Llen(data []string) (int, error) {
	// 检查参数数量
	if len(data) != 1 {
		return 0, fmt.Errorf("invalid number of arguments")
	}
	// 检查key是否存在
	if _, ok := s.MList[data[0]]; !ok {
		// 如果不存在，返回0
		return 0, nil
	}
	// 返回数组的长度
	return len(s.MList[data[0]]), nil
}
