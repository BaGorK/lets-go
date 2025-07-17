package main

import (
	"fmt"
	"runtime"
	"time"
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

	// For loop example
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += 1
	}
	fmt.Println("The sum of numbers from 0 to 100 is:", sum)
	// While loop example
	whileSum := 0
	for whileSum < 10 {
		whileSum += 1
	}
	fmt.Println("The while loop sum is:", whileSum)

	// If statement example
	if sum > 50 {
		fmt.Println("The sum is greater than 50")
	} else if sum == 50 {
		fmt.Println("The sum is equal to 50")
	} else {
		fmt.Println("The sum is less than 50")
	}

	// If statement with short variable declaration
	if os := runtime.GOOS; os == "linux" {
		fmt.Println("Running on Linux")
	} else {
		fmt.Printf("Not running on Linux, running on %s\n", os)
	}

	// Switch statement example
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s.\n", os)
	}

	// Switch with no condition
	// Switch without a condition is the same as switch true.
	// This construct can be a clean way to write long if-then-else chains.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon.")
	case t.Hour() < 17:
		fmt.Println("It's afternoon.")
	default:
		fmt.Println("It's evening.")
	}

	// time
	fmt.Println(time.Now().Year())    // 2025
	fmt.Println(time.Now().Month())   // July
	fmt.Println(time.Now().Weekday()) // Wednesday
	fmt.Println(time.Now().Hour())    // 19
	fmt.Println(time.Now().Minute())  // 30
	fmt.Println(time.Now().Second())  // 17
	fmt.Println(time.Now().Clock())   // 19 30 17
	fmt.Println(time.Now().Local())   // 2025-07-16 19:30:17.156561984 +0300 EAT
	fmt.Println(time.Now().Date())    // 2025 July 16
	fmt.Println(time.TimeOnly)        // 15:04:05

	// calculator application
	calc()
}
