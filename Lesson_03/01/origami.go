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
	coordinateMap := make(map[string]bool)
	maxColumn, maxRow := 0, 0

	// lettura file
	file, err := os.Open("./Lesson_03/01/input.txt")
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
			parts := strings.Split(line, ",")
			colonna, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			riga, _ := strconv.Atoi(parts[1])
			coordinates = append(coordinates, Coord{riga, colonna})
			// Create the key as "riga-colonna"
			key := fmt.Sprintf("%d-%d", riga, colonna)
			coordinateMap[key] = true
			if(riga > maxRow) {
				maxRow = riga
			}
			if(colonna > maxColumn) {
				maxColumn = colonna
			}
		}
	}

	// create paper and draw coordinates
	paper := make([][]string, maxRow+1)
	for i:=0; i<=maxRow; i++ {
		paper[i] = make([]string, maxColumn+1)
		for j:=0; j<=maxColumn; j++ {
			key := fmt.Sprintf("%d-%d", i, j)
			if coordinateMap[key] {
				paper[i][j] = "#"
			} else {
				paper[i][j] = "."
			}
		}
	}
	
	// esegui piegature
	for i:=0; i<len(instructions); i++ {
		isVerticalLine := instructions[i].direction == "x"
		foldingRow := len(paper)
		foldingColumn := len(paper[0])
		if isVerticalLine {
			xLine := instructions[i].index
			foldingColumn = xLine
		} else {
			yLine := instructions[i].index
			foldingRow = yLine
		}
		for j:=0; j<foldingRow; j++ {
			for k:=0; k<foldingColumn; k++ {
				var overlapValue string 
				if isVerticalLine {
					overlapValue = paper[j][maxColumn-k]
				} else {
					overlapValue = paper[maxRow-j][k]
				}
				if overlapValue == "#" {
					paper[j][k] = "#"
				}
			}
			// scarto tutte le colonne non necessarie qui
		}
		// scarto tutte le righe non necessarie qui
		if(!isVerticalLine) {
			paper = paper[0:foldingRow]
			maxRow = len(paper) -1
		} else {
			for j:=0; j<=maxRow; j++ {
				paper[j] = paper[j][0:foldingColumn]
			}
			maxColumn = len(paper[0]) -1
		}
	}


	// print the grid
	for i:=0; i<=maxRow; i++ {
		fmt.Println(strings.Join(paper[i], ""))
	}

}
