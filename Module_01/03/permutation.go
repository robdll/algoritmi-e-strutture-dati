// Permutazione da riordinare

// Si consideri un vettore di N strutture, ciascuna composta da un nome e da una chiave identifica-
// tiva da 1 a N. Le chiavi sono uniche, cioè non sono ripetute.

// Progettate un algoritmo che riordini il vettore in questo modo: nella posizione 0 c’è l’elemento
// di chiave N, nella posizione 1 c’è l’elemento di chiave N −1 e così via fino alla posizione N −1
// che dovrà contenere l’elemento di chiave 1.
// Esempio. Se il vettore v contiene questi elementi:

// 6 Francesco, 1 Andrea, 5 Elisa, 2 Beatrice, 3 Carlo, 4 Dino, 7 Giorgia, 9 Irene, 8 Henry
// l’algoritmo deve riordinare il vettore in questo modo:
// 9 Irene, 8 Henry, 7 Giorgia, 6 Francesco, 5 Elisa, 4 Dino, 3 Carlo, 2 Beatrice, 1 Andrea

package main

import "fmt"

type person struct {
	name string
	key int
}

func main() {
	var sequence = []person{
		{key: 6, name: "Francesco"},
		{key: 1, name: "Andrea"},
		{key: 5, name: "Elisa"},
		{key: 2, name: "Beatrice"},
		{key: 3, name: "Carlo"},
		{key: 4, name: "Dino"},
		{key: 7, name: "Giorgia"},
		{key: 9, name: "Irene"},
		{key: 8, name: "Henry"},
	}
	
	var sortedVector = make([]person, len(sequence))

	for i := len(sequence) - 1; i >= 0; i-- {
		var currentPerson = sequence[i]
		sortedVector[currentPerson.key - 1] = currentPerson
	}
	reverse(sortedVector)
	fmt.Println(sortedVector)
}

func reverse(slice []person) {
	var leftIndex = 0
	var rightIndex = len(slice) - 1
	for leftIndex < rightIndex {
		slice[leftIndex], slice[rightIndex] = slice[rightIndex], slice[leftIndex]
		leftIndex++
		rightIndex--
	}
}