package mylist

import "fmt"

type List struct {
	Head *Node
	Tail *Node
	// Len  int
}

type Node struct {
	Prev *Node
	Next *Node
	Val  interface{}
}

func New() *List {
	head := new(Node)
	tail := new(Node)
	head.Next = tail
	tail.Prev = head
	return &List{Head: head, Tail: tail}
}

// AddToHead 添加到头部
func (l *List) AddToHead(v interface{}) *Node {
	newNode := &Node{
		Prev: l.Head,
		Next: l.Head.Next,
		Val:  v,
	}
	// l.Len++
	// 将新节点插入到head和head.Next之间
	newNode.Next = l.Head.Next
	newNode.Prev = l.Head
	l.Head.Next = newNode
	newNode.Next.Prev = newNode
	return l.Head.Next
}

// AddToTail 添加到尾部
func (l *List) AddToTail(v interface{}) *Node {
	newNode := &Node{
		Prev: l.Tail.Prev,
		Next: l.Tail,
		Val:  v,
	}
	// l.Len++
	// 将新节点插入到tail和tail.Prev之间
	newNode.Next = l.Tail
	newNode.Prev = l.Tail.Prev
	l.Tail.Prev = newNode
	newNode.Prev.Next = newNode
	return l.Tail.Prev
}

// Remove 删除节点
func (n *Node) Remove() error {
	if n == nil || n.Prev == nil || n.Next == nil {
		return fmt.Errorf("node is a dummy node")
	}
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	return nil
}

// 弹出一个头节点值
func (l *List) PopHead() interface{} {
	if l.Head.Next == l.Tail {
		return nil
	}
	node := l.Head.Next
	node.Remove()
	return node.Val
}

// 弹出一个尾节点值
func (l *List) PopTail() interface{} {
	if l.Head.Next == l.Tail {
		return nil
	}
	node := l.Tail.Prev
	node.Remove()
	return node.Val
}

// 遍历链表
func (l *List) Traverse() []interface{} {
	var res []interface{}
	node := l.Head.Next
	for node != l.Tail || node == nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return res
}
