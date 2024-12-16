package mylist

import (
	"log"
	"testing"
)

func TestNode_Remove(t *testing.T) {
	l := New()
	n := l.AddToHead(1)
	 l.AddToTail(2)
	v1 := l.Traverse()
	log.Println(v1)
	n.Remove()
	v2 := l.Traverse()
	log.Println(v2)
	// n2.Remove()
	v4 := l.PopTail()
	log.Println(v4)
}
