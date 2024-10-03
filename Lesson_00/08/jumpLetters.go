// 8 Inizia con ’a’, inizia con ’b’
// Si vuole scrivere un programma che legga una sequenza di dieci stringhe e stampi il numero di
// stringhe che cominciano con la lettera a e il numero di stringhe che cominciano con b.

package main

import "fmt"

func main() {
	var words = make([]string, 10)
	var counter int = 0
	for i := 0; i<10; i++ {
		fmt.Print("Provide word number ", i+1, ":")
		fmt.Scan(&words[i])
		//check if the word starts with a or b
		if(words[i][0] == 'a' || words[i][0] == 'b') {
			counter++
		}
	}
	fmt.Println("The number of words starting with 'a' or 'b' (lowercase) is: ", counter)
}