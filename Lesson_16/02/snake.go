package main

import (
	"bufio"
	. "fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func throwDice() int {
	return rand.Intn(6)+1
}

func main(){

	file, err := os.Open("gameGrid.txt")
	if err != nil {
		Println("Errore nell'apertura del file:", err)
		return
	}
	defer file.Close()
	var board [30]int
	//var i int = 0
	for i := 1; i<30; i++ {
		board[i] = 0
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := strings.Fields(scanner.Text())
		val1, _ := strconv.Atoi(input[0])
		val2, _ := strconv.Atoi(input[1])
		board[val1-1] = val2
	}

	// Println("numero random 1: ", throwDice())
	Println("board", board)

	//playGame(i, board)
	mosseMinime := MosseMinimeSenzaUsoScaleSerpenti(board)
	Println(mosseMinime)

}

func MosseMinimeSenzaUsoScaleSerpenti(board [30]int) int{
		
	lancioMax:= 6
	var casellacorrente,verify int	//si parte dalla prima casella
	var lanci int
	//provo simulando il lancio con numero piu alto, se ciò porta ad una scala o serprente
	//o superi la casella 30 il lancio sara rifatto dalla casella diminuito di uno il dado
	for {
		
		if casellacorrente < 29{
			
			verify = board[casellacorrente]
			Print("Prendo ",verify)
			if verify != 0 {
				
				casellacorrente -= lancioMax
				lancioMax -= 1
				casellacorrente += lancioMax
				lanci --
				Println("if",casellacorrente)
			}else{
				lancioMax = 6
				casellacorrente += lancioMax
				lanci ++
				Println(" lancio 6 procedi casella: ",casellacorrente)
			}
		}else if casellacorrente > 29{
			Print("sei andato in over")
			break
		}else if casellacorrente == 29 {
			Print("hai vinto con: ")
			break
		}
	}

	return lanci
}


func playGame(i int, board [30]int) {
	for {
		Print("lancio da casella ",i, " ")
		numDice:= throwDice()
		i+= numDice
		if i == 30{
			Println("Hai vinto! con il lancio: ", numDice)
			break
		}else if i > 30 {
			Println("sei andato oltre le 30 caselle con il lancio: ", numDice)	
			break
		}
		val:=board[i]
		if val != 0{
			i = val
		}
		Println("a casella", i, "con il lancio dado numero: ", numDice)
	}
}