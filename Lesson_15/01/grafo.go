package main

import (
	"fmt"
	LinkedList "linkedList_project/linkedList"
	"strconv"
	"strings"
) 

type grafo struct {
	n int 
	adiacenti []LinkedList.LinkedList
}

func main() {
	var n int
	fmt.Print("Enter the number of nodes: ")
	fmt.Scan(&n)
	g := nuovoGrafo(n)
	initGrafo(g, n)
}


func nuovoGrafo (n int) *grafo {
	s := LinkedList.NewLinkedList()
	for i := 0; i < n; i++ {
		LinkedList.AddNode(s, i)
	}
	return &grafo{n, make([]LinkedList.LinkedList, n)}
}

func initGrafo (g *grafo, n int) {
	// ask the user to insert the edges of each node
	for i := 0; i < n; i++ {
		fmt.Print("Enter the edges of node ", i, " separated by comma: ")
		var edges string
		fmt.Scan(&edges)
		edgeList := strings.Split(edges, ",")
		for _, edge := range edgeList {
			if _, err := strconv.Atoi(edge); err != nil {
				fmt.Println("Invalid input. Please enter a comma separated list of integers.")
				i-- // decrement i to retry the current node
				break
			}
		}
	}
}