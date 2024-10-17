package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var instructions = []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
		"",
		"fold along y=7",
		"fold along x=5",
	}

	// calc paper size
	maxColumn, maxRow := 0, 0
	coordsSection  := true
	foldingSequence := make([]string, 0)

	for i:=0; i < len(instructions); i++ {
		if (instructions[i] == "") {
			coordsSection = false
			continue
		}
		if(coordsSection) {
			coords := strings.Split(instructions[i], ",")
			a, _ := strconv.Atoi(coords[0])
			b, _ := strconv.Atoi(coords[1])
			maxColumn = max(maxColumn, a)
			maxRow = max(maxRow, b)
		} else {
			foldString := strings.Split(instructions[i], " ")[2]
			foldingSequence = append(foldingSequence, foldString)
		}
	}

	// create paper
	paper := make([][]string, maxRow+1)
	for i:=0; i<=maxRow; i++ {
		paper[i] = make([]string, maxColumn+1)
	}

	// draw # symbols
	for i:=0; instructions[i] != ""; i++ {
		coords := strings.Split(instructions[i], ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		paper[y][x] = "#"
	}

	// draw a . symbols
	for i:=0; i<=maxRow; i++ {
		for j:=0; j<=maxColumn; j++ {
			if paper[i][j] == "" {
				paper[i][j] = "."
			}
		}
	}

	// per ogni istruzione di piegatura
	for i:=0; i<len(foldingSequence); i++ {
		// per ogni riga della carta o riga/2 se pieghi lungo y
		// check first char of foldingSequence to see if it's x or y

		isVerticalFold := strings.Split(foldingSequence[i], "=")[0] == "y"
		fmt.Println("Piegatura",isVerticalFold, foldingSequence[i])
		xFoldingPoint := maxColumn
		yFoldingPoint := maxRow
		if isVerticalFold {
			xFoldingPoint = (maxColumn-1)/2
			fmt.Println("E' un taglio verticale, quindi lavoriamo su tutte le righe e su #colonne:",xFoldingPoint)
		} else {
			yFoldingPoint = (maxRow-1)/2
			fmt.Println("E' un taglio orizzontale, quindi lavoriamo su tutte le colonne e su #righe:",yFoldingPoint)
		}
		fmt.Println("foldX",xFoldingPoint,"foldY", yFoldingPoint)
		for j:=0; j<xFoldingPoint; j++ {
			fmt.Println("riga", j)
			for k:=0; k<xFoldingPoint; k++ {
				fmt.Println("colonna", k)
			}
		}
	}


	// print the map
	for i:=0; i<=maxRow; i++ {
		fmt.Println(strings.Join(paper[i], ""))
		// riga lunghezza meno i
		// fmt.Println(strings.Join(paper[maxRow-i], ""))
	}
	fmt.Println(foldingSequence)

	
}

// avrei fatto un ciclo for iniziale per vedere se vuole che tagli alla y o x fino a che boh prema un tasto che vuole per uscire dalla piegatura, poi if che appunto
// mi dice se ha scelto un x o y, e dentro ognuno facevi un for per appunto fare il fold facendo poi vedere allutente la piegatura corrente e poi di nuovo la domanda
// almeno cosi fai il ciclo for per y o per x