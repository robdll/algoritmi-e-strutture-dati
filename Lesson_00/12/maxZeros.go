// Massimo numero di zeri
// Scrivere un programma maxZeros.go che legge da standard input una sequenza di interi
// terminata da -1 e stampa il numero che contiene il maggior numero di 0. Nel caso in cui vi
// siano più numeri che contengono il massimo numero di 0, il programma stampa l?ultimo in-
// contrato. Ad esempio, se la sequenza letta è 3040 145 80 1707 105002 78 1970 6005 -1
// il programma stampa 105002

package main

import "fmt"

func main() {
	var number int
	var maxZeros int
	var numberWithMaxZeros int
	for number != -1 {
		fmt.Print("Insert a number, -1 to stop: ")
		fmt.Scan(&number)
		numberString := fmt.Sprint(number)
		counter := 0
		for _, digit := range numberString {
			if digit == '0' {
				counter++
			}
		}
		if(counter > maxZeros) {
			maxZeros = counter
			numberWithMaxZeros = number
		}
	}
	fmt.Println("number with most 0 is: ", numberWithMaxZeros, " with ", maxZeros, " zeros")
}