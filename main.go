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

type Student struct {
	name   string
	gender string
	age    int
}

func setStudet(name string, gender string, age int) Student {
	// st := Student{name, gender, age}
	st := Student{name, gender, age} // This is a struct literal, which is a way to create a new struct with field names.
	// or st := Student{}
	// st.name = name
	// st.gender = gender
	// st.age = age

	return st
}

func printStudent(s Student) {
	fmt.Println(s)
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
	// calc()

	// Types
	// Pointers
	// i, j := 42, 2701
	//
	// p := &i         // point to i
	// fmt.Println(*p) // read i through the pointer
	// *p = 21         // set i through the pointer
	// fmt.Println(i)  // see the new value of i
	//
	// p = &j         // point to j
	// *p = *p / 37   // divide j through the pointer
	// fmt.Println(j) // see the new value of j
	i, j := 42, 2701
	p := &i                                        // point to i
	fmt.Println("Value of i through pointer:", *p) // read i through the pointer
	*p = *p - 20
	fmt.Println("New value of i after modification:", i) // see the new value of i
	p = &j
	fmt.Println("Value of j through pointer:", *p) // read j through the pointer

	// Structs
	// A struct is a collection of fields.
	type Vertex struct {
		x int
		y int
	}
	v := Vertex{1, 2}         // Create a new Vertex
	v.x = 4                   // Set the x field
	fmt.Println("Vertex:", v) // Print the Vertex

	// Arrays
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println(arr[0], arr[1]) // Print the array elements
	fmt.Println(arr)            // Print the entire Arrays

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)                                 // Print the array of primes
	fmt.Println("Length of primes array:", len(primes)) // Print the length of the primes array

	/*
		Slices

		An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

		The type []T is a slice with elements of type T.

		Slices are like references to arrays

		A slice does not store any data, it just describes a section of an underlying array.

		Changing the elements of a slice modifies the corresponding elements of its underlying array.

		Other slices that share the same underlying array will see those changes.
	*/
	var s []int
	if len(primes) > 0 {
		s = primes[0:4]
	} else {
		s = []int{} // Initialize an empty slice if primes is empty
	}
	fmt.Println("Slice of primes:", s) // Print the slice of primes

	names := [4]string{"Name 1", "Name 2", "Name 3", "Name 4"}
	fmt.Println("Names array:", names) // Print the names array

	names1 := names[0:2] // Name 1 and Name 2
	names2 := names[1:3] // Name 2 and Name 3

	fmt.Println("Slice names slice 1:", names1) // Print the slice a
	fmt.Println("Slice names slice 2:", names2) // Print the slice b

	names2[0] = "New Name 2"                                       // Modify the first element of names2
	fmt.Println("Modified Names 2 slice:", names2)                 // Print the modified slice b
	fmt.Println("Original names array after modification:", names) // Print the original names array to see the effect of the modification

	names[0] = "New Name 1"                                        // Modify the first element of names
	fmt.Println("Modified Names array:", names)                    // Print the modified names array
	fmt.Println("Slice names slice 1 after modification:", names1) // Print the slice a to see the effect of the modification

	// Array of Structs
	var students []Student
	students = append(students, setStudet("st 1", "male", 23))
	students = append(students, setStudet("st 2", "female", 19))
	students = append(students, Student{name: "st 3", gender: "male", age: 24}) // Using struct literal to create a new Student
	fmt.Println("Students slice:", students)                                    // Print the students slice
}
