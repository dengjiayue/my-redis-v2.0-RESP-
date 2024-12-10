package src

import (
	"my_redis/src/timewheel"
	"strconv"
	"time"
)

//过期时间实现
//

// 扫描部分key并检查是否过期
func (s *Server) CheckExpire(num int) {
	// 遍历所有key
	for k, v := range s.Ex {
		if v < time.Now().Unix() {
			// 过期了
			delete(s.Ex, k)
			// 删除key
			delete(s.M, k)
		}
		num--
		if num == 0 {
			break
		}
	}
}

// 定时过期
func (s *Server) TimingExpire() {
	// 定时执行
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		// 定时执行
		s.CheckExpire(100)
	}
}

// 惰性过期(输出是否过期 )
func (s *Server) LazyExpire(key string) bool {
	// 检查key是否过期
	if s.Ex[key] < time.Now().Unix() {
		// 过期了
		delete(s.Ex, key)
		// 删除key
		delete(s.M, key)
		return true
	}
	return false
}

// 立即过期
// 使用时间轮实现立即过期
// exTime为毫秒过期时间
func (s *Server) ImmediatelyExpire(key string, exTime int64) {
	late := time.Millisecond * time.Duration(exTime)
	timewheel.Delay(late, key, func() {
		// 删除数据
		delete(s.Ex, key)
		// 过期时间记录
		delete(s.M, key)
	})
}

// 设置过期时间: 向Ex写数据
// exTime为过期时间(毫秒)
func (s *Server) SetExTime(key string, exTime string) error {
	//将exTime转换为数字
	exTimeNum, err := strconv.Atoi(exTime)
	if err != nil {
		return err
	}
	// 将exTime转换为过期时间戳(毫秒级)
	ex := time.Now().UnixMilli() + int64(exTimeNum)
	s.Ex[key] = ex
	return nil
}
