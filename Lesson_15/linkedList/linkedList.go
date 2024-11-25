package linkedlist

import "fmt"

type ListNode struct {
	value interface{} // use interface{} to be generic type
	next  *ListNode
}

type LinkedList struct {
	head *ListNode
}

// creates and returns a new LinkedList.
func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func newNode(value interface{}) *ListNode {
	return &ListNode{value: value}
}

// complessità O(1)
func AddNode(list *LinkedList, value interface{}) {
	node := newNode(value)
	node.next = list.head
	list.head = node
	fmt.Println("Valore inserito")
}

// complessità O(n)
func printList(list *LinkedList, ascend bool) {
	var elements []int
	for node := list.head; node != nil; node = node.next {

		if value, ok := node.value.(int); ok {
			elements = append(elements, value)
		} else {

			fmt.Println("Errore: tipo non valido per node.value")
		}
	}
	if ascend {
		for i := len(elements) - 1; i >= 0; i-- {
			fmt.Println(elements[i])
		}
	} else {
		for i := 0; i < len(elements); i++ {
			fmt.Println(elements[i])
		}
	}
}

// complessità O(n)
func SearchList(list *LinkedList, value interface{}) bool {
	found := false
	for node := list.head; node != nil; node = node.next {
		if node.value == value {
			found = true
		}
	}
	return found
}

// complessità O(n)
func removeNode(list *LinkedList, value int) {
	if list.head == nil {
		fmt.Println("Nessun elemento nella lista")
		return
	}
	if list.head.value == value {
		list.head = list.head.next
		fmt.Println("Valore rimosso")
		return
	}
	for node := list.head; node.next != nil; node = node.next {
		if node.next.value == value {
			node.next = node.next.next
			fmt.Println("Valore rimosso")
			return
		}
	}
}

// complessità O(n)
func countElements(list *LinkedList) {
	count := 0
	for node := list.head; node != nil; node = node.next {
		count++
	}
	fmt.Println("Total elements in the list:", count)
}
