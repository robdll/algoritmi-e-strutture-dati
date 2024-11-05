package main

import . "fmt"

func main() {

	Input := []int{2, 4, 3, 1, 5}

	Println("questo e prima: ", Input)
	Println("questo e dopo: ", QuickSort(Input))
}

func QuickSort(slice []int) []int {

	lengh := len(slice)

	if lengh == 2 {

		return slice
	} else {

		A := slice[:(lengh / 2)]
		B := slice[(lengh / 2):]


	}

}
