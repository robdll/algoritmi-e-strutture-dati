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
		// Improvement 1: we could create a map to keep track of all the # position (e.g. { 5: [1, 3, 5] })
	}

	// create paper
	paper := make([][]string, maxRow+1)
	for i:=0; i<=maxRow; i++ {
		paper[i] = make([]string, maxColumn+1)
	}

	// draw #
		// Improvement 1 (use the map here
	for i:=0; instructions[i] != ""; i++ {
		coords := strings.Split(instructions[i], ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		paper[y][x] = "#"
	}

	// draw a . where there is no #
	// Improvement 2 skip the previous loop and just either draw a . or a # in the next loop with the map
	for i:=0; i<=maxRow; i++ {
		for j:=0; j<=maxColumn; j++ {
			if paper[i][j] == "" {
				paper[i][j] = "."
			}
		}
	}

	// get folding instructions
	// foldingInstruction = instructions[len(instructions)-1]

	// print the map
	for i:=0; i<=maxRow; i++ {
		fmt.Println(strings.Join(paper[i], ""))
		// riga lunghezza meno i
		// fmt.Println(strings.Join(paper[maxRow-i], ""))
	}
	fmt.Println(foldingSequence)

}