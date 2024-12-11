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
	var i int = 0 //casella i + 1
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
	Println("inserisci la casella di partenza: ")
	for{
		Scan(&i)
		if board[i] != 0{
			Println("posizione non valida, inserisci un'altra casella: ")
		}else{
			break
		}

	}
	//playGame(i, board)
	mosseMinime := MosseMinimeSenzaUsoScaleSerpenti(board, i+1)
	Println(mosseMinime)


	// casella di paretenza i + 1 da fare dopo
}

func MosseMinimeSenzaUsoScaleSerpenti(board [30]int, startbox int) int{
		
	maxThrow:= 6
	var verify int = 0	//si parte dalla prima casella data i
	var throwTOT int
	//var dadi []int 	se vogliamo dire la serie dei dadi lanciati...
	//provo simulando il lancio con numero piu alto, se ciò porta ad una scala o serprente
	//o superi la casella 30 il lancio sara rifatto dalla casella diminuito di uno il dado
	for {
		
		if startbox < 29{
			
			verify = board[startbox]
			Print("Prendo ",verify)
			if verify != 0 {
				
				startbox -= maxThrow
				maxThrow -= 1
				startbox += maxThrow
				
				Println("if",startbox)
			}else{
				maxThrow = 6
				startbox += maxThrow
				throwTOT ++
				Println(" lancio 6 procedi casella: ",startbox)
			}
		}else if startbox > 29{
			
			startbox -=maxThrow
			lancioperf := 29 - startbox
			Print("sei andato in over, dovevi fare: ", lancioperf)
			break
		}else if startbox == 29 {
			Print("hai vinto con: ")
			break
		}
	}
	//Println(dadi)
	return throwTOT
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