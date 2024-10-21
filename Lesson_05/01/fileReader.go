package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    // Open the file
    file, err := os.Open("./Lesson_05/01/input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a reader
    reader := bufio.NewReader(file)

    var lines []string

    for {
        // Read each line
        line, isPrefix, err := reader.ReadLine()
        if err != nil {
            break
        }

        // Convert the byte slice to a string
        strLine := string(line)
        lines = append(lines, strLine)

        // If the line is too long and wasn't fully read, continue reading until it's complete
        for isPrefix {
            line, isPrefix, _ = reader.ReadLine()
            strLine += string(line)
        }
    }

    // Print out the lines to verify
    fmt.Println("Lines read from the file:")
    for _, line := range lines {
        fmt.Println(line)
    }
}
