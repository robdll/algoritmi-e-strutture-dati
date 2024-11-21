package main

import (
	. "fmt"
)

type listNode struct {
	intero    int
	puntatore *listNode
}

type linkedList struct {
	iniziolista *listNode
}

func main() {

	var node *listNode
	node = new(listNode)
	Println(node)
	node.intero = 10
	Println(node)
	Println(node.intero)

}
