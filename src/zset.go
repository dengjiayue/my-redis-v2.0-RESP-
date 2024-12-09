package src

import (
	"fmt"
	"strconv"
)

// 实现my-redis多数据类型支持
// 实现zset类型

// 实现zset类型的zadd方法
func (s *Server) Zadd(data []string) error {
	// 检查参数数量是否正确
	if len(data) < 3 || len(data)%2 != 1 {
		return fmt.Errorf("wrong number of arguments for 'zadd' command")
	}
	// 检查key是否存在
	if _, ok := s.MZset[data[0]]; !ok {
		// 如果不存在，创建一个新的空map
		s.MZset[data[0]].Hash = make(map[string]bool)
	}
	for i := 1; i < len(data)-1; i += 2 {
		// 获取score
		score, err := strconv.ParseUint(data[i], 10, 64)
		if err != nil {
			return err
		}
		// 设置field和value
		s.MZset[data[0]].Hash[data[i+1]] = true
		//向跳跃表中添加数据
		s.MZset[data[0]].SkipList.Add(score, data[i+1])
	}
	// 返回nil表示成功
	return nil
}
