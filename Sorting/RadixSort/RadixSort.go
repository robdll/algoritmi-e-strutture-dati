package main

import . "fmt"

//isolare le cifre

func main() {

	array := []int{124, 432, 876, 954}
	slice := []int{}
	for _, v := range array {
		var cifra int
		for v > 0 {
			cifra = v % 10
			v = v / 10
			slice = append(slice, cifra)
		}
	}

	Print(slice)
}
