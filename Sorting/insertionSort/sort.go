// implement insertion sort

package main

import "fmt"

func main() {
	s := []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Before: ", s)
	insertSort(s)
	fmt.Println("After: ", s)
}

func insertSort(slice []int) {
	// iterate n-1 times
 	for currentIdx:=1; currentIdx<len(slice); currentIdx++ {
		// store current element
		currentElement := slice[currentIdx]
		fmt.Println("Iteration #", currentIdx, "currentElement: ", currentElement)
		// iterate on each element on the left (of our current element) until they are bigger than the current one
		j := currentIdx-1
		for j>=0 && slice[j]>currentElement {
			fmt.Println("Element ", slice[j], "is moved to the right")
			// move the element to the right
			slice[j+1] = slice[j]
			j--
		}
		// insert the current element in the right position
		fmt.Println(currentElement, "goes to position", j+1 )
		slice[j+1] = currentElement
 	}
}