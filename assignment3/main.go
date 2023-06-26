// Package main is executable package for this project.
package main

import (
	"assignment3/calculator"
	"fmt"
	"strings"
)

func main() {

	isCalculatorOn := true
	var calUse string

	fmt.Println("Welcome to the assignment 3 calculator!")

	for isCalculatorOn {

		calculator.Calc()

		fmt.Println("Do you want to continue using the calculator (Y/N)")
		fmt.Scanln(&calUse)
		calUse = strings.ToUpper(calUse)

		if calUse == "Y" {
			fmt.Println("Please continue using the calculator")
		} else {
			fmt.Println("Exiting the calculator...")
			isCalculatorOn = false
		}
	}
}
