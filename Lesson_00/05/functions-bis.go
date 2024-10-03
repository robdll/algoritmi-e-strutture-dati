// 1. Scrivere una funzione quanteConA(parole []string) int che, data una slice di stringhe,
// restituisca quante stringhe contengono il carattere ‘a’.
// 2. Scrivere una funzione primaConA(parole []string) string che, data una slice di
// stringhe, restituisca la prima parola che contentiene il carattere ‘a’, o la stringa vuota.
// 3. Scrivere una funzione piuCorta(parole []string) int che, data una slice di stringhe,
// restituisca la lunghezza della stringa più corta in termini di byte.

package main

import "fmt"

func main () {
	words := []string{"ciao", "foo", "stai", "bene", "grazie"}
	fmt.Println("quanteConA: Expected: 3 - Received: ", quanteConA(words))
	fmt.Println("primaConA: Expected: ciao - Received: ", primaConA(words))
	fmt.Println("piuCorta: Expected: 3 - Received: ", piuCorta(words))
}

func quanteConA(parole []string) int {
	count := 0
	for _, p := range parole {
		for _, c := range p {
			if c == 'a' {
				count++
				break
			}
		}
	}
	return count
}

func primaConA(parole []string) string {
	for _, p := range parole {
		for _, c := range p {
			if c == 'a' {
				return p
			}
		}
	}
	return ""
}

func piuCorta(parole []string) int {
	min := len(parole[0])
	for _, p := range parole {
		if len(p) < min {
			min = len(p)
		}
	}
	return min
}