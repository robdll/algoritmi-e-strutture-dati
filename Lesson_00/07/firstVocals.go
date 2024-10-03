// 7 Prime vocali
// Si vuole scrivere un programma che legga una sequenza di stringhe e stampi, per ogni stringa, la
// sua prima vocale (per semplicità assumiamo che le stringhe contengano solo lettere minuscole).
// Nel caso in cui una stringa non contenga vocali il programma stampa “la parola non contiene
// vocali”.

package main

import "fmt"

func main() {
	var word string
	for {
		fmt.Print("Provide a word, 'stop' to exit: ")
		fmt.Scan(&word)
		if word == "stop" {
			break
		}
		firstVocal := getFirstVocal(word)
		if firstVocal == "" {
			fmt.Println("The word does not contain any vocals")
		} else {
			fmt.Println("The first vocal is:", firstVocal)
		}
	}
}

func getFirstVocal(word string) string {
	vocals := "aeiou"
	for _, char := range word {
		if isVocal(char, vocals) {
			return string(char)
		}
	}
	return ""
}

func isVocal(currentLetter rune, vocals string) bool {
	for _, vocal := range vocals {
		if currentLetter == vocal {
			return true
		}
	}
	return false
}