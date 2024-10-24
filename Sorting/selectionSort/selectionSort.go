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
	for idx := 0; idx < len(slice)-1; idx++ {
		minIdx := idx
		// find the minimum element in the remaining unsorted slice
		for j := idx + 1; j < len(slice); j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		// move minimum element to the index of the unsorted slice
		slice[idx], slice[minIdx] = slice[minIdx], slice[idx]
	}
}