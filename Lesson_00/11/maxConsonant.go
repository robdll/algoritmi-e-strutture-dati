// 11 Massimo numero consonanti
// Si vuole scrivere una funzione che, presa come argomento una slice di stringhe, restituisca il nu-
// mero massimo di consonanti contenute in una stringa (assumiamo per semplicità che le stringhe
// contengano solo caratteri minuscoli)

package main

import (
	"fmt"
	"strings"
)

func main() {
	var slice = []string{"ciao", "come", "stais", "bene", "graziella"}
	fmt.Println(maxConsonant(slice))
}

func maxConsonant(slice []string) int {
	vocals := "aeiou"
	max := 0
	for _, word := range slice {
		counter := 0
		for _, letter := range word {
			if !strings.Contains(vocals, string(letter)) {
				counter++
			}
		}
		if counter > max {
			max = counter
		}
	}
	return max
}