// implement selection sort

package main

import "fmt"

func main() {
	s := []int{15, 96, 44, 22, 54, 28, 83}
	fmt.Println("Before: ", s)
	
	fmt.Println("After: ", selectionSort(s))
}

func selectionSort(slice []int) []int {
	fmt.Println("selectionSort: ", slice)
	if len(slice) == 1 {
		return slice
	} else {
		// find the minimum element in the slice
		maxIdx := 0
		for i := 0; i < len(slice); i++ {
			if slice[i] > slice[maxIdx] {
				maxIdx = i
			}
		}
		// move the minimum element to the first index
		slice[len(slice)-1], slice[maxIdx] = slice[maxIdx], slice[len(slice)-1]
		// recursively sort the remaining slice
		return append(selectionSort(slice[:len(slice)-1]), slice[len(slice)-1])
	}
}