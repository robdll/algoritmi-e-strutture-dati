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

type person struct {
	name string
	key int
}

func main() {
	var sequence = []person{
		{key: 1, name: "Andrea"},
		{key: 2, name: "Beatrice"},
		{key: 6, name: "Francesco"},
		{key: 5, name: "Elisa"},
		{key: 8, name: "Henry"},
		{key: 9, name: "Irene"},
		{key: 4, name: "Dino"},
		{key: 3, name: "Carlo"},
		{key: 7, name: "Giorgia"},
}

	var requiredSwap = len(sequence) -1
	for numSwap, unorderedLeftIndex := 0, 0; numSwap != requiredSwap; {
		if sequence[unorderedLeftIndex].key != len(sequence) - unorderedLeftIndex {
			sequence[unorderedLeftIndex], sequence[len(sequence)-sequence[unorderedLeftIndex].key] = sequence[len(sequence)-sequence[unorderedLeftIndex].key], sequence[unorderedLeftIndex]
			numSwap++
		} else {
			unorderedLeftIndex++
			requiredSwap--
		}
	}

	for _, p := range sequence {
		println(p.name, p.key)
	}

}