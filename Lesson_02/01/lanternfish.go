package main

import "fmt"

func main() {
	sequence := []int{3,4,3,1,2}
	simulationDays := 80

	fmt.Println("Initial Sequence", sequence)

	for i := 1; i < simulationDays; i++ {
		newBorn := []int{}
		for j := 0; j < len(sequence); j++ {
			sequence[j]--
			if sequence[j] == -1 {
				newBorn = append(newBorn, 8)
				sequence[j] = 6
			}
		}
		sequence = append(sequence, newBorn...)
		fmt.Println("After", i, "days:")
		fmt.Println(sequence)
	}
	fmt.Println("After", simulationDays, "days there are", len(sequence), "lanternfish")
}