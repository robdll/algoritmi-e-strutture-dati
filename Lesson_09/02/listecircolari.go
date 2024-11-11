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
	Println(list.head)
	Println(list.head.valore)
	input2 := 30
	addNode(&list, input2)
	Println(list.head)
	Println(list.head.valore)
	Println(list.head.next)

}

