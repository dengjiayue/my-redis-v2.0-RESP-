package mytimewheel

import (
	"log"
	mylist "my_redis/my_list"
	"time"
)

var tw *TimeWheel

func init() {
	tw = New(time.Second, 3000)
	tw.Start()
}

// 实现时间轮算法
type TimeWheel struct {
	Interval       time.Duration       //时间格的基本时间跨度
	Ticker         *time.Ticker        // 时间轮的计时器
	Slots          []*mylist.List      //时间格
	CurrentPos     int                 //当前指针，指向时间格的下标
	Locations      map[string]location //记录任务（key）所在的位置
	SlotNum        int                 //时间格的数量
	AddTaskChan    chan *task          //添加任务的通道
	DeleteTaskChan chan string         //删除任务的通道
	IsClose        bool                // 关闭的标志
	CloseChan      chan struct{}       // 关闭的通道
}

// 任务位置
type location struct {
	taskElem *mylist.Node //指向双向链表中任务节点的指针
}

// 任务
type task struct {
	key    string
	circle int           // 循环次数
	delay  time.Duration // 延迟时间
	job    func()        // 任务函数
}

// 初始化时间轮
func New(interval time.Duration, slotNum int) *TimeWheel {
	if interval <= 0 || slotNum <= 0 {
		return nil
	}
	tw := &TimeWheel{
		Interval:       interval,
		Slots:          make([]*mylist.List, slotNum),
		CurrentPos:     0,
		Locations:      make(map[string]location),
		SlotNum:        slotNum,
		AddTaskChan:    make(chan *task),
		DeleteTaskChan: make(chan string),
		CloseChan:      make(chan struct{}),
	}
	tw.initSlots()
	return tw
}

// 初始化时间格任务链表
func (tw *TimeWheel) initSlots() {
	for i := 0; i < len(tw.Slots); i++ {
		tw.Slots[i] = mylist.New()
	}
}

// 启动时间轮
func (tw *TimeWheel) Start() {
	tw.Ticker = time.NewTicker(tw.Interval)
	go tw.run()
}

// 运行时间轮
func (tw *TimeWheel) run() {
	for {
		select {
		case <-tw.Ticker.C:
			tw.handelTask()
		case task := <-tw.AddTaskChan:
			tw.addTask(task)
		case key := <-tw.DeleteTaskChan:
			tw.deleteTask(key)
		case <-tw.CloseChan:
			tw.Ticker.Stop()
			tw.IsClose = true
			return
		}
	}
}

// 处理任务
func (tw *TimeWheel) handelTask() {
	// log.Printf("时间轮执行任务: %d", tw.CurrentPos)
	taskList := tw.Slots[tw.CurrentPos%tw.SlotNum]
	v := taskList.Traverse()
	log.Printf("时间 %d任务数量: %d", tw.CurrentPos, len(v))
	// log.Printf("任务数量: %d", taskList.Len())
	for {
		e := taskList.PopHead()
		if e == nil {
			break
		}
		log.Printf("任务: %v\n", e.(*task))
		task := e.(*task)
		if task.circle == 0 {
			go task.job()
			log.Printf("任务执行完毕: %s", task.key)
			delete(tw.Locations, task.key)
		} else {
			task.circle--
		}
	}
	tw.CurrentPos++
}

// 添加任务
func (tw *TimeWheel) addTask(task *task) {
	//查看任务 key 是否存在
	if _, ok := tw.Locations[task.key]; ok {
		// 删除
		tw.deleteTask(task.key)
	}
	// 计算任务所在的位置
	pos := (tw.CurrentPos + int(task.delay/tw.Interval)) % tw.SlotNum
	// 计算任务的循环次数(圈数=(当前时间格+任务延迟时间)/时间格数量)(舍弃小数,只取整数部分)
	circle := (tw.CurrentPos + int(task.delay/tw.Interval)) / tw.SlotNum

	task.circle = circle

	e := tw.Slots[pos].AddToHead(task)
	tw.Locations[task.key] = location{
		taskElem: e,
	}
	log.Printf("%d 任务添加成功: %s", pos, task.key)
}

// 删除任务
func (tw *TimeWheel) deleteTask(key string) {
	// 查看任务 key 是否存在
	if _, ok := tw.Locations[key]; !ok {
		return
	}
	loc := tw.Locations[key]
	err := loc.taskElem.Remove()
	if err != nil {
		log.Printf("删除任务失败: %s", key)
		return
	}
	log.Printf("任务删除成功: %s", key)
	delete(tw.Locations, key)
}
