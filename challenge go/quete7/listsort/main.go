package main

import (
	"fmt"
)

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *NodeI

	link = listPushBack(link, 5)
	link = listPushBack(link, 4)
	link = listPushBack(link, 3)
	link = listPushBack(link, 2)
	link = listPushBack(link, 1)

	PrintList(ListSort(link))
}

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	var node_list []NodeI
	for l != nil {
		node_list = append(node_list, *l)
		l = l.Next
	}
	for i := 0; i < len(node_list); i++ {
		for j := i + 1; j < len(node_list); j++ {
			if node_list[i].Data > node_list[j].Data {
				node_list[i], node_list[j] = node_list[j], node_list[i]
			}
		}
	}
	var new_list *NodeI
	for i := 0; i < len(node_list); i++ {
		new_list = listPushBack(new_list, node_list[i].Data)
	}
	return new_list
}
