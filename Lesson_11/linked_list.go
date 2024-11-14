package main

import "fmt"

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
	size int
}

/*
func newNode(val int) *listNode {
	return &listNode{val, nil}
}*/

/*
func newNode(val int) *listNode {
	var new listNode
	new.item = val
	new.next = nil // valore di default
	return &new
}*/

func newNode(val int) *listNode {
	node := new(listNode)
	node.item = val
	node.next = nil // valore di default
	return node
}

func len(l linkedList) int {
	return l.size
}

// se inserisco in testa, devo riassegnare la lista, quindi va restituita
func addNewNode(l linkedList, val int) linkedList {
	node := newNode(val)
	node.next = l.head
	l.head = node // modifica il campo head della struct l locale alla funzione!
	fmt.Print("l dentro func addNewNode: ")
	printList(l)
	return l
}

// passo il riferimento alla struttura lista, modifico direttamente la sua head, senza bisogno di restituire
func addNewNodePointer(l *linkedList, val int) {
	node := newNode(val)
	node.next = l.head
	l.head = node // modifica ll campo head della struct list riferita da l!
	fmt.Print("l dentro func addNewNode: ")
	printList(*l)
}

func printList(l linkedList) {
	p := l.head
	for p != nil {
		fmt.Print(p.item, " ")
		p = p.next
	}
	fmt.Println()
}

func searchList(l linkedList, val int) (bool, *listNode) {
	p := l.head
	for p != nil {
		if p.item == val {
			return true, p
		}
		p = p.next
	}
	return false, nil
}

func deleteItem(l *linkedList, val int) bool {
	var curr, prev *listNode = l.head, nil
	for curr != nil {
		if curr.item == val {
			if prev == nil { // trovato in testa
				l.head = l.head.next
			} else {
				prev.next = curr.next
			}
			return true
		}
		prev = curr
		curr = curr.next
	}

	return false // val non trovato
}

func main() {
	var list linkedList

	var n listNode // problematico: se riutilizzo n, sovrascrive la struttura che fa da nodo iniziale. Invece ogni volta bisogna creare un nuovo nodo!
	n.item = 20
	n.next = nil
	list.head = &n

	/*
		n.item = 50
		n.next = list.head
		list.head = &n

		printList(list) // list ha un solo nodo, con next che punta a sè stesso (loop infinito)
	*/

	/*
		var node *listNode
		node = new(listNode) // crea una nuova struttura nodo
		(*node).item = 10
		(*node).next = list.head
		list.head = node

		fmt.Println("list:", list)
		fmt.Println("list.head:", list.head)

		node = new(listNode)
		node.item = 7 // non è necessario dereferenziare
		node.next = list.head
		list.head = node

		fmt.Println("list:", list)
		fmt.Println("list.head:", list.head)

		node = &listNode{30, list.head} // un altro modo per creare una struttura listNode
		list.head = node
		fmt.Println("list:", list)
		fmt.Println("list.head:", list.head)

		printList(list)*/

	var l2 linkedList

	addNewNode(l2, 100)
	printList(l2)
	fmt.Println(l2) // lista vuota - il campo head di l2 non è stato modificato

	l2 = addNewNode(l2, 200) // il campo head di l2 viene modificato con l'assegnamento
	printList(l2)

	l2 = addNewNode(l2, 300) // il campo head di l2 viene modificato con l'assegnamento
	printList(l2)

	addNewNodePointer(&l2, 80) // il campo head di l2 viene modificato dalla funzione
	printList(l2)

	res, pointer := searchList(l2, 80)
	fmt.Println("search 80:", res, pointer)
	res, pointer = searchList(l2, 800)
	fmt.Println("search 800:", res, pointer)

	fmt.Println(deleteItem(&l2, 800))
	printList(l2)

	fmt.Println(deleteItem(&l2, 300))
	printList(l2)

	fmt.Println(deleteItem(&l2, 200))
	printList(l2)
}
