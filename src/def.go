package src

const (
	SET = 1
	GET = 2
)

type Server struct {
	M               map[string]string
	DealQueue       chan chan []string //请求处理通道
	DisengagedQueue chan chan []string //空闲的处理通道队列
	// MaxChan         uint32                  // 最大通道数量
}

func NewServer(max_chan uint32) *Server {
	q := make(chan chan []string, max_chan)
	//初始化空闲的处理通道队列
	for i := uint32(0); i < max_chan; i++ {
		q <- (make(chan []string))
	}
	return &Server{
		M:               make(map[string]string),
		DealQueue:       make(chan chan []string),
		DisengagedQueue: q,
		// MaxChan:         max_chan,
	}
}

// 优雅的关闭
func (s *Server) Close() {
	// 关闭请求处理通道队列
	close(s.DealQueue)
	//关闭所有请求通道
	for ch := range s.DealQueue {
		close(ch)
	}
	// 关闭空闲的处理通道队列
	close(s.DisengagedQueue)
	// 关闭所有空闲的处理通道队列
	for ch := range s.DisengagedQueue {
		close(ch)
	}

}
