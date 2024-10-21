package main

import "fmt"

func main() {
	fmt.Println(f_rec(7))
}

// fibonacci recursively
func f_rec(n int) uint64 {
	if n == 1 || n == 2 {
		return 1
	}
	return f_rec(n-1) + f_rec(n-2)
}
