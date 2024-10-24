// implement selection sort

package main

import "fmt"

func main() {
	s := []int{15, 96, 44, 22, 54, 28, 83}
	fmt.Println("Before: ", s)
	selectionSort(s)
	fmt.Println("After: ", s)
}

func selectionSort(slice []int) {
	// iterate n-1 times
	for idx := len(slice)-1; idx > 0; idx-- {
		maxIdx := idx
		// find the minimum element in the remaining unsorted slice
		for j := idx - 1; j >= 0; j-- {
			if slice[j] > slice[maxIdx] {
				maxIdx = j
			}
		}
		// move minimum element to the index of the unsorted slice
		slice[idx], slice[maxIdx] = slice[maxIdx], slice[idx]
	}
}