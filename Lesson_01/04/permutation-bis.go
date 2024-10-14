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
	var sequence = []int{7, 4, 5, 1, 3, 6, 2}
	readValues := make(map[int]bool)
	counter := 0
	for _, value := range sequence {
		readValues[value] = true
		if(readValues[value+1]) {
			counter++
		}
	}
	fmt.Println(counter)
}