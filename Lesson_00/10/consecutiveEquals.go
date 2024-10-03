// 10 Coppie di numeri uguali
// Scrivere un programma che legga da standard input una sequenza di numeri terminata da zero e
// conti quante sono le coppie di numeri naturali consecutivi uguali.

package main

import "fmt"

func main() {
	var previous int
	var current int
	var count int
	fmt.Print("Provide a number: ")
	fmt.Scan(&current)
	previous = current
	for current != 0 {
		fmt.Print("Provide another number, 0 to stop: ")
		fmt.Scan(&current)
		if current == previous {
			count++
		}
		previous = current
	}
	fmt.Println("Number of consecutive equal numbers: ", count)
}