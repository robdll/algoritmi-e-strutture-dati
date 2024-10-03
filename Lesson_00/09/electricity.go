// 9 Elettricità
// Vogliamo leggere una sequenza di N interi (almeno 3), che descrivono il consumo di elettricità
// in N giorni, e stampare i giorni in cui il consumo è stato inferiore rispetto sia al giorno prima
// che al giorno dopo. I giorni sono numerati a partire da 1

package main

import "fmt"

func main() {
	var N int
	fmt.Print("Inserisci il numero di giorni: ")
	fmt.Scan(&N)
	if N < 3 {
		fmt.Println("Devi inserire almeno 3 giorni")
		return
	}
	var previous int
	var current int
	var next int
	fmt.Print("Inserisci il consumo del primo giorno: ")
	fmt.Scan(&previous)
	fmt.Print("Inserisci il consumo del secondo giorno: ")
	fmt.Scan(&current)
	for i := 3; i <= N; i++ {
		fmt.Print("Inserisci il consumo del giorno ", i, ": ")
		fmt.Scan(&next)
		if current < previous && current < next {
			fmt.Println("Il giorno", i-1, "ha un consumo inferiore rispetto sia al giorno precedente che al giorno successivo")
		}
		previous = current
		current = next
	}
}