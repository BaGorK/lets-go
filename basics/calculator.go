package main

import "fmt"

func calc() {
	var operation string
	var number1, number2 float64

	fmt.Println("===============================")
	fmt.Println("CALCULATOR GO APPLICATION 1.0.0")
	fmt.Println("===============================")
	fmt.Println("Which operation would you like to perform? (add, subtract, multiply, divide)")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter the first number:")
	fmt.Scanf("%f", &number1)
	fmt.Println("Enter the second number:")
	fmt.Scanf("%f", &number2)

	switch operation {
	case "add":
		result := number1 + number2
		fmt.Printf("The result of adding %.2f and %.2f is: %.2f\n", number1, number2, result)
	case "subtract":
		result := number1 - number2
		fmt.Printf("The result of subtracting %.2f from %.2f is: %.2f\n", number2, number1, result)
	case "multiply":
		result := number1 * number2
		fmt.Printf("The result of multiplying %.2f and %.2f is: %.2f\n", number1, number2, result)
	case "divide":
		if number2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		} else {
			result := number1 / number2
			fmt.Printf("The result of dividing %.2f by %.2f is: %.2f\n", number1, number2, result)
		}
	default:
		fmt.Println("Error: Invalid operation. Please choose from add, subtract, multiply, or divide.")
	}
	// The %.2f is a format verb for printing floating-point numbers with exactly 2 digits after the decimal point.
}
