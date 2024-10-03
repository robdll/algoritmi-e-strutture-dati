// 1. Scrivere una funzione stranoProdotto(numeri []int) int che, data come parametro
// una slice di interi, trovi quelli che sono maggiori di 7 e multipli di 4 e ne restituisca il
// prodotto. Ad esempio, se la slice contiene i numeri 12, 3, 4, 8, 9, 2, la funzione dovrà
// restituire il numero 96 (pari al prodotto di 12 per 8).
// 2. Scrivere una funzione pariDispari(numeri []int) che, data come parametro una slice
// di interi, stampi, per ciascun numero, se è pari o dispari.
// 3. Scrivere una funzione minDispari(numeri []int) int che, data una slice di interi,
// restituisca il più piccolo numero dispari (la slice può contenere sia numeri positivi che
// negativi).

package main

import "fmt"

func main () {
 sequence := []int{12, 3, 4, 8, 9, 2}
 fmt.Println("stranoProdotto: Expected: 96 - Received: ", stranoProdotto(sequence))
 fmt.Println("pariDispari:")
 pariDispari(sequence)
 fmt.Println("minDispari: Expected: 3 - Received: ", minDispari(sequence))

} 

func stranoProdotto(numeri []int) int {
	prod := 1
	for _, n := range numeri {
		if n > 7 && n % 4 == 0 {
			prod *= n
		}
	}
	return prod
}

func pariDispari(numeri []int) {
	for n := range numeri {
		if n % 2 == 0 {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}
}

func minDispari(numeri []int) int {
	min := 0
	found := false
	for _, n := range numeri {
		if n % 2 != 0 && (!found || n < min) {
			min = n
			found = true
		}
	}
	return min
}
