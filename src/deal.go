package src

func (s *Server) Deal() {
	for {
		ch := <-s.DealQueue
		// 解析数据
		//获取第一个uint8的数据为请求类型
		data := <-ch

		// 处理请求
		// fmt.Println("Type:", Type)
		switch data[0] {
		case "set", "SET":
			s.Set(data[1:], ch)
		case "get", "GET":
			s.Get(data[1], ch)
		default:
			ch <- []string{"请求类型错误"}
		}
	}
}

func (s *Server) Set(data []string, ch chan []string) {

	// 处理请求
	s.M[data[0]] = data[1]
	// 返回结果(0代表成功)
	ch <- []string{"OK"}
	//放回空闲队列
	s.DisengagedQueue <- ch
}

func (s *Server) Get(data string, ch chan []string) {
	// 处理请求
	// 从map中获取值
	value, ok := s.M[data]
	if !ok {
		ch <- []string{""}
	} else {
		ch <- []string{value}
	}
}
