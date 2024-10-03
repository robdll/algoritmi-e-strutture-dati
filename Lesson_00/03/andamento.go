// 3 Andamento
// Data da standard input una serie di interi positivi terminata da 0, stampare ‘+’ ogni volta che il
// nuovo valore è maggiore o uguale al precedente e ‘-’ altrimenti.

package main

import "fmt"

func main() {
	var previous int
	var current int
	fmt.Print("Provide a number: ")
	fmt.Scan(&current)
	previous = current
	for current != 0 {
		fmt.Print("Provide a number: ")
		fmt.Scan(&current)
		if current >= previous {
			fmt.Println("+")
		} else {
			fmt.Println("-")
		}
		previous = current
	}
	
}