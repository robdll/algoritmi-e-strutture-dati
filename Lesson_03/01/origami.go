package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Instruction struct {
	direction string
	index int
}

func main() {

	var instructions []Instruction
	var coordinates []Coord

	file, err := os.Open("./Lesson_03/01/input_small.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan()  {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.Contains(line, "f") {
			parts := strings.Split(line, "=")
			letter := strings.Split(parts[0], " ")[2]
			index, _ := strconv.Atoi(parts[1])
			instructions = append(instructions, Instruction{letter, index})
		} else {
			fmt.Println("line", line)
			parts := strings.Split(line, ",")
			riga, _ := strconv.Atoi(parts[0])
			colonna, _ := strconv.Atoi(parts[1])
			coordinates = append(coordinates, Coord{riga, colonna})
		}
	}

	fmt.Println("instructions", instructions)
	fmt.Println("coordinates", coordinates)


	// calc paper size
	// maxColumn, maxRow := 0, 0
	// coordsSection  := true
	// foldingSequence := make([]string, 0)

	// // for _, line := range lines {
	// for i:=0; i < len(instructions); i++ {
	// 	// if first char is a letter

	// 	firstChar := instructions[i][0]
	// 	if firstChar == 'f'{
	// 		coordsSection = false
	// 	}
	// 	if(coordsSection) {
	// 		coords := strings.Split(instructions[i], ",")
	// 		a, _ := strconv.Atoi(coords[0])
	// 		b, _ := strconv.Atoi(coords[1])
	// 		maxColumn = max(maxColumn, a)
	// 		maxRow = max(maxRow, b)
	// 	} else {
	// 		foldString := strings.Split(instructions[i], " ")[2]
	// 		foldingSequence = append(foldingSequence, foldString)
	// 	}
	// }

	// // create paper
	// paper := make([][]string, maxRow+1)
	// for i:=0; i<=maxRow; i++ {
	// 	paper[i] = make([]string, maxColumn+1)
	// }

	// // draw # symbols
	// for i:=0; instructions[i] != ""; i++ {
	// 	coords := strings.Split(instructions[i], ",")
	// 	x, _ := strconv.Atoi(coords[0])
	// 	y, _ := strconv.Atoi(coords[1])
	// 	paper[y][x] = "#"
	// }

	// // draw a . symbols
	// for i:=0; i<=maxRow; i++ {
	// 	for j:=0; j<=maxColumn; j++ {
	// 		if paper[i][j] == "" {
	// 			paper[i][j] = "."
	// 		}
	// 	}
	// }

	// for i:=0; i<=maxRow; i++ {
	// 	fmt.Println(strings.Join(paper[i], ""))
	// }
	// fmt.Println(foldingSequence)

	// // per ogni istruzione di piegatura
	// for i:=0; i<len(foldingSequence); i++ {
	// 	// per ogni riga della carta o riga/2 se pieghi lungo y
	// 	// check first char of foldingSequence to see if it's x or y

	// 	isVerticalLine := strings.Split(foldingSequence[i], "=")[0] == "x"
	// 	foldingRow := len(paper)
	// 	foldingColumn := len(paper[0])
	// 	if isVerticalLine {
	// 		xLine, _ := strconv.Atoi(strings.Split(foldingSequence[i], "=")[1])
	// 		foldingColumn = xLine
	// 	} else {
	// 		yLine, _ := strconv.Atoi(strings.Split(foldingSequence[i], "=")[1])
	// 		foldingRow = yLine
	// 	}
	// 	for j:=0; j<foldingRow; j++ {
	// 		for k:=0; k<foldingColumn; k++ {
	// 			var overlapValue string 
	// 			if isVerticalLine {
	// 				// fmt.Println("foldingColumn", foldingColumn)
	// 				overlapValue = paper[j][maxColumn-k]
	// 			} else {
	// 				overlapValue = paper[maxRow-j][k]
	// 			}
	// 			if overlapValue == "#" {
	// 				paper[j][k] = "#"
	// 			}
	// 		}
	// 		// scarto tutte le colonne non necessarie qui
			
	// 	}
	// 	// scarto tutte le righe non necessarie qui
	// 	if(!isVerticalLine) {
	// 		paper = paper[0:foldingRow]
	// 		maxRow = len(paper) -1
	// 	} else {
	// 		for j:=0; j<=maxRow; j++ {
	// 			paper[j] = paper[j][0:foldingColumn]
	// 		}
	// 		maxColumn = len(paper[0]) -1
	// 	}
	// }


	// // print the map
	// for i:=0; i<=maxRow; i++ {
	// 	fmt.Println(strings.Join(paper[i], ""))
	// 	// riga lunghezza meno i
	// 	// fmt.Println(strings.Join(paper[maxRow-i], ""))
	// }

	
}

// avrei fatto un ciclo for iniziale per vedere se vuole che tagli alla y o x fino a che boh prema un tasto che vuole per uscire dalla piegatura, poi if che appunto
// mi dice se ha scelto un x o y, e dentro ognuno facevi un for per appunto fare il fold facendo poi vedere allutente la piegatura corrente e poi di nuovo la domanda
// almeno cosi fai il ciclo for per y o per x