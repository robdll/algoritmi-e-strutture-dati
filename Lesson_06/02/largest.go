package main

import "fmt"

func main() {
	fmt.Println(largest([]int{1, 7, 2, 3, 4, 5}))
}

func largest(numbers []int) int {
	n := len(numbers)
	if n == 1 {
		return numbers[0]
	}
	bigger := max( numbers[n-1], largest( numbers[:n-1] ) )
	return bigger
}
