// Raccogliere permutazioni
// Data una permutazione di 1..N, vogliamo raccogliere i numeri in ordine crescente cominciando
// ad analizzarli da sinistra. Scrivete un programma che stabilisce quante volte avremo bisogno di
// tornare verso sinistra.
// Esempio: Nella permutazione 4 5 1 3 6 2 l’output è 2, poiché 1 si trova andando sempre
// verso destra, poi si prosegue verso destra per raccogliere 2, ma per raccogliere 3 bisogna tornare
// indietro verso sinistra; bisogna tornare ancora indietro per raccogliere 4, dopodiché 5 e 6 si
// trovano in ordine proseguendo verso destra.

package main

import "fmt"

func main() {
	var permutation = []int{7, 4, 5, 1, 3, 6, 2}
	var counter = 0
	var current = 0
	for j := 0; current < len(permutation); j++ {
		counter++
		if(j%2 == 0) {
			for i := 0; i < len(permutation); i++ {
				if(permutation[i] == current + 1) {
					current++
				}
			}
		} else {
			for i := len(permutation) - 1; i >= 0; i-- {
				if(permutation[i] == current + 1) {
					current++
				}
			}
		}
	}
	fmt.Println(counter)
}