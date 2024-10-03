// 1. Scrivete un programma che legga una riga di caratteri terminata da a-capo, poi legga
// un’altra lettera e stampi quante volte la lettera compare nella prima riga.
// 2. Scrivete un programma che legga una riga di caratteri terminata da a-capo e che visualizzi
// un istogramma con una barra per ogni carattere dell’alfabeto, lunga quanto il numero
// delle sue occorrenze. Il programma non deve visualizzare le barre delle lettere che non
// compaiono e non deve fare distinzione fra maiuscole e minuscole.
// 3. Due parole costituiscono un anagramma se l’una si ottiene dall’altra permutando le let-
// tere (es: attore, teatro). Scrivete un programma che legga due parole e verifichi se sono anagrammi

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("1. Count letter")
	// countLetter()
	fmt.Println("2. Count letters")
	countLetters()
}

func countLetter() {
	var sequence string
	var character string
	fmt.Print("Provide a string: ")
	fmt.Scan(&sequence)
	fmt.Print("Provide a character: ")
	fmt.Scan(&character)
	counter := 0
	for _, c := range sequence {
		if string(c) == character {
			counter++;
		}
	}
	fmt.Println("The character", character, "appears", counter, "times in the string", sequence)
}

func countLetters() {
	var sequence string
	fmt.Print("Provide a string: ")
	fmt.Scan(&sequence)
	for i := 0; i < 26; i++ {
		counter := 0
		lowerChar := rune('a' + i)
		upperChar := rune('A' + i)
		for j := 0; j < len(sequence); j++ {
			currentChar := rune(sequence[j])
			if(currentChar == lowerChar || currentChar == upperChar) {
				counter++
			}
		}
		if counter != 0 {
			character := string(upperChar)
			fmt.Print(character, ": ")
			for k := 0; k < counter; k++ {
					fmt.Print("*")
			}
			fmt.Println()
		}
	}
}

func isAnagram() {
	var first string
	var second string
	fmt.Print("Provide the first word: ")
	fmt.Scan(&first)
	fmt.Print("Provide the second word: ")
	fmt.Scan(&second)
	if len(first) != len(second) {
		fmt.Println("The two words are not anagrams")
		return
	}
	for _, c := range first {
		index := strings.IndexRune(second, c)
		if index == -1 {
		}
	}
	fmt.Println("The two words are anagrams")
}
