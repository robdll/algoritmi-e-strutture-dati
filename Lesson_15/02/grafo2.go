package main

import (
	"fmt"
	"strconv"
	"strings"

	LinkedListPackage "Grafi/linkedList"
)

type grafo struct {
	n         int
	adiacenti []LinkedListPackage.LinkedList
}

type vertice struct {
	informazioni informazioni
	chiave       string
	adj          []*vertice // insieme dei vertici adiacenti
}
type informazioni struct {
	nome  string
	età   int
	hobby string
}

func main() {
	var n int
	fmt.Print("Enter the number of nodes: ")
	fmt.Scan(&n)
	g := nuovoGrafo(n)
	initGrafo(g, n)
	// fmt.Println("The graph is:")
	// printGrafo(g)
	// checkEdgeExistence(g)

}

func nuovoGrafo(n int) *grafo {
	var nome, hobby string
	var età int

	// Slice di liste collegate per rappresentare le adiacenze
	adiacenti := make([]LinkedListPackage.LinkedList, n)

	for i := 0; i < n; i++ {
		// Inizializza una nuova lista collegata per il nodo corrente
		adiacenti[i] = *LinkedListPackage.NewLinkedList()

		// Raccogli le informazioni dell'utente
		fmt.Printf("Nodo %d:\n", i+1)
		fmt.Print("Inserisci il nome: ")
		fmt.Scan(&nome)
		fmt.Print("Inserisci l'età: ")
		fmt.Scan(&età)
		fmt.Print("Inserisci l'hobby: ")
		fmt.Scan(&hobby)

		// Crea un vertice con le informazioni raccolte
		verticeInfo := vertice{
			informazioni: informazioni{
				nome:  nome,
				età:   età,
				hobby: hobby,
			},
			chiave: strconv.Itoa(i),
			adj:    []*vertice{},
		}

		// Aggiungi il vertice alla lista collegata
		LinkedListPackage.AddNode(&adiacenti[i], verticeInfo)
	}

	// Crea e restituisci il grafo
	return &grafo{n: n, adiacenti: adiacenti}
}

func initGrafo(g *grafo, n int) {
	// ask the user to insert the edges of each node
	for i := 0; i < n; i++ {
		fmt.Print("Enter the edges of node ", i, " separated by comma: ")
		var edges string
		fmt.Scan(&edges)
		// split the string by comma  (e.g. "1,2,3" -> ["1", "2", "3"])
		edgeList := strings.Split(edges, ",")
		for _, edge := range edgeList {
			num, err := strconv.Atoi(edge)
			// if string is not a comma separated list of integers, or if the integer is pointing to a not-existing node, retry the current node
			if err != nil || num < (-1) || num >= n {
				fmt.Println("Invalid input. Please enter a comma separated list of integers.")
				// decrement i to retry the current node
				i--
				break
			} else if num != -1 {
				LinkedListPackage.AddNode(&g.adiacenti[i], num)
			}
		}
	}
}

func printGrafo(g *grafo) {
	for i := 0; i < g.n; i++ {
		fmt.Print("  ", i)
	}
	for i := 0; i < g.n; i++ {
		fmt.Println()
		fmt.Print(i, " ")
		currentList := g.adiacenti[i]
		for j := 0; j < g.n; j++ {
			hasEdge := LinkedListPackage.SearchList(&currentList, j)
			if hasEdge {
				fmt.Print("1  ")
			} else {
				fmt.Print("0  ")
			}
		}
	}
}

func checkEdgeExistence(g *grafo) {
	var input string
	fmt.Print("\nEnter the two nodes to check if there is an edge between them (separated by comma): ")
	fmt.Scan(&input)
	// split the string by comma  (e.g. "1,2" -> ["1", "2"])
	nodes := strings.Split(input, ",")
	firstInt, err1 := strconv.Atoi(nodes[0])
	secondInt, err2 := strconv.Atoi(nodes[1])
	// if string is not a comma separated list of integers, or if the integer is pointing to a not-existing node, retry the current node
	if err1 != nil || err2 != nil || firstInt < 0 || firstInt >= g.n || secondInt < 0 || secondInt >= g.n {
		fmt.Println("Invalid input. Please enter a comma separated list of integers.")
		checkEdgeExistence(g)
	} else {
		hasEdge := LinkedListPackage.SearchList(&g.adiacenti[firstInt], secondInt)
		if hasEdge {
			fmt.Println("The edge exists.")
		} else {
			fmt.Println("The edge does not exist.")
		}
	}

}
