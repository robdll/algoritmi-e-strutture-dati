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

	// Initialize board
	var board []int
	for i := 1; i<30; i++ {
		board = append(board, 0)
	}
	
	// Read board Sizefrom file
	file, err := os.Open("gameGrid.txt")
	if err != nil {
		Println("Errore nell'apertura del file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// if board is provided by user, we should check that there are no cases where 6 consecutive squares have either a snake or a stair
		input := strings.Fields(scanner.Text())
		val1, _ := strconv.Atoi(input[0])
		val2, _ := strconv.Atoi(input[1])
		board[val1-1] = val2
	}
	Println("Board: ", board)

	// get Starting point
	var startingSquare int = getStartingSquare(board)

	minThrows := calcMinDiceThrowToWin(board, startingSquare, false)
	Println(minThrows)
}

func calcMinDiceThrowToWin(board []int, startingSquare int, allowSnakeAndStair bool) int{
		
	currentSquare := startingSquare
	totThrow := 0

	for {
		if currentSquare == 29{
			break
		}
		totThrow++

		currentSquare = currentSquare + 6
		if currentSquare >= 29 {
			break
		}
		for board[currentSquare] != 0 {
			currentSquare = currentSquare - 1
		}
	}
	return totThrow
}

// 	var verify int = 0	//si parte dalla prima casella data i
// 	var throwTOT int
// 	//var dadi []int 	se vogliamo dire la serie dei dadi lanciati...
// 	//provo simulando il lancio con numero piu alto, se ciò porta ad una scala o serprente
// 	//o superi la casella 30 il lancio sara rifatto dalla casella diminuito di uno il dado
	
// 			verify = board[startingSquare]
// 			Print("Prendo ",verify)
// 			if verify != 0 {
				
// 				startingSquare -= maxThrow
// 				maxThrow -= 1
// 				startingSquare += maxThrow
				
// 				Println("if",startingSquare)
// 			}else{
// 				maxThrow = 6
// 				startingSquare += maxThrow
// 				throwTOT ++
// 				Println(" lancio 6 procedi casella: ",startingSquare)
// 			}
// 		}else if startingSquare > 29{
			
// 			startingSquare -=maxThrow
// 			lancioperf := 29 - startingSquare
// 			Print("sei andato in over, dovevi fare: ", lancioperf)
// 			break
// 		} else if startingSquare == 29 {
// 			Print("hai vinto con: ")
// 			break
// 		}
// 	}
// 	//Println(dadi)
// 	return throwTOT
// }


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

func getStartingSquare (board []int) int {
	var startingSquare int = 0
	if len(os.Args) == 2 {
		providedArg := os.Args[1]
		number, err := strconv.Atoi(providedArg)
		if err == nil && number-1 > 0 && number-1 < len(board) {
			startingSquare = number-1
			if board[startingSquare] != 0 {
				Print("La casella di partenza ", startingSquare+1, " è una scala o serpente, quindi si parte dala prima casella disponibile")
				startingSquare = 0
				for board[startingSquare] != 0 {
					startingSquare++
				}
			}
		}
	}
	return startingSquare
}