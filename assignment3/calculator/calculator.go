// Package calculator caclulates the value with the given two numbers.
package calculator

import (
	"assignment3/operations"
	"fmt"
)

// Calc calculates the value of the given two numbers.
func Calc() {

	var userNum1, userNum2, operation, result float32
	var resdiv, remainder int
	var operationType string

	fmt.Println("enter the first number: ")
	fmt.Scanln(&userNum1)

	fmt.Println("enter the second number: ")
	fmt.Scanln(&userNum2)

	fmt.Println("What operation do you want to perform?\n Type 1 for addition\n Type 2 for subtraction\n Type 3 for multiplication\n Type 4 for division")
	fmt.Scanln(&operation)

	if operation == 1 {
		result = operations.Add(userNum1, userNum2)
		operationType = "Addition"
	} else if operation == 2 {
		result = operations.Sub(userNum1, userNum2)
		operationType = "Subtraction"
	} else if operation == 3 {
		result = operations.Mul(userNum1, userNum2)
		operationType = "Multiplication"
	} else if operation == 4 {
		if userNum2 == 0 {
			fmt.Println("Not possible to divide by zero")
		} else {
			resdiv, remainder = operations.Div(userNum1, userNum2)
			operationType = "Division"
		}
	} else {
		fmt.Println("Operation not supported. Please try again.")
	}

	if operation == 1 || operation == 2 || operation == 3 {
		fmt.Println("You performed", operationType, ", the Result is: ", result)
	} else if operation == 4 && userNum2 != 0 {
		fmt.Println("You performed", operationType, "the Result is: ", resdiv, "and the Remainder is: ", remainder)
	}
}
