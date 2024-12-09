package src

import "fmt"

// 封装my redis多数据类型支持

// hash 类型

// 实现hash类型的hset方法
func (s *Server) Hset(data []string) error {
	// 检查参数数量是否正确
	if len(data) > 2 {
		return fmt.Errorf("wrong number of arguments for 'hset' command")
	}
	// 检查key是否存在
	if _, ok := s.MHash[data[0]]; !ok {
		// 如果不存在，创建一个新的空map
		s.MHash[data[0]] = make(map[string]string)
	}
	for i := 1; i < len(data)-1; i += 2 {
		// 设置field和value
		s.MHash[data[0]][data[i]] = data[i+1]
	}
	// 返回nil表示成功
	return nil
}

// 实现hash类型的hget方法
func (s *Server) Hget(data []string) (string, error) {
	// 检查参数数量是否正确
	if len(data) != 2 {
		return "", fmt.Errorf("wrong number of arguments for 'hget' command")
	}
	// 检查key是否存在
	if _, ok := s.MHash[data[0]]; !ok {
		// 如果不存在，返回错误
		return "", nil
	}
	// 获取field对应的value
	value, ok := s.MHash[data[0]][data[1]]
	if !ok {
		// 如果field不存在，返回错误
		return "", nil
	}
	// 返回value
	return value, nil
}

func (s *Server) Hgetall(data []string) ([]string, error) {
	// 检查参数数量是否正确
	if len(data) != 1 {
		return nil, fmt.Errorf("wrong number of arguments for 'hgetall' command")
	}
	// 检查key是否存在
	if _, ok := s.MHash[data[0]]; !ok {
		// 如果不存在，返回错误
		return nil, nil
	}
	// 初始化一个空的slice
	hash := []string{}
	// 获取所有field和value
	for k, v := range s.MHash[data[0]] {
		hash = append(hash, k, v)

	}
	// 返回所有field和value
	return hash, nil
}
