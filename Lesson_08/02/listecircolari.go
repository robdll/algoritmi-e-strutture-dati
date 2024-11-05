package main

import . "fmt"

type listnode struct {
	valore int
	next   *listnode
}

type linkedList struct {
	head *listnode
}

func newNode(valore int) *listnode {
	return &listnode{valore: valore}
}

func addNode(list *linkedList, valore int) {
	node := newNode(valore)
	node.next = list.head
	list.head = node
	Println("Valore inserito")
}

func main() {

	list := linkedList{}
	input := 20

	addNode(&list, input)
	Println(list, list.head)
	Print(list.head.valore)
}
