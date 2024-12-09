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
	var i int = 0
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

	playGame(i, board)

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