package mytimewheel

import (
	"time"
)

func Delay(delay time.Duration, key string, job func()) {
	//异步添加任务
	go func() {
		tw.AddTaskChan <- &task{
			delay: delay,
			key:   key,
			job:   job,
		}
	}()
	// log.Printf("add task: %s", key)
}

func RemoveTask(key string) {
	tw.DeleteTaskChan <- key
}

func StopTimeWheel() {
	tw.CloseChan <- struct{}{}
}
