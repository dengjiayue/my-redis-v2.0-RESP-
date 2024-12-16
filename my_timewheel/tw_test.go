package mytimewheel

import (
	"log"
	"testing"
	"time"
)

func TestTw(t *testing.T) {
	ch := make(chan struct{})
	Delay(2*time.Second, "test", func() {
		ch <- struct{}{}
	})
	time.Sleep(time.Second * 1)
	//更换延迟时间
	Delay(time.Second*5, "test", func() {
		ch <- struct{}{}
	})
	<-ch
	// log.Println("test")
}

func TestTw2(t *testing.T) {
	Delay(time.Second*3, "test2", func() {
		log.Println("test2")
	})
	time.Sleep(time.Second * 1)
	v := tw.Slots[3].Traverse()
	log.Println(v)
	Delay(time.Second*5, "test2", func() {
		log.Println("test2")
	})
	v2 := tw.Slots[2].Traverse()
	log.Println(v2)
}
