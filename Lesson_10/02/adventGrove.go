package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	// read input.txt file line by line
	// and print each line
	file, err := os.Open("./Lesson_10/02/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan()  {
		line := scanner.Text()
		fmt.Println(line)
	}
	file.Close()
}
