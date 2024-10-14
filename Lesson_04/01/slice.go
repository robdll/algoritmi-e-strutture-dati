package main

import "fmt"

func main() {
	s0 := []int{0, 0} // definisce uno slice di interi con due elementi 0, 0
	s1 := append(s0, 2) // definisce uno slice di interi con tre elementi 0, 0, 2
	s1[0] = 100 // modifica il primo elemento dello slice s1 100, 0, 2, ma non quello di s0 poiché la capacità di s1 è maggiore di quella di s0
	s2 := append(s1, 3, 5, 7) // definisce uno slice di interi con sei elementi 100, 0, 2, 3, 5, 7
	s3 := append(s2, s0...) // definisce uno slice di interi con otto elementi 100, 0, 2, 3, 5, 7, 0, 0
	s3[2] = 99 // modifica il terzo elemento dello slice s3 100, 0, 2, 99, 5, 7, 0, 0
	s4 := append(s3[3:6], s3[2:]...) // definisce uno slice di interi con cinque elementi 99, 5, 7, 0, 0
	fmt.Println("s0: ", s0)
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("s3: ", s3)
	fmt.Println("s4: ", s4)
}

