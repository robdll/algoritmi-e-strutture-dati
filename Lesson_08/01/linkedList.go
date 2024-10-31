package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type listNode struct {
	value int
	next *listNode
}

type linkedList struct {
	head *listNode
}

func main() {
	
	reader := bufio.NewReader(os.Stdin)
	list := linkedList{}

	fmt.Println("Questo programma permette di creare una lista concatenata e offre differenti funzioni per la sua manipolazione.")
	fmt.Println("Istruzioni disponibili:")
	fmt.Println("+ n | se n non appartiene all'insieme lo inserisce")
	fmt.Println("- n | Se n appartiene all’insieme lo elimina")
	fmt.Println("? n | stampa un messaggio se n appartiene all’insieme")
	fmt.Println("c   | stampa il numero di elementi dell'insieme")
	fmt.Println("p   | stampa gli elementi dell'insieme in ordine")
	fmt.Println("o   | stampa gli elementi dell'insieme in ordine inverso")
	fmt.Println("d   | cancella tutti gli elementi dell'insieme")
	fmt.Println("f   | termina il programma")

	for {
		fmt.Println("Inserire un istruzione:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Remove any trailing newline character
		parameters := strings.Split(input, " ")
		command := parameters[0]
		var value int
		if(len(parameters) > 1) {
			val, err := strconv.Atoi(parameters[1])
			if err != nil {
				fmt.Println("Errore: il valore deve essere un intero")
				continue
			}
			value = val
		}



		switch command {
			case "+":
				addNode(&list, value)
			case "-":
				removeNode(&list, value)
			case "?":
				searchList(&list, value)
			case "c":
				countElements(&list)
			case "p":
				printList(&list, true)
			case "o":
				printList(&list, false)
			case "d":
				list.head = nil
			case "f":
				return
			default:
				fmt.Println("Comando non valido")
		}
	}
	
}


func newNode (value int) *listNode {
	return &listNode{value: value}
}

func addNode (list *linkedList, value int) {
	node := newNode(value)
	node.next = list.head
	list.head = node
	fmt.Println("Valore inserito")
}

func printList (list *linkedList, ascend bool) {
	var elements []int
	for node := list.head; node != nil; node = node.next {
		elements = append(elements, node.value)
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

func searchList (list *linkedList, value int) {
	found := false
	for node := list.head; node != nil; node = node.next {
		if node.value == value {
			fmt.Println("Il valore ", value, " è presente nella lista")
			found = true
		}
	}
	if !found {
		fmt.Println("Il valore ", value, " non è presente nella lista")
	}
}

func removeNode (list *linkedList, value int) {
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

func countElements (list *linkedList) {
	count := 0
	for node := list.head; node != nil; node = node.next {
		count++
	}
	fmt.Println("Total elements in the list:", count)
}