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
		case "hset", "HSET":
			err := s.Hset(data[1:])
			if err != nil {
				ch <- []string{err.Error()}
			} else {
				ch <- []string{"+OK"}
			}
		case "hget", "HGET":
			value, err := s.Hget(data[1:])
			if err != nil {
				ch <- []string{ErrorResponse(err.Error())}
			} else {
				ch <- []string{SingleStringResponse(value)}
			}
		case "hgetall", "HGETALL":
			value, err := s.Hgetall(data[1:])
			if err != nil {
				ch <- []string{ErrorResponse(err.Error())}
			} else {
				ch <- value
			}
		case "lpush", "LPUSH":
			err := s.Lpush(data[1:])
			if err != nil {
				ch <- []string{ErrorResponse(err.Error())}
			} else {
				ch <- []string{SuccessResponse()}
			}
		case "rpop", "RPOP":
			value, err := s.Rpop(data[1:])
			if err != nil {
				ch <- []string{ErrorResponse(err.Error())}
			} else {
				ch <- []string{SingleStringResponse(value)}
			}
		case "llen", "LLEN":
			value, err := s.Llen(data[1:])
			if err != nil {
				ch <- []string{ErrorResponse(err.Error())}
			} else {
				ch <- []string{IntegerResponse(value)}
			}
		case "sadd", "SADD":
			err := s.Sadd(data[1:])
			if err != nil {
				ch <- []string{ErrorResponse(err.Error())}
			} else {
				ch <- []string{SuccessResponse()}
			}
		case "smembers", "SMEMBERS":
			value, err := s.Smembers(data[1:])
			if err != nil {
				ch <- []string{ErrorResponse(err.Error())}
			} else {
				ch <- value
			}
		case "zadd", "ZADD":
			err := s.Zadd(data[1:])
			if err != nil {
				ch <- []string{ErrorResponse(err.Error())}
			} else {
				ch <- []string{SuccessResponse()}
			}

		default:
			ch <- []string{"请求类型错误"}
		}
	}
}

func (s *Server) Set(data []string, ch chan []string) {

	// 处理请求
	s.M[data[0]] = data[1]
	// 返回结果(0代表成功)
	ch <- []string{SuccessResponse()}
}

func (s *Server) Get(data string, ch chan []string) {
	// 处理请求
	// 从map中获取值
	value, ok := s.M[data]
	if !ok {
		ch <- []string{""}
	} else {
		ch <- []string{SingleStringResponse(value)}
	}
}
