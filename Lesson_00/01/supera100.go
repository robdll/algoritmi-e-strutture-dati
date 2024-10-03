// 1 Supera 100
// Scrivere un programma supera100.go che legge da standard input una sequenza di interi po-
// sitivi terminata da -1 e stampa il primo numero che supera 100, se presente; altrimenti stampa
// “nessun numero maggiore di 100”

package main

import "fmt"

func main() {
	var number = 0
	var found = false
	for number != -1 {
		fmt.Print("Provide a number: ")
		fmt.Scan(&number)
		if number > 100 && !found {
			fmt.Print("Found a number above 100: ", number)
			found = true
		}
	}
	if !found {
		fmt.Println("nessun numero maggiore di 100")
	}
}