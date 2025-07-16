package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// This is a simple Go application
	fmt.Println("Hello, World!")
	fmt.Println("Welcome to my first Go application!")
	fmt.Println("This is a simple program to demonstrate Go syntax.")
	fmt.Println("Let's build something amazing with Go!")
	fmt.Println("Happy coding!")
	fmt.Println("Goodbye!")

	// Example of using the add function
	result := add(3, 4)
	fmt.Printf("The sum of 3 and 4 is: %d\n", result)

	// Example of using the swap function
	var a string
	var b string
	a, b = swap("World", "Hello")
	fmt.Printf("After swapping: a = %s, b = %s\n", a, b)
	x, y := "World", "Hello"
	x, y = swap(x, y)
	fmt.Printf("After swapping: x = %s, y = %s\n", x, y)

	c, python, java := true, false, "no!" // Multiple assignment
	fmt.Println(c, python, java)
}
