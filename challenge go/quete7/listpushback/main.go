package main

import (
	"fmt"
)

func main() {

	link := &List{}

	ListPushBack(link, "Hello")
	fmt.Println(link.Head)
	ListPushBack(link, "man")
	fmt.Println(link.Head.Next)
	ListPushBack(link, "how are you")
	fmt.Println(link.Head.Next.Next)
	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
}
