package main

import "fmt"

// Define a function that takes two integers and returns an integer
func add(a int, b int) int {
    return a + b
}

func main() {
    // Assign the function 'add' to a variable
    var myFunc func(int, int) int
    myFunc = add

    // Now, you can use 'myFunc' to call the function
    result := myFunc(3, 4)
    fmt.Println(result) // Output: 7
}