package mytimewheel

import "time"

func Delay(delay time.Duration, key string, job func()) {
	tw.AddTaskChan <- &task{
		delay: delay,
		key:   key,
		job:   job,
	}
}

func RemoveTask(key string) {
	tw.DeleteTaskChan <- key
}

func StopTimeWheel() {
	tw.CloseChan <- struct{}{}
}
