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
	maxX, maxY := 0, 0
	emptyLine := 0
	for i:=0; instructions[i] != ""; i++ {
		coords := strings.Split(instructions[i], ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		maxX = max(maxX, x)
		maxY = max(maxY, y)
		emptyLine++
		// Improvement 1: we could create a map to keep track of all the # position (e.g. { 5: [1, 3, 5] })
	}
	emptyLine++

	// create paper
	paper := make([][]string, maxY+1)
	for i:=0; i<=maxY; i++ {
		paper[i] = make([]string, maxX+1)
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
	for i:=0; i<=maxY; i++ {
		for j:=0; j<=maxX; j++ {
			if paper[i][j] == "" {
				paper[i][j] = "."
			}
		}
	}

	// get folding instructions
	// foldingInstruction = instructions[len(instructions)-1]

	// print the map
	for i:=0; i<=maxY; i++ {
		fmt.Println(strings.Join(paper[i], ""))
	}




}