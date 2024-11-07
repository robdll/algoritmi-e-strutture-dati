package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main () {

	planetTable := make(map[string]string)

	// lettura input
	file, err := os.Open("./Lesson_09/02/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan()  {
		// read line and separate parent from child
		line := scanner.Text()
		parts := strings.Split(line, ")")
		parent := parts[0]
		satellite := parts[1]
		planetTable[satellite] = parent
	}


	totOrbits := 0
	for planetName := range planetTable {
		planetOrbits := planetDept(planetTable, planetName)
		totOrbits += planetOrbits
	}
	fmt.Println(totOrbits)

}

func planetDept(planetTable map[string]string, planetName string) int {
	if planetName == "COM" {
		return 0
	}
	return 1 + planetDept(planetTable, planetTable[planetName])
}