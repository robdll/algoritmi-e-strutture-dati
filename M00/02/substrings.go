// Data una stringa di N caratteri nell’alfabeto a b c, stampare il numero di sottostringhe che
// iniziano con a e finiscono con b (tali sottostringhe possono sovrapporsi).
// Esempio: nella stringa ccbaacbabbcbab il numero di sottostringhe a-b è 15. Notate che
// ciascuna delle prime due a (cioè, le due più a sinistra) appaiono ciascuna in 5 sottostringhe
// a-b).

package main

import "fmt"

func main() {
	// declare a string only made with the characters a, b, c
	s := "ccbaacbabbcbab"
	// declare a counter
	countA := 0
	countSubstrings := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			countA++
		}
		if s[i] == 'b' {
			countSubstrings += countA
		}
	}
	fmt.Println(countSubstrings)
}