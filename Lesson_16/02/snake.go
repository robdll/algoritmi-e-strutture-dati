package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

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

	minRolls := calcMinDiceRollToWin(board, startingSquare)
	Println("Numero minimo di lanci per vincere: ", minRolls)
}

func calcMinDiceRollToWin(board []int, startingSquare int) int {
	rolls := 0

	// creo un array di len(board) elementi

	// creo una coda con la square di partenza

	// finchè la coda non è vuota o ho trovato la square 30

		// prendo il primo elemento della coda

		// se il valore della square è 30, ho vinto e ritorno il numero di lanci

		// se la square è già stata visitata, continuo con il prossimo elemento (break)

		// segno la square come visitata

		// metto in coda tutti i vicini di quella square (square+1..6) (se square)

		// incremento il numero di lanci

		return rolls

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