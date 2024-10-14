package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Simple Calculator")

	var num1 float64
	var num2 float64
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Print("1. Add\n2. Subtract\n3. Multiply\n4. Divide\n")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println(add(num1, num2))
	case 2:
		fmt.Println(sub(num1, num2))
	case 3:
		fmt.Println(mul(num1, num2))
	case 4:
		result, err := div(num1, num2)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		fmt.Println(result)
	default:
		fmt.Println("Invalid choice")
	}
}

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
