package main

import "fmt"

type node struct {
	val  int   // value of the node
	next *node // pointer to the next node
}

func main() {
	head := &node{val: 3}
	head.next = &node{val: 2}
	head.next.next = &node{val: 0}
	head.next.next.next = &node{val: 1}
	head.next.next.next.next = &node{val: 7}

	f(head, 1)	// 7
	f(head, 5)	// 3
	f(head, 10) // nothing
}

// Ricevendo il puntatore alla testa di una lista e un intero k la funzione f 
// stampa il valore alla posizione lunghezza - k e restituisce la lunghezza della lista.
// la complessità è O(n) dove n è la lunghezza della lista.
func f( p *node, k int ) int {
	var a int
	if p == nil {
		return 0
	}
	a = 1 + f(p.next, k) 
	if a == k { 
		fmt.Println(p.val)
	}
	return a
}