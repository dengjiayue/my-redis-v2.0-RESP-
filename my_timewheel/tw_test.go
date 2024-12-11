package mytimewheel

import (
	"testing"
	"time"
)

func TestTw(t *testing.T) {
	ch := make(chan struct{})
	Delay(time.Second, "test", func() {
		ch <- struct{}{}
	})
	<-ch
	// log.Println("test")
}
