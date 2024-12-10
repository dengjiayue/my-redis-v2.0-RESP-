package src

const (
	SET = 1
	GET = 2
)

// 有序集合: 使用跳跃表+哈希表实现
type Zset struct {
	//跳跃表
	SkipList *SkipList
	//哈希表
	Hash map[string]bool
}

// 跳跃表
type SkipList struct {
	Head *SkipListNode
	Tail *SkipListNode
	Len  int
}

// 跳表的初始化
func NewSkipList() *SkipList {
	head := &SkipListNode{
		Front: nil,
		Back:  nil,
		Up:    nil,
		Down:  nil,
		Score: 0,
		Value: "",
	}
	tail := &SkipListNode{
		Front: nil,
		Back:  nil,
		Up:    nil,
		Down:  nil,
		Score: 0,
		Value: "",
	}
	head.Front = tail
	tail.Back = head
	return &SkipList{
		Head: head,
		Tail: tail,
		Len:  0,
	}
}

//实现跳跃表增加与删除数据

// 增加
func (s *SkipList) Add(score uint64, value string) error {
	//TODO : 实现跳跃表增加...
	return nil
}

// 删除
func (s *SkipList) Remove(score uint64, value string) error {
	//TODO : 实现跳跃表删除...
	return nil
}

// 跳跃表节点
type SkipListNode struct {
	Front *SkipListNode
	Back  *SkipListNode
	Up    *SkipListNode
	Down  *SkipListNode
	Score uint64
	Value string
}

type Server struct {
	M               map[string]string
	DealQueue       chan chan []string //请求处理通道
	DisengagedQueue chan chan []string //空闲的处理通道队列
	MHash           map[string]map[string]string
	MList           map[string][]string
	MSet            map[string]map[string]bool
	MZset           map[string]*Zset
	Ex              map[string]int64 // 过期时间字典
	// MStream         map[string]map[string]string
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
