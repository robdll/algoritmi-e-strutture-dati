// implement insertion sort

package main

import "fmt"

func main() {
	/*
	*  Using slice
	*/
	sequence := []int{}
	var number int
	fmt.Print("Insert a number, 0 to stop: ")
	fmt.Scan(&number)
	for number != 0 {
		sequence = updateSequence(sequence, number)
		fmt.Println("New Sequence: ", sequence)
		fmt.Println("Insert a number, 0 to stop: ")
		fmt.Scan(&number)
	}

}

func updateSequence(slice []int, number int) []int {
	if(len(slice) == 0) {
		slice = append(slice, number)
	} else {
		// get slice length 
		index := 0
		for ; index < len(slice) && slice[index] < number ; index ++ {}
		rightPart := append([]int{number}, slice[index:]...)
		slice = append(slice[:index], rightPart...)
	}
	fmt.Println("Slice Cap ", cap(slice))
	return slice
}

