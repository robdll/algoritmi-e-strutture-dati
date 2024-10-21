package main

import "fmt"

func main() {
	fmt.Println(multiply(5, 3))
}

// multiply recursively
func multiply(x, y int) int {
	if y == 0 {
		return 0
	}
	toAdd := multiply(x, y-1)
	return x + toAdd
}
