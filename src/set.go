package src

import "fmt"

//实现 my-redis多数据类型支持
//实现set(集合)类型
//使用map[string]bool来实现set类型

// 实现set类型的sadd方法
func (s *Server) Sadd(data []string) error {
	// 检查参数数量是否正确
	if len(data) > 2 {
		return fmt.Errorf("wrong number of arguments for 'sadd' command")
	}
	// 检查key是否存在
	if _, ok := s.MSet[data[0]]; !ok {
		// 如果不存在，创建一个新的空的map
		s.MSet[data[0]] = make(map[string]bool)
	}
	for i := 1; i < len(data); i++ {
		// 设置field和value
		s.MSet[data[0]][data[i]] = true
	}
	// 返回nil表示成功
	return nil
}

// 实现set类型的smembers方法
func (s *Server) Smembers(data []string) ([]string, error) {
	// 检查参数数量是否正确
	if len(data) != 1 {
		return nil, fmt.Errorf("wrong number of arguments for 'smembers' command")
	}
	// 检查key是否存在
	if _, ok := s.MSet[data[0]]; !ok {
		// 如果不存在，返回错误
		return nil, nil
	}
	// 初始化一个空的slice
	set := []string{}
	// 遍历map，将所有的field添加到slice中
	for k := range s.MSet[data[0]] {
		set = append(set, k)
	}
	// 返回slice
	return set, nil
}
